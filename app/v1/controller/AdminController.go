package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

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

// 分配角色
func AssignRole(ctx *gin.Context) {
	params, _ := utils.DataMapByRequest(ctx)
	int_id, _ := params["uid"].(uint)
	// 获取数据库句柄
	DB := config.GetDB()
	// 获取用户
	var admin model.Admin
	DB.First(&admin, int_id)
	admin.Group = params["groupStr"].(string)
	result := DB.Save(&admin)
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "分配成功", nil)
	}
}
