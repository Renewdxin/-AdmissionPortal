package web

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	authcase "github.com/Renewdxin/selfMade/internal/ports/app/auth"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlerAdapter struct {
	authCase authcase.AuthorizeApplicationPort
	jwtPorts middleware.JwtApplicationPort
}

func NewAuthHandlerAdapter(authCase authcase.AuthorizeApplicationPort, jwtPorts middleware.JwtApplicationPort) *AuthHandlerAdapter {
	return &AuthHandlerAdapter{
		authCase: authCase,
		jwtPorts: jwtPorts,
	}
}

func (handler AuthHandlerAdapter) Login(c *gin.Context) {
	var account auth.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		logger.Logger.Log(logger.WarnLevel, "Log In Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}

	if err := handler.authCase.LogIn(account.ID, account.Password); err != nil {
		logger.Logger.Logf(logger.ErrorLevel, " Password Or Account Invalid, id: %v, password: %v", account.ID, account.Password)
		c.JSON(http.StatusOK, gin.H{
			"msg": "Password Or Account Invalid",
		})
		return
	}

	token, err := handler.jwtPorts.GenerateToken(account.ID, "gvgkjbjkttsrt")
	if err != nil {
		logger.Logger.Log(logger.WarnLevel, "Token Generate Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "please try again",
		})
		return
	}

	c.Header("tok", token)
	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome",
	})
}

func (handler AuthHandlerAdapter) Register(c *gin.Context) {
	var account auth.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		logger.Logger.Log(logger.WarnLevel, "register account bind error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	if err := handler.authCase.RegisterByID(account.ID, account.Password); err != nil {
		logger.Logger.Log(logger.WarnLevel, "Token Generate Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "register finished",
	})
}

func (handler AuthHandlerAdapter) ChangePassword(c *gin.Context) {
	id := c.GetString("id")
	newPassword := c.GetString("newPassword")
	oldPassword := c.GetString("oldPassword")

	if err := handler.authCase.ChangePassword(id, oldPassword, newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (handler AuthHandlerAdapter) ForgetPassword(c *gin.Context) {
	id := c.GetString("id")
	newPassword := c.GetString("newPassword")

	if err := handler.authCase.ForgetPasswordByEmail(id, newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
