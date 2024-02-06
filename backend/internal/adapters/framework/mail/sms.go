package mail

import (
	"github.com/Renewdxin/selfMade/internal/ports/framework/mail"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

type SMSAdapter struct {
	Client *openapi.Client
}

func NewSMSAdapter(accessKeyId *string, accessKeySecret *string) (SMSAdapter, error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	client, err := openapi.NewClient(config)
	if err != nil {
		return SMSAdapter{}, err
	}

	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	return SMSAdapter{Client: client}, nil
}

func (adapter SMSAdapter) SendNotifySms(number, sign string) (_err error) {
	// query Params
	queries := map[string]interface{}{}
	queries["TemplateCode"] = tea.String("NOTIFY_TEMPLATE")
	// body Params
	body := map[string]interface{}{}
	body["PhoneNumberJson"] = tea.String(number)
	body["SignNameJson"] = tea.String(sign)
	body["TemplateParamJson"] = nil
	body["SmsUpExtendCodeJson"] = nil
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
		Body:  body,
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = adapter.Client.CallApi(mail.Params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}

func (adapter SMSAdapter) SendCodeSms(number, sign string) (_err error) {
	// query Params
	queries := map[string]interface{}{}
	queries["TemplateCode"] = tea.String(os.Getenv("CODE_TEMPLATE"))
	// body Params
	body := map[string]interface{}{}
	body["PhoneNumberJson"] = tea.String(number)
	body["SignNameJson"] = tea.String(sign)
	body["TemplateParamJson"] = nil
	body["SmsUpExtendCodeJson"] = nil
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
		Body:  body,
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = adapter.Client.CallApi(mail.Params, request, runtime)
	if _err != nil {
		return _err
	}

	return _err
}
