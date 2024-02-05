package mail

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

type SMSAdapter struct{}

func NewSMSAdapter() SMSAdapter {
	return SMSAdapter{}
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func (adapter SMSAdapter) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

/**
 * API 相关
 * @param path params
 * @return OpenApi.Params
 */
func (adapter SMSAdapter) CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
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
	_result = params
	return _result
}

func (adapter SMSAdapter) SendNotifySms(number, sign string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := adapter.CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	params := adapter.CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["TemplateCode"] = tea.String("NOTIFY_TEMPLATE")
	// body params
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
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}

func (adapter SMSAdapter) SendCodeSms(number, sign string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := adapter.CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	params := adapter.CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["TemplateCode"] = tea.String(os.Getenv("CODE_TEMPLATE"))
	// body params
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
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}
