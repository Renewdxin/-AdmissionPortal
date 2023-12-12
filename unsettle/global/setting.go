package global

import (
	"github.com/Renewdxin/selfMade/unsettle/pkg/logger"
	"github.com/Renewdxin/selfMade/unsettle/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
)
