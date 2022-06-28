package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

// Register 注册
func Register(ctx *gin.Context) {
	DB := config.GetDB()
	// 获取参数
	var params model.Users
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误:"+err.Error(), nil)
		return
	}
	fmt.Println(params)
	if len(params.Password) < 6 {
		utils.Fail(ctx, "密码不能小于6位！", nil)
		return
	}
	if len(params.Username) == 0 {
		params.Username = utils.RandomString(10)
	}

	if len(params.Usernick) == 0 {
		params.Usernick = utils.RandomString(10)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(ctx, "加密错误", nil)
		return
	}
	params.Password = string(hashPassword)
	// uuid
	params.UserID = utils.UUID()
	// 创建
	err = DB.Create(&params).Error
	if err != nil {
		utils.Fail(ctx, "注册失败", nil)
		return
	}
	//返回结果
	utils.Success(ctx, "注册成功", nil)
}

// Login 登录
func Login(ctx *gin.Context) {
	// 初始化数据库句柄
	DB := config.GetDB()
	// 定义使用模型
	var params model.Users
	var user model.Users
	// 绑定获取请求参数
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误:"+err.Error(), nil)
		return
	}
	if len(params.Username) == 0 {
		utils.Fail(ctx, "用户名不能为空", nil)
		return
	}
	if len(params.Password) == 0 {
		utils.Fail(ctx, "密码不能为空", nil)
		return
	}
	if len(params.Password) < 6 {
		utils.Fail(ctx, "密码不能小于6位！", nil)
		return
	}
	// 获取用户
	DB.Where("username = ?", params.Username).First(&user)
	if user.ID == 0 {
		utils.Fail(ctx, "该用户未注册", nil)
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		utils.Fail(ctx, "密码错误", nil)
		return
	}
	// 生成token
	token, tokenErr := config.ReleaseToken(user.ID)
	if tokenErr != nil {
		utils.Fail(ctx, "生成token失败", nil)
		return
	}
	// 获取 本机真实IP
	ip, _ := utils.ExternalIp()
	user.Loginip = ip.String()
	// 更新
	resultErr := DB.Save(&user).Error
	if resultErr != nil {
		utils.Fail(ctx, "登录失败", nil)
		return
	}
	//返回结果
	utils.Success(ctx, "登录成功", gin.H{
		"token": token,
	})
}
