package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/gorilla/websocket"

	"github.com/aiviaio/okex"
	"github.com/aiviaio/okex/events"
)

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 25 * time.Second
	PingPeriod = 15 * time.Second
)

type BizType string

const (
	// PublicPath is the public websocket api path
	BizPublic BizType = "public"
	// PrivatePath is the private websocket api path
	BizPrivate BizType = "private"
	// BusinessPath is the business websocket api path
	BizBusiness BizType = "business"
)

// ClientWs is the websocket api client
//
// https://www.okex.com/docs-v5/en/#websocket-api
type ClientWs struct {
	url           map[BizType]okex.BaseURL
	apiKey        string
	secretKey     []byte
	passphrase    string
	conn          map[BizType]*websocket.Conn
	mu            map[BizType]*sync.RWMutex
	ctx           context.Context
	Cancel        context.CancelFunc
	DoneChan      chan interface{}
	ErrChan       chan *events.Error
	SubscribeChan chan *events.Subscribe
	UnsubscribeCh chan *events.Unsubscribe
	LoginChan     chan *events.Login
	SuccessChan   chan *events.Success
	sendChan      map[BizType]chan []byte
	lastTransmit  map[BizType]*time.Time
	AuthRequested *time.Time
	Authorized    bool
	Private       *Private
	Public        *Public
	Business      *Business
	Trade         *Trade
	log           logr.Logger
}

type ClientOption func(c *ClientWs)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, url map[BizType]okex.BaseURL, opts ...ClientOption) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		url:          url,
		apiKey:       apiKey,
		secretKey:    []byte(secretKey),
		passphrase:   passphrase,
		conn:         make(map[BizType]*websocket.Conn),
		mu:           map[BizType]*sync.RWMutex{BizPublic: {}, BizPrivate: {}, BizBusiness: {}},
		ctx:          ctx,
		Cancel:       cancel,
		sendChan:     map[BizType]chan []byte{BizPublic: make(chan []byte, 3), BizPrivate: make(chan []byte, 3), BizBusiness: make(chan []byte, 3)},
		DoneChan:     make(chan interface{}, 32),
		lastTransmit: make(map[BizType]*time.Time),
		log:          logr.New(nil),
	}

	c.Private = NewPrivate(c)
	c.Public = NewPublic(c)
	c.Business = NewBusiness(c)
	c.Trade = NewTrade(c)

	for _, o := range opts {
		o(c)
	}
	return c
}

