package web

import (
	"fmt"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	authcase "github.com/Renewdxin/selfMade/internal/ports/app/auth"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/Renewdxin/selfMade/internal/ports/core/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		logger.Logger.Logf(logger.ErrorLevel, "Failed to bind account JSON during login: %v", err)
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

	appKey := os.Getenv("APP_KEY")
	token, err := handler.jwtPorts.GenerateToken(account.ID, appKey)
	if err != nil {
		logger.Logger.Logf(logger.ErrorLevel, "Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "please try again",
		})
		return
	}

	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome",
	})
}

// ... rest of the code remains the same
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

func (handler AuthHandlerAdapter) ChangePasswordByCode(c *gin.Context) {
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

func (handler AuthHandlerAdapter) ChangePasswordByPwd(c *gin.Context) {
	id := c.GetString("id")
	newPassword := c.GetString("newPassword")

	if err := handler.authCase.ForgetPasswordByID(id, newPassword); err != nil {
		logger.Logger.Logf(logger.WarnLevel, "Password Change Error, id : %v", id)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error!",
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
		logger.Logger.Logf(logger.InfoLevel, "id : %v Failed to validate code", id)
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "wrong code",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
