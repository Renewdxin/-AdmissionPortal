package web

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JWTHandlerAdapter struct {
	jwtPorts middleware.JwtApplicationPort
}

// NewJWTHandlerAdapter creates a new instance of JWTHandlerAdapter.
func NewJWTHandlerAdapter(jwtPorts middleware.JwtApplicationPort) *JWTHandlerAdapter {
	return &JWTHandlerAdapter{jwtPorts: jwtPorts}
}

func (j *JWTHandlerAdapter) JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		claims, err := j.jwtPorts.ParseToken(token)
		if err != nil {
			handleTokenError(c, err)
			return
		}

		// Check if the token is about to expire and refresh it
		if shouldRefreshToken(claims) {
			newToken, err := j.jwtPorts.GenerateToken(claims.UserID, claims.AppKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
				c.Abort()
				return
			}
			c.Header("Authorization", "Bearer "+newToken)
		}

		c.Next()
	}
}

// shouldRefreshToken checks if the token should be refreshed based on its expiration time.
func shouldRefreshToken(claims *middleware.Claims) bool {
	// Define the duration before expiration to refresh the token
	const refreshDuration = time.Minute * 30
	expirationTime := time.Unix(claims.ExpiresAt, 0)
	return time.Until(expirationTime) < refreshDuration
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
