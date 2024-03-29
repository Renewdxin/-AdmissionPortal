package middleware

import (
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type JWTAdapter struct{}

// NewJWTAdapters is a constructor function for JWTAdapter.
func NewJWTAdapters() *JWTAdapter {
	return &JWTAdapter{}
}

// GetJWTSecret returns the JWT secret key from the environment variable.
func (j JWTAdapter) GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// GenerateToken generates a JWT token with the specified user ID and application key.
func (j JWTAdapter) GenerateToken(userid string, AppKey string) (string, error) {
	expireTime := time.Now().Add(2 * time.Hour) // Use time.Hour for clarity
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
func (j JWTAdapter) ParseToken(tokenString string) (*middleware.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &middleware.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.GetJWTSecret(), nil
	})
	if err != nil {
		logrus.WithField("error", err).Error("Error parsing token")
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*middleware.Claims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		err := jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
		logrus.WithField("error", err).Error("Invalid token")
		return nil, err
	}
}
