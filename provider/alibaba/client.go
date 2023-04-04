package alibaba

import (
	"os"

	"github.com/open-tdp/libcloud/provider"
	"github.com/open-tdp/libcloud/setting"

	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	au "github.com/alibabacloud-go/tea-utils/v2/service"
	tea "github.com/alibabacloud-go/tea/tea"

	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	swas "github.com/alibabacloud-go/swas-open-20200601/client"
)

type Client struct {
	*provider.ReqeustParam
	config  *ac.Config
	runtime *au.RuntimeOptions
}

func NewClient(rq *provider.ReqeustParam) *Client {

	c := &Client{rq, nil, nil}

	c.NewConfig()
	c.NewRuntime()

	return c

}

func (c *Client) NewConfig() {

	if setting.Debug {
		os.Setenv("DEBUG", "tea")
	}

	config := &ac.Config{
		AccessKeyId:     tea.String(c.SecretId),
		AccessKeySecret: tea.String(c.SecretKey),
		RegionId:        tea.String(c.RegionId),
	}

	// 回传参数
	c.config = config

}

func (c *Client) NewRuntime() {

	runtime := &au.RuntimeOptions{
		// 自动重试机制
		Autoretry:   tea.Bool(true),
		MaxAttempts: tea.Int(2),

		// 超时配置（单位 ms）
		ConnectTimeout: tea.Int(5000),
		ReadTimeout:    tea.Int(10000),
	}

	// 回传参数
	c.runtime = runtime

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
