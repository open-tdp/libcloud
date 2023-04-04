package alibaba

import (
	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	swas "github.com/alibabacloud-go/swas-open-20200601/client"
)

func (c *Client) Alidns() (*alidns.Client, error) {

	return alidns.NewClient(c.config)

}

func (c *Client) Ecs() (*ecs.Client, error) {

	return ecs.NewClient(c.config)

}

func (c *Client) Swas() (*swas.Client, error) {

	return swas.NewClient(c.config)

}
