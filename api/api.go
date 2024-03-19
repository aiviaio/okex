package api

import (
	"context"

	"github.com/aiviaio/okex"
	"github.com/aiviaio/okex/api/rest"
	"github.com/aiviaio/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex.Destination, restOpts []rest.ClientOption, wsOpts []ws.ClientOption) (*Client, error) {
	restURL := okex.RestURL
	wsPubURL := okex.PublicWsURL
	wsPriURL := okex.PrivateWsURL
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.AwsPublicWsURL
		wsPriURL = okex.AwsPrivateWsURL
	case okex.DemoServer:
		restURL = okex.DemoRestURL
		wsPubURL = okex.DemoPublicWsURL
		wsPriURL = okex.DemoPrivateWsURL
	case okex.OmegaServer:
		restURL = okex.OmegaRestURL
		wsPubURL = okex.OmegaPublicWsURL
		wsPriURL = okex.OmegaPrivateWsURL
	case okex.BusinessServer:
		restURL = okex.RestURL
		wsPubURL = okex.BusinessWsURL
		wsPriURL = okex.BusinessWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination, restOpts...)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseURL{true: wsPriURL, false: wsPubURL}, wsOpts...)

	return &Client{r, c, ctx}, nil
}
