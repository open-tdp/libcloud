package provider

// 请求参数

type ReqeustParam struct {
	SecretId  string `note:"访问密钥 Id"`
	SecretKey string `note:"访问密钥 Key"`
	RegionId  string `note:"资源所在区域"`
	Endpoint  string `note:"指定接口域名"`
	Service   string `note:"产品名称"`
	Version   string `note:"接口版本"`
	Action    string `note:"接口名称"`
	Payload   any    `note:"结构化数据"`
}

// 请求结果

type ResponseResult struct {
	Error   ResponseError `note:"错误信息"`
	Message []string      `note:"消息内容"`
	Payload any           `note:"请求结果"`
}

// 错误信息

type ResponseError struct {
	Code    string `note:"错误代码"`
	Message string `note:"错误信息"`
}

func (e *ResponseError) Create(err any) {

	if er, ok := err.(error); ok {
		e.Message = er.Error()
		e.Code = "Nil"
		return
	}

	if er, ok := err.(string); ok {
		e.Message = er
		e.Code = "Nil"
		return
	}

	e.Message = "Unkown"
	e.Code = "Nil"

}
