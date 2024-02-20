package web

import (
	"fmt"
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
	fmt.Println("start")

	if err := c.ShouldBindJSON(&account); err != nil {
		logger.Logger.Log(logger.WarnLevel, "Log In Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}
	fmt.Println(account)

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
		logger.Logger.Logf(logger.WarnLevel, "Token Generate Error, id : %v", account.ID)
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
		logger.Logger.Logf(logger.WarnLevel, "Password Change Error, id : %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error! please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (handler AuthHandlerAdapter) ForgetPassword(c *gin.Context) {
	id := c.GetString("id")
	newPassword := c.GetString("newPassword")

	if err := handler.authCase.ForgetPasswordByID(id, newPassword); err != nil {
		logger.Logger.Logf(logger.WarnLevel, "Password Change Error, id : %v", id)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error! please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (handler AuthHandlerAdapter) CodeSend(c *gin.Context) {
	id := c.GetString("id")
	if err := handler.authCase.CodeSend(id); err != nil {
		logger.Logger.Log(logger.ErrorLevel, "Failed to send code, plz try again")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error! please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (handler AuthHandlerAdapter) CodeVerify(c *gin.Context) {
	code := c.GetString("code")
	id := c.GetString("id")
	if err := handler.authCase.CodeVerify(id, code); err != nil {
		logger.Logger.Log(logger.InfoLevel, "Failed to validate code")
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "wrong code",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
