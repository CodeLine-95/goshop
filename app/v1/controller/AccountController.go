package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"goshop/utils/captcha"
	"goshop/utils/time"
)

//判断手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.Admin
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 注册
func Register(ctx *gin.Context) {
	DB := config.GetDB()
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	username := params["username"].(string)
	password := params["password"].(string)
	usernick := params["usernick"].(string)

	if len(password) < 6 {
		utils.Fail(ctx, "密码不能小于6位！", nil)
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
		utils.Fail(ctx, "加密错误", nil)
		return
	}

	user := model.Admin{
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
	DB := config.GetDB()
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	username := params["username"].(string)
	password := params["password"].(string)
	captchaId := params["captchaId"].(string)
	captchaCode := params["captchaCode"].(string)
	// 表单验证
	if len(username) == 0 {
		utils.Fail(ctx, "用户名不能为空", nil)
		return
	}
	if len(password) == 0 {
		utils.Fail(ctx, "密码不能为空", nil)
		return
	}
	if len(captchaCode) == 0 {
		utils.Fail(ctx, "验证码不能为空", nil)
		return
	}
	// 验证用户密码
	var user model.Admin
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		utils.Fail(ctx, "该用户未注册", nil)
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utils.Fail(ctx, "密码错误", nil)
		return
	}
	// 验证验证码
	if !captcha.VerfiyCaptcha(captchaId, captchaCode) {
		utils.Fail(ctx, "验证码不正确", nil)
		return
	}
	// 生成token
	token, err := config.ReleaseToken(user)
	if err != nil {
		utils.Fail(ctx, "生成token失败", nil)
		return
	}
	// 获取 本机真实IP
	ip, _ := utils.ExternalIp()
	// 更新user
	var LocalTime time.LocalTime
	updateData := make(map[string]interface{})
	updateData["loginip"] = ip.String()
	updateData["login_at"] = LocalTime.FormatDateString(LocalTime.String())
	DB.Model(&user).Where("id = ?", user.ID).Updates(updateData)
	// 返回值
	utils.Success(ctx, "登录成功", gin.H{
		"token": token,
	})
}

// 生成验证码
func GetCaptcha(ctx *gin.Context) {
	captcha.GenerateCapcha(ctx)
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	UserResultData := make(map[string]interface{})
	userModel := user.(model.Admin)
	UserResultData["username"] = userModel.Username
	UserResultData["usernick"] = userModel.Usernick
	UserResultData["avatar"] = userModel.Avatar
	UserResultData["phone"] = userModel.Phone
	UserResultData["email"] = userModel.Email
	UserResultData["loginip"] = userModel.Loginip
	UserResultData["logintime"] = userModel.LoginAt
	UserResultData["group"] = userModel.Group
	// 返回值
	utils.Success(ctx, "获取成功", UserResultData)
}
