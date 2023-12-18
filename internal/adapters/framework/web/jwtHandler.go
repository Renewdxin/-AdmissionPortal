package web

import (
	jwt2 "github.com/Renewdxin/selfMade/internal/adapters/app/middleware/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWTHandlerAdapter struct {
	jwtAdapter jwt2.JWTAdapters
}

func NewJWTHandlerAdapter(jwtAdapter jwt2.JWTAdapters) *JWTHandlerAdapter {
	return &JWTHandlerAdapter{jwtAdapter: jwtAdapter}
}

func (j JWTHandlerAdapter) JWTHandler() gin.HandlerFunc {
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
			_, err := j.jwtAdapter.ParseToken(token)
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
