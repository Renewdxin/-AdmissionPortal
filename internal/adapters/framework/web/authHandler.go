package web

import (
	authcase "github.com/Renewdxin/selfMade/internal/ports/app/auth"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authCase authcase.AuthcasePorts
}

func NewAuthHandler(authCase authcase.AuthcasePorts) *AuthHandler {
	return &AuthHandler{
		authCase: authCase,
	}
}

func (handler AuthHandler) Login(c *gin.Context) {
	var account auth.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	if err := handler.authCase.LogIn(account.ID, account.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome",
	})
}

func (handler AuthHandler) Register(c *gin.Context) {
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

func (handler AuthHandler) ChangePassword(c *gin.Context) {
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

func (handler AuthHandler) ForgetPassword(c *gin.Context) {
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
