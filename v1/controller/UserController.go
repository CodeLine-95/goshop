package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"goshop/common"
	"goshop/model"
	"goshop/utils"
	"net/http"
)

//判断手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	usernick := ctx.PostForm("usernick")

	if len(password) < 6 {
		utils.Response(ctx, http.StatusUnprocessableEntity, 422, "密码不能小于6位！", nil)
		return
	}

	if len(username) == 0 {
		username = utils.RandomString(10)
	}

	if len(usernick) == 0 {
		usernick = utils.RandomString(10)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(ctx, http.StatusUnprocessableEntity, 500, "加密错误", nil)
		return
	}

	user := model.User{
		Username: username,
		Usernick: usernick,
		Password: string(hashPassword),
	}
	DB.Create(&user)
	//返回结果

	utils.Success(ctx, "注册成功", nil)
}
