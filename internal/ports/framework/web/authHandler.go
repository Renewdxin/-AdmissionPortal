package web

import "github.com/gin-gonic/gin"

type AuthHandlerPorts interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Delete(c *gin.Context)
}
