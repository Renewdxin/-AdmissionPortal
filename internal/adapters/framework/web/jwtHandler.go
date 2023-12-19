package web

import (
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWTHandlerAdapter struct {
	jwtPorts middleware.JwtPort
}

func NewJWTHandlerAdapter(jwtPorts middleware.JwtPort) *JWTHandlerAdapter {
	return &JWTHandlerAdapter{jwtPorts: jwtPorts}
}

func (j *JWTHandlerAdapter) JWTHandler() gin.HandlerFunc {
	// get token
	// token exist ? validate : abort
	// next or abort
	return func(c *gin.Context) {
		token, exist := c.GetQuery("token")
		if !exist {
			token = c.GetHeader("token")
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": jwt.ErrSignatureInvalid})
			c.Abort()
			return
		} else {
			_, err := j.jwtPorts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					c.JSON(http.StatusOK, gin.H{"msg": jwt.ValidationErrorExpired})
				default:
					c.JSON(http.StatusUnauthorized, gin.H{"msg": jwt.ValidationErrorNotValidYet})
				}
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
