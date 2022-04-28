package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"goshop/model"
	"goshop/utils"
	"time"
)

var jwtKey = []byte(viper.GetString("server.jwtKey"))

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成jwt token
func ReleaseToken(user model.Admin) (string, error) {
	ip, _ := utils.ExternalIp()
	expirTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    ip.String(),
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// 解析jwt token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
