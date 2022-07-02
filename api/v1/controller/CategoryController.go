package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"goshop/utils/Paginate"
	"strconv"
)

type ParamsRequest struct {
	Page     int64  `form:"page" binding:"required" json:"page"`
	PageSize int64  `form:"pageSize" binding:"required" json:"pageSize"`
	Name     string `form:"name" json:"name"`
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
	ParamsFilter := utils.ParamsFilter("page,pageSize", paramsMap)
	// 获取列表
	DB := config.GetDB()
	var Result []*model.Category
	resErr := DB.Scopes(Paginate.Paginate(strconv.FormatInt(params.Page, 10), strconv.FormatInt(params.PageSize, 10))).Where(ParamsFilter).Find(&Result).Error
	if resErr != nil {
		utils.Fail(ctx, resErr.Error(), nil)
		return
	}
	var category model.Category
	ResultLists := category.ToTree(Result)
	utils.Success(ctx, "获取成功", ResultLists)
}
