package ws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aiviaio/okex"
	"github.com/aiviaio/okex/events"
	"github.com/aiviaio/okex/events/public"
	requests "github.com/aiviaio/okex/requests/ws/public"
)

// Public
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels
type Business struct {
	*ClientWs
	cCh chan *public.Candlesticks
}

// NewBusiness returns a pointer to a fresh Business
func NewBusiness(c *ClientWs) *Business {
	return &Business{ClientWs: c}
}

// Candlesticks
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Business) Candlesticks(req requests.Candlesticks, ch ...chan *public.Candlesticks) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.cCh = ch[0]
	}
	return c.Subscribe(BizBusiness, []okex.ChannelName{}, m)
}

// UCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Business) UCandlesticks(req requests.Candlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.cCh = nil
	}
	return c.Unsubscribe(BizBusiness, []okex.ChannelName{}, m)
}

func (c *Business) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		default:
			chName := fmt.Sprint(ch)
			// candlestick channels
			if strings.Contains(chName, "candle") {
				e := public.Candlesticks{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				go func() {
					if c.cCh != nil {
						c.cCh <- &e
					}
				}()
				return true
			}
		}
	}
	return false
}
