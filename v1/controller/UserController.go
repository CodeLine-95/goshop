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

// 登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 表单验证
	if len(username) == 0 {
		utils.Response(ctx, http.StatusBadRequest, 400, "用户名不能为空", nil)
		return
	}
	if len(password) == 0 {
		utils.Response(ctx, http.StatusBadRequest, 400, "密码不能为空", nil)
		return
	}
	// 验证用户密码
	var user model.User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		utils.Response(ctx, http.StatusBadRequest, 400, "该用户未注册", nil)
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utils.Response(ctx, http.StatusBadRequest, 400, "密码错误", nil)
		return
	}
	// 生成token
	token, err := common.ReleaseToken(user)
	if err != nil {
		utils.Response(ctx, http.StatusBadRequest, 400, "生成token失败", nil)
		return
	}
	// 获取 本机真实IP
	ip, _ := utils.ExternalIp()
	// 更新user
	DB.Model(&model.User{}).Where("id = ?", user.ID).Update("loginip", ip.String())
	// 返回值
	utils.Success(ctx, "登录成功", gin.H{
		"token": token,
	})
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	UserResultData := make(map[string]interface{})
	UserResultData["username"] = user.(model.User).Username
	UserResultData["usernick"] = user.(model.User).Usernick
	UserResultData["phone"] = user.(model.User).Phone
	UserResultData["email"] = user.(model.User).Email
	UserResultData["loginip"] = user.(model.User).Loginip
	// 返回值
	utils.Success(ctx, "登录成功", UserResultData)
}
