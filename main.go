package main

import (
	"github.com/Renewdxin/selfMade/global"
	"github.com/Renewdxin/selfMade/internal/router"
	"github.com/Renewdxin/selfMade/pkg/jwt"
	setting2 "github.com/Renewdxin/selfMade/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	setupSetting()
}

func main() {
	r := router.NewRouter()
	r.POST("/auth", jwt.AuthHandler)
	r.POST("/home", jwt.AuthHandler, func(context *gin.Context) {
		username := context.MustGet("username").(string)
		context.JSON(http.StatusOK, gin.H{"weclome:": username})
	})
	r.Run(":9000")
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}
