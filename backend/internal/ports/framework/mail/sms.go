package mail

import openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

type SMSPorts interface {
	CreateClient(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error)
	CreateApiInfo() (_result *openapi.Params)
	SendNotifySms(number, sign string) (_err error)
	SendCodeSms(number, sign string) (_err error)
}
