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

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
//func (adapter SMSAdapter) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error) {
//	config := &openapi.Config{
//		// 必填，您的 AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 必填，您的 AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//	}
//	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
//	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
//	_result = &openapi.Client{}
//	_result, _err = openapi.NewClient(config)
//	return _result, _err
//}
//
///**
// * API 相关
// * @param path Params
// * @return OpenApi.Params
// */
//func (adapter SMSAdapter) CreateApiInfo() (_result *openapi.Params) {
//	params := &openapi.Params{
//		// 接口名称
//		Action: tea.String("SendBatchSms"),
//		// 接口版本
//		Version: tea.String("2017-05-25"),
//		// 接口协议
//		Protocol: tea.String("HTTPS"),
//		// 接口 HTTP 方法
//		Method:   tea.String("POST"),
//		AuthType: tea.String("AK"),
//		Style:    tea.String("RPC"),
//		// 接口 PATH
//		Pathname: tea.String("/"),
//		// 接口请求体内容格式
//		ReqBodyType: tea.String("formData"),
//		// 接口响应体内容格式
//		BodyType: tea.String("json"),
//	}
//	_result = params
//	return _result
//}

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
