okex
====

[![Go Reference](https://pkg.go.dev/badge/github.com/aiviaio/okex.svg)](https://pkg.go.dev/github.com/aiviaio/okex)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/aiviaio/okex.svg)](https://github.com/aiviaio/okex)
[![GoReportCard example](https://goreportcard.com/badge/github.com/aiviaio/okex)](https://goreportcard.com/report/github.com/aiviaio/okex)
[![GitHub license](https://img.shields.io/github/license/aiviaio/okex.svg)](https://github.com/aiviaio/okex/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/aiviaio/okex.svg)](https://GitHub.com/aiviaio/okex/releases/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

*NOTICE:*
> PACKAGE IS CURRENTLY UNDER HEAVY DEVELOPMENT AND THERE IS NO GUARANTY FOR STABILITY UNTIL V1 RELEASE.

Okex V5 Golang API

A complete golang wrapper for [Okex](https://www.okex.com) V5 API. Pretty simple and easy to use. For more info about
Okex V5 API [read here](https://www.okx.com/docs-v5/en/).

Installation
-----------------

```bash
go get github.com/aiviaio/okex@v1.0.5
```

Usage
-----------

```go
package main

import (
 "context"
 "log"

 "github.com/aiviaio/okex"
 "github.com/aiviaio/okex/api"
 "github.com/aiviaio/okex/events"
 "github.com/aiviaio/okex/events/private"
 "github.com/aiviaio/okex/events/public"
 ws_private_requests "github.com/aiviaio/okex/requests/ws/private"
 ws_public_requests "github.com/aiviaio/okex/requests/ws/public"
)

func main() {
 apiKey := "YOUR-API-KEY"
 secretKey := "YOUR-SECRET-KEY"
 passphrase := "YOUR-PASS-PHRASE"
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

 oCh := make(chan *private.Order)
 iCh := make(chan *public.Instruments)

 // to receive unique events individually in separated channels
 client.Ws.SetChannels(errChan, subChan, uSubChan, lCh, sCh)

 // subscribe into orders private channel
 // it will do the login process and wait until authorization confirmed
 err = client.Ws.Private.Order(ws_private_requests.Order{
  InstType: okex.SwapInstrument,
 }, oCh)
 if err != nil {
  log.Fatalln(err)
 }

 // subscribe into instruments public channel
 // it doesn't need any authorization
 err = client.Ws.Public.Instruments(ws_public_requests.Instruments{
  InstType: okex.SwapInstrument,
 }, iCh)
 if err != nil {
  log.Fatalln("Instruments", err)
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
  case b := <-client.Ws.DoneChan:
   log.Printf("[End]:\t%v", b)
   return
  }
 }
}
```

Supporting APIs
---------------

* [Rest](https://www.okx.com/docs-v5/en/#rest-api)
  * [Trade](https://www.okx.com/docs-v5/en/#rest-api-trade) (except demo special trading endpoints)
  * [Funding](https://www.okx.com/docs-v5/en/#rest-api-funding)
  * [Account](https://www.okx.com/docs-v5/en/#rest-api-account)
  * [SubAccount](https://www.okx.com/docs-v5/en/#rest-api-subaccount)
  * [Market Data](https://www.okx.com/docs-v5/en/#rest-api-market-data)
  * [Public Data](https://www.okx.com/docs-v5/en/#rest-api-public-data)
  * [Trading Data](https://www.okx.com/docs-v5/en/#rest-api-trading-data)

* [Ws](https://www.okx.com/docs-v5/en/#websocket-api)
  * [Private Channel](https://www.okx.com/docs-v5/en/#websocket-api-private-channel) (except demo special trading
      endpoints)
  * [Public Channel](https://www.okx.com/docs-v5/en/#websocket-api-public-channels)
  * [Trade](https://www.okx.com/docs-v5/en/#websocket-api-trade)

Features
--------

* All [requests](/requests), [responses](/responses), and [events](events) are well typed and will convert into the
  language built-in types instead of using API's strings. *Note that zero values will be replaced with non-existing
  data.*
* Fully automated authorization steps for both [REST](/api/rest) and [WS](/api/ws)
* To receive websocket events you can choose [RawEventChan](/api/ws/client.go#L25)
  , [StructuredEventChan](/api/ws/client.go#L28), or provide your own
  channels. [More info](https://github.com/aiviaio/okex/wiki/Handling-WS-events)
