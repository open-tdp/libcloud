package cloudflare

import (
	"context"

	"github.com/open-tdp/libcloud/provider"

	cf "github.com/cloudflare/cloudflare-go"
)

type Client struct {
	*provider.ReqeustParam
	ctx context.Context
	api *cf.API
}

func NewClient(rq *provider.ReqeustParam) *Client {

	c := &Client{rq, context.Background(), nil}

	c.NewApi()

	return c

}

func (c *Client) NewApi() error {

	api, err := cf.NewWithAPIToken(c.SecretKey)

	c.api = api

	return err

}
