package main

import (
	"context"
	"log"
	"os"

	"github.com/aiviaio/okex"
	"github.com/aiviaio/okex/api"
	"github.com/aiviaio/okex/events"
	"github.com/aiviaio/okex/events/private"
	"github.com/aiviaio/okex/events/public"
	ws_private_requests "github.com/aiviaio/okex/requests/ws/private"
	ws_public_requests "github.com/aiviaio/okex/requests/ws/public"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	apiKey := os.Getenv("OKEX_API_KEY")
	secretKey := os.Getenv("OKEX_SECRET_KEY")
	passphrase := os.Getenv("OKEX_PASSPHRASE")
	dest := okex.DemoServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, dest, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Rest.Account.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %+v", response)

	errChan := make(chan *events.Error)
	subChan := make(chan *events.Subscribe)
	uSubChan := make(chan *events.Unsubscribe)
	lCh := make(chan *events.Login)
	sCh := make(chan *events.Success)
	// to receive unique events individually in separated channels
	client.Ws.SetChannels(errChan, subChan, uSubChan, lCh, sCh)

	oCh := make(chan *private.Order)
	// subscribe into orders private channel
	// it will do the login process and wait until authorization confirmed
	err = client.Ws.Private.Order(ws_private_requests.Order{
		InstType: okex.SwapInstrument,
	}, oCh)
	if err != nil {
		log.Fatalln(err)
	}

	iCh := make(chan *public.Instruments)
	// subscribe into instruments public channel
	// it doesn't need any authorization
	err = client.Ws.Public.Instruments(ws_public_requests.Instruments{
		InstType: okex.SwapInstrument,
	}, iCh)
	if err != nil {
		log.Fatalln("Instruments", err)
	}

	tCh := make(chan *public.Tickers)
	/* err = client.Ws.Public.Tickers(ws_public_requests.Tickers{
		InstID: "BTC-USD-SWAP",
	}, tCh)
	if err != nil {
		log.Fatalln("Instruments", err)
	} */

	cCh := make(chan *public.Candlesticks)
	err = client.Ws.Business.Candlesticks(ws_public_requests.Candlesticks{
		InstID:  "BTC-USD-SWAP",
		Channel: okex.CandleStick1m,
	}, cCh)
	if err != nil {
		log.Fatalln("Candlesticks", err)
	}

	// starting on listening
	for {
		select {
		case <-lCh:
			log.Print("[Authorized]")
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("[Subscribed]\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("[Unsubscribed]\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("[Error]\t%+v", err)
		case o := <-oCh:
			log.Print("[Event]\tOrder")
			for _, p := range o.Orders {
				log.Printf("\t%+v", p)
			}
		case i := <-iCh:
			log.Print("[Event]\tInstrument")
			for _, p := range i.Instruments {
				log.Printf("\t%+v", p)
			}
		case t := <-tCh:
			log.Print("[Event]\tTicker")
			for _, p := range t.Tickers {
				log.Printf("\t%+v", p)
			}
		case c := <-cCh:
			log.Print("[Event]\tCandlestick")
			for _, p := range c.Candles {
				log.Printf("\t%+v", p)
			}
		case b := <-client.Ws.DoneChan:
			log.Printf("[End]:\t%v", b)
			return
		}
	}
}
