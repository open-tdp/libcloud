package cloudflare

import (
	"context"
	"strconv"

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

// 处理错误

func (c *Client) Error(err any) *provider.ResponseError {

	if er, ok := err.(*cf.Error); ok {
		code := er.Messages[0].Code
		msg := er.Messages[0].Message
		return &provider.ResponseError{
			Code:    strconv.Itoa(code),
			Message: msg,
		}
	}

	if er, ok := err.(error); ok {
		return &provider.ResponseError{
			Code:    "Nil",
			Message: er.Error(),
		}
	}

	if er, ok := err.(string); ok {
		return &provider.ResponseError{
			Code:    "Nil",
			Message: er,
		}
	}

	return &provider.ResponseError{
		Code:    "Nil",
		Message: "Unkown",
	}

}
