package tencent

import (
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

func (c *Client) Cbs() (client *cbs.Client, err error) {

	return cbs.NewClient(c.credential, c.RegionId, c.profile)

}

func (c *Client) Cvm() (client *cvm.Client, err error) {

	return cvm.NewClient(c.credential, c.RegionId, c.profile)

}

func (c *Client) Dnspod() (client *dnspod.Client, err error) {

	return dnspod.NewClient(c.credential, c.RegionId, c.profile)

}

func (c *Client) Lighthouse() (client *lighthouse.Client, err error) {

	return lighthouse.NewClient(c.credential, c.RegionId, c.profile)

}
