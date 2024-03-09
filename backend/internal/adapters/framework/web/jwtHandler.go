package web

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWTHandlerAdapter struct {
	jwtPorts middleware.JwtApplicationPort
}

// NewJWTHandlerAdapter creates a new instance of JWTHandlerAdapter.
func NewJWTHandlerAdapter(jwtPorts middleware.JwtApplicationPort) *JWTHandlerAdapter {
	return &JWTHandlerAdapter{jwtPorts: jwtPorts}
}

// JWTHandler returns a middleware function that validates JWT tokens.
func (j *JWTHandlerAdapter) JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		_, err := j.jwtPorts.ParseToken(token)
		if err != nil {
			handleTokenError(c, err)
			return
		}

		c.Next()
	}
}

// extractToken extracts the token from the query parameters or headers.
func extractToken(c *gin.Context) string {
	token, exists := c.GetQuery("token")
	if !exists {
		token = c.GetHeader("Authorization")
	}
	return token
}

// handleTokenError handles errors returned from token parsing.
func handleTokenError(c *gin.Context, err error) {
	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		switch ve.Errors {
		case jwt.ValidationErrorExpired:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error processing token"})
	}
	c.Abort()
}
