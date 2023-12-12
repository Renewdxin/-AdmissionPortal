package web

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler AuthHandler) Login(c *gin.Context) {

}

func (handler AuthHandler) Register(c *gin.Context) {

}

func (handler AuthHandler) Delete(c *gin.Context) {

}
