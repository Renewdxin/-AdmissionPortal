package web

import "github.com/gin-gonic/gin"

type HomeHandlerPort interface {
	HomePage(c *gin.Context)
}
