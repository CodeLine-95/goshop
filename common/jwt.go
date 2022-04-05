package common

import (
	"github.com/dgrijalva/jwt-go"
	"goshop/model"
	"time"
)

var jwtKey = []byte("job_hand_some")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成jwt token
func ReleaseToken(user model.User) (string, error) {
	expirTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "jobhandsome.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
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
