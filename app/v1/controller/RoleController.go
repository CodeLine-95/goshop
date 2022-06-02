package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

// 获取角色列表
func RoleLists(ctx *gin.Context) {
	DB := config.InitDB()
	params, _ := utils.DataMapByRequest(ctx)
	// 查询数据
	var roles model.Roles
	Result, totalCount := roles.FindAll(DB, params)
	TreeResult := roles.ToTree(Result)
	// 返回值
	utils.Success(ctx, "获取成功", gin.H{
		"count": totalCount,
		"data":  TreeResult,
	})
}

// 创建角色
func AddRole(ctx *gin.Context) {
	params, _ := utils.DataMapByRequest(ctx)
	// 获取数据库句柄
	DB := config.GetDB()
	// 写入数据库
	var roles model.Roles
	err := roles.AddRole(DB, params)
	// 返回值
	if err.Error != nil {
		utils.Fail(ctx, err.Error(), nil)
	} else {
		utils.Success(ctx, "创建角色成功", nil)
	}
}

// 编辑角色
func EditRole(ctx *gin.Context) {
	params, _ := utils.DataMapByRequest(ctx)
	// 获取数据库句柄
	DB := config.GetDB()
	// 写入数据库
	var roles model.Roles
	err := roles.EditRole(DB, params)
	// 返回值
	if err.Error != nil {
		utils.Fail(ctx, err.Error(), nil)
	} else {
		utils.Success(ctx, "编辑角色成功", nil)
	}
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
