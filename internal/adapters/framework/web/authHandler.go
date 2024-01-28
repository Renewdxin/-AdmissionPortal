package web

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "account error",
		})
		return
	}
	fmt.Println(account)

	if err := handler.authCase.LogIn(account.ID, account.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	token, err := handler.jwtPorts.GenerateToken(account.ID, "hahaha")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "hahah",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "welcome",
		"token": token,
	})
}

func (handler AuthHandlerAdapter) Register(c *gin.Context) {
	var account auth.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	if err := handler.authCase.RegisterByEmail(account.ID, account.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
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
