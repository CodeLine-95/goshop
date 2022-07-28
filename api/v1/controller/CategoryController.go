package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

type ParamsRequest struct {
	Name string `form:"name" json:"name"`
	Pid  int64  `form:"pid" json:"pid"` // pid 上级节点ID
}

// GetCategoryLists 获取分类列表
func GetCategoryLists(ctx *gin.Context) {
	var params ParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}
	// 过滤page和pageSize
	paramsMap, _ := utils.AnyToMap(params)
	ParamsFilter := utils.ParamsFilter("page,pageSize,name", paramsMap)
	// 获取列表
	DB := config.GetDB()
	var Result []*model.Category
	var resErr error
	// 如果name条件不为空，追加模糊查询：position('搜索字符' in 字段)
	if len(params.Name) > 0 {
		resErr = DB.Where(ParamsFilter).Where("position(? in name)", params.Name).Find(&Result).Error
	} else {
		resErr = DB.Where(ParamsFilter).Find(&Result).Error
	}
	if resErr != nil {
		utils.Fail(ctx, resErr.Error(), nil)
		return
	}
	utils.Success(ctx, "获取成功", Result)
}
