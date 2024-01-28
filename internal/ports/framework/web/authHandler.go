package web

import "github.com/gin-gonic/gin"

type AuthHandlerPort interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	//Delete(c *gin.Context)
	ChangePassword(c *gin.Context)
	ForgetPassword(c *gin.Context)
}
