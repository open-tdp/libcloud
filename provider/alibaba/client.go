package alibaba

import (
	"github.com/open-tdp/libcloud/provider"

	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	swas "github.com/alibabacloud-go/swas-open-20200601/client"
)

type Client struct {
	*provider.ReqeustParam
	config *ac.Config
}

func NewClient(rq *provider.ReqeustParam) *Client {

	c := &Client{rq, nil}

	c.NewConfig()

	return c

}

func (c *Client) NewConfig() {

	c.config = &ac.Config{
		AccessKeyId:     &c.SecretId,
		AccessKeySecret: &c.SecretKey,
		RegionId:        &c.RegionId,
	}

}

func (c *Client) Alidns() (*alidns.Client, error) {

	return alidns.NewClient(c.config)

}

func (c *Client) Ecs() (*ecs.Client, error) {

	return ecs.NewClient(c.config)

}

func (c *Client) Swas() (*swas.Client, error) {

	return swas.NewClient(c.config)

}
