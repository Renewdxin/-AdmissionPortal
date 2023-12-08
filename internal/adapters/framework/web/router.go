package web

import (
	"github.com/Renewdxin/selfMade/internal/adapters/app/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	apiAccount := r.Group("/accounts")
	apiAccount.Use(middleware.JWTHandler())
	{
		apiAccount.POST("/login")
		apiAccount.POST("/logout")
		apiAccount.POST("/signup")
		apiAccount.POST("/password/forget/:id")
		apiAccount.POST("/password/change/:id")
	}

	apiPofile := r.Group("/profile")
	apiPofile.Use(middleware.JWTHandler())
	{
		apiPofile.GET("/Info/:id")
		apiPofile.GET("/account/:id")
	}
	return r
}
