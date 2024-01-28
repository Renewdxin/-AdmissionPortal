package web

import "github.com/gin-gonic/gin"

type JWTHandlerPort interface {
	JWTHandler() gin.HandlerFunc
}
