package mail

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

var Params = &openapi.Params{
	// 接口名称
	Action: tea.String("SendBatchSms"),
	// 接口版本
	Version: tea.String("2017-05-25"),
	// 接口协议
	Protocol: tea.String("HTTPS"),
	// 接口 HTTP 方法
	Method:   tea.String("POST"),
	AuthType: tea.String("AK"),
	Style:    tea.String("RPC"),
	// 接口 PATH
	Pathname: tea.String("/"),
	// 接口请求体内容格式
	ReqBodyType: tea.String("formData"),
	// 接口响应体内容格式
	BodyType: tea.String("json"),
}

type SMSPorts interface {
	//CreateClient(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error)
	//CreateApiInfo() (_result *openapi.Params)
	SendNotifySms(number, sign string) (_err error)
	SendCodeSms(number, sign string) (_err error)
}
