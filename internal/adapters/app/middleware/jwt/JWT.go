package jwt

import (
	"errors"
	"github.com/Renewdxin/selfMade/internal/adapters/framework/global"
	"github.com/Renewdxin/selfMade/unsettle/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(AppID string, AppKey string) (string, error) {
	now := time.Now()
	expireTime := now.Add(7200)
	claims := Claims{
		AppID:  util.EncodingMD5(AppID),
		AppKey: util.EncodingMD5(AppKey),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(GetJWTSecret())
}

func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); tokenClaims.Valid && ok {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
