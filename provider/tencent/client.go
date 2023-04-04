package tencent

import (
	"strings"

	"github.com/open-tdp/libcloud/provider"
	"github.com/open-tdp/libcloud/setting"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

type Client struct {
	*provider.ReqeustParam
	credential *tc.Credential
	profile    *tp.ClientProfile
}

func NewClient(rq *provider.ReqeustParam) *Client {

	c := &Client{rq, nil, nil}

	c.NewCredential()
	c.NewProfile()

	return c

}

func (c *Client) NewCredential() {

	// 初始化
	credential := tc.NewCredential(c.SecretId, c.SecretKey)

	// 回传参数
	c.credential = credential

}

func (c *Client) NewProfile() {

	// 初始化
	profile := tp.NewClientProfile()

	// 调试模式
	profile.Debug = setting.Debug

	// 网络错误重试
	profile.NetworkFailureMaxRetries = 2

	// API 限频重试
	profile.RateLimitExceededMaxRetries = 2

	// 地域容灾机制
	profile.DisableRegionBreaker = false
	profile.BackupEndpoint = "ap-hongkong." + th.RootDomain

	// 按地域设置接口
	if c.Endpoint != "" {
		profile.HttpProfile.Endpoint = c.Endpoint // 完整域名
	} else if c.RegionId != "" {
		if !strings.HasSuffix(c.RegionId, "-ec") {
			profile.HttpProfile.Endpoint = c.Service + "." + c.RegionId + "." + th.RootDomain
		}
	}

	// 回传参数
	c.profile = profile

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
