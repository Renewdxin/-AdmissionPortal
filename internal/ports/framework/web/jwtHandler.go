package web

import "github.com/gin-gonic/gin"

type JWTHandlerPorts interface {
	JWTHandler() gin.HandlerFunc
}
