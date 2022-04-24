package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"strconv"
)

// 创建角色
func AddRole(ctx *gin.Context) {
	params, _ := utils.DataMapByRequest(ctx)
	ParentID, _ := strconv.Atoi(params["ParentID"])
	Sort, _ := strconv.Atoi(params["Sort"])
	Status, _ := strconv.Atoi(params["Status"])
	role := model.Role{
		Name:     params["Name"],
		Alias:    params["Alias"],
		ParentID: uint(ParentID),
		Sort:     uint(Sort),
		Remark:   params["Remark"],
		Status:   uint(Status),
	}

	// 获取数据库句柄
	DB := config.GetDB()
	// 写入数据库
	result := DB.Create(&role)
	// 返回值
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "创建角色成功", nil)
	}
}
