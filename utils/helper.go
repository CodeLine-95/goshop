package utils

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, msg, data)
}

// Returns an int >= min, < max
//func RandomInt(min, max int) int {
//	return min + rand.Intn(max-min)
//}

// 生成指定长度的随机字符
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmMNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
