package web

import "github.com/gin-gonic/gin"

type AuthHandlerPort interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	//Delete(c *gin.Context)
	ChangePasswordByCode(c *gin.Context)
	ChangePasswordByPwd(c *gin.Context)
	CodeSend(c *gin.Context)
	CodeVerify(c *gin.Context)
}
