package jwt

import (
	"errors"
	"github.com/Renewdxin/selfMade/global"
	"github.com/Renewdxin/selfMade/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
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

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthHandler(c *gin.Context) {
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	if user.Username == "cyl" && user.Password == "123456" {
		//生成token
		tokenString, _ := GenerateToken(user.Username, user.Password)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
