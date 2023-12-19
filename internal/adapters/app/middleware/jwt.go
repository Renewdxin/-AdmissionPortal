package middleware

import (
	"github.com/Renewdxin/selfMade/internal/adapters/framework/global"
	"github.com/Renewdxin/selfMade/internal/ports/app/middleware"
	"github.com/Renewdxin/selfMade/internal/ports/framework/logger"
	"github.com/Renewdxin/selfMade/unsettle/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTAdapters struct {
	logger logger.LoggerPorts
}

func NewJWTAdapters(logger logger.LoggerPorts) *JWTAdapters {
	return &JWTAdapters{
		logger: logger,
	}
}

func (j JWTAdapters) GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func (j JWTAdapters) GenerateToken(userid string, AppKey string) (string, error) {
	now := time.Now()
	expireTime := now.Add(7200)
	claims := middleware.Claims{
		UserID: util.EncodingMD5(userid),
		AppKey: util.EncodingMD5(AppKey),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(j.GetJWTSecret())
}

func (j JWTAdapters) ParseToken(tokenString string) (*middleware.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &middleware.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*middleware.Claims); tokenClaims.Valid && ok {
		return claims, nil
	}

	j.logger.Log(4, "INVALID TOKEN")
	return nil, err
}
