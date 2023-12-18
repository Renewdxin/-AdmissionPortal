package middleware

import "github.com/dgrijalva/jwt-go"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
type Claims struct {
	UserID string `json:"user_id"`
	AppKey string `json:"app_key"`
	jwt.StandardClaims
}

type JwtPort interface {
	GetJWTSecret() []byte
	GenerateToken(userid string, AppKey string) (string, error)
	ParseToken(tokenString string) (*Claims, error)
}