// Connect into the server
//
// https://www.okex.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect(p BizType) error {
	if c.conn[p] != nil {
		return nil
	}

	err := c.dial(p)
	if err == nil {
		return nil
	}
	c.log.Error(err, "failed to dial ws connection")

	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()
	counter := 0
	for {
		select {
		case <-ticker.C:
			err = c.dial(p)
			if err == nil {
				return nil
			}
			counter++
			c.log.Error(err, "failed to dial ws connection", "attempt", counter)
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// CheckConnect into the server
func (c *ClientWs) CheckConnect(p BizType) bool {
	if c.conn[p] != nil {
		return true
	}
	return false
}

// Login
//
// https://www.okex.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login() error {
	if c.Authorized {
		return nil
	}

	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
		return nil
	}

	now := time.Now()
	c.AuthRequested = &now
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}

	return c.Send(BizPrivate, okex.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okex.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(p BizType, ch []okex.ChannelName, args map[string]string) error {
	count := 1
	if len(ch) != 0 {
		count = len(ch)
	}

	tmpArgs := make([]map[string]string, count)
	tmpArgs[0] = args

	for i, name := range ch {
		tmpArgs[i] = map[string]string{}
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(p, okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(p BizType, ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(p, okex.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(p BizType, op okex.Operation, args []map[string]string, extras ...map[string]string) error {
	if op != okex.LoginOperation {
		err := c.Connect(p)
		if err == nil {
			if (p == BizPrivate || p == BizBusiness) && !c.Authorized {
				err = c.WaitForAuthorization()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}

	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.sendChan[p] <- j

	return nil
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(errCh chan *events.Error, subCh chan *events.Subscribe, unSub chan *events.Unsubscribe, lCh chan *events.Login, sCh chan *events.Success) {
	c.ErrChan = errCh
	c.SubscribeChan = subCh
	c.UnsubscribeCh = unSub
	c.LoginChan = lCh
	c.SuccessChan = sCh
}

// WaitForAuthorization waits for the auth response and try to log in if it was needed
func (c *ClientWs) WaitForAuthorization() error {
	if c.Authorized {
		return nil
	}

	if err := c.Login(); err != nil {
		return err
	}

	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for range ticker.C {
		if c.Authorized {
			return nil
		}
	}

	return nil
}

func (c *ClientWs) dial(p BizType) error {
	c.mu[p].Lock()
	conn, res, err := websocket.DefaultDialer.Dial(string(c.url[p]), nil)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}

		c.mu[p].Unlock()

		return fmt.Errorf("error %d: %w", statusCode, err)
	}
	defer res.Body.Close()
	c.log.Info("response ws", "status", res.Status)

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			c.mu[p].Lock()
			c.conn[p].Close()
			c.mu[p].Unlock()
		}()
		err := c.receiver(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: receiver") {
				c.ErrChan <- &events.Error{
					Event: "error",
					Msg:   err.Error(),
				}
			}
			c.log.Error(err, "receiver error")
		}
	}()

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			c.mu[p].Lock()
			c.conn[p].Close()
			c.mu[p].Unlock()
		}()
		err := c.sender(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: sender") {
				c.ErrChan <- &events.Error{
					Event: "error",
					Msg:   err.Error(),
				}
			}
			c.log.Error(err, "sender error")
		}
	}()

	c.conn[p] = conn
	c.mu[p].Unlock()

	return nil
}

func (c *ClientWs) sender(p BizType) error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for {
		c.mu[p].RLock()
		dataChan := c.sendChan[p]
		c.mu[p].RUnlock()

		select {
		case data := <-dataChan:
			c.mu[p].RLock()
			err := c.conn[p].SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to set write deadline for ws connection, error: %w", err)
			}

			w, err := c.conn[p].NextWriter(websocket.TextMessage)
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to get next writer for ws connection, error: %w", err)
			}

			if _, err = w.Write(data); err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to write data via ws connection, error: %w", err)
			}

			c.mu[p].RUnlock()

			if err := w.Close(); err != nil {
				return fmt.Errorf("failed to close ws connection, error: %w", err)
			}
		case <-ticker.C:
			c.mu[p].RLock()
			if c.conn[p] != nil && (c.lastTransmit[p] == nil || (c.lastTransmit[p] != nil && time.Since(*c.lastTransmit[p]) > PingPeriod)) {
				go func() {
					c.sendChan[p] <- []byte("ping")
				}()
			}

			c.mu[p].RUnlock()
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}

func (c *ClientWs) receiver(p BizType) error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			c.mu[p].RLock()
			err := c.conn[p].SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to set read deadline for ws connection, error: %w", err)
			}

			mt, data, err := c.conn[p].ReadMessage()
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to read message from ws connection, error: %w", err)
			}
			c.mu[p].RUnlock()

			now := time.Now()
			c.mu[p].Lock()
			c.lastTransmit[p] = &now
			c.mu[p].Unlock()

			if mt == websocket.TextMessage && string(data) != "pong" {
				e := &events.Basic{}
				if err := json.Unmarshal(data, e); err != nil {
					return fmt.Errorf("failed to unmarshall message from ws, error: %w", err)
				}
				go c.process(data, e)
			}
		}
	}
}

func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)

	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()

	return fmt.Errorf("operation cancelled: %s", msg)
}

// TODO: break each case into a separate function
func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.ErrChan != nil {
				c.ErrChan <- &e
			}
		}()

		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		if c.SubscribeChan != nil {
			c.SubscribeChan <- &e
		}

		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.UnsubscribeCh != nil {
				c.UnsubscribeCh <- &e
			}
		}()

		return true
	case "login":
		if time.Since(*c.AuthRequested).Seconds() > 30 {
			c.AuthRequested = nil
			_ = c.Login()
			break
		}

		c.Authorized = true

		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.LoginChan != nil {
				c.LoginChan <- &e
			}
		}()

		return true
	}

	if c.Private.Process(data, e) {
		return true
	}

	if c.Public.Process(data, e) {
		return true
	}

	if c.Business.Process(data, e) {
		return true
	}

	if e.ID != "" {
		if e.Code != 0 {
			ee := *e
			ee.Event = "error"

			return c.process(data, &ee)
		}

		e := events.Success{}
		_ = json.Unmarshal(data, &e)

		if c.SuccessChan != nil {
			c.SuccessChan <- &e
		}

		return true
	}

	return false
}

// WithLogger option sets new logger for websocket client
func WithLogger(sink logr.LogSink) ClientOption {
	return func(c *ClientWs) {
		c.log = c.log.WithSink(sink)
	}
}
