package middleware

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/logger"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTAdapter is a struct that encapsulates functionality related to JWT (JSON Web Token) operations.
type JWTAdapter struct{}

// NewJWTAdapters is a constructor function for JWTAdapter, initializing it with the provided logger.
func NewJWTAdapters() *JWTAdapter {
	return &JWTAdapter{}
}

// GetJWTSecret returns the JWT secret key from the global configuration.
func (j JWTAdapter) GetJWTSecret() []byte {
	return []byte("xiyou3g")
}

// GenerateToken generates a JWT token with the specified user ID and application key.
// It returns the generated token as a string and any error encountered during the process.
func (j JWTAdapter) GenerateToken(userid string, AppKey string) (string, error) {
	expireTime := time.Now().Add(7200)
	claims := middleware.Claims{
		UserID: userid,
		AppKey: AppKey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "xiyou3g",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(j.GetJWTSecret())
}

// ParseToken parses the provided JWT token string and returns the claims if the token is valid.
// It returns an error if the token is invalid or any other error occurs during parsing.
func (j JWTAdapter) ParseToken(tokenString string) (*middleware.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &middleware.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*middleware.Claims); tokenClaims.Valid && ok {
		return claims, nil
	}

	// Log a message for an invalid token.
	logger.Logger.Log(4, "INVALID TOKEN")
	return nil, err
}
