package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"goshop/utils/Paginate"
	"strconv"
)

type ParamsRequest struct {
	Name string `form:"name" json:"name"`
	Pid  int64  `form:"pid" json:"pid"` // pid 上级节点ID
}

type GoodsParamsRequest struct {
	CateId   int64  `form:"cateId" json:"cateId"`
	Page     string `form:"page" json:"page"`
	PageSize string `form:"pageSize" json:"pageSize"`
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

// GetCategoryGoodsLists 获取分类下的全部商品
func GetCategoryGoodsLists(ctx *gin.Context) {
	var params GoodsParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	var cate model.Category
	DB := config.GetDB()
	pids := cate.GetCategoryIds(DB, params.CateId)

	var GoodsResult []*model.Goods
	var resErr error
	if len(pids) > 0 {
		pidsByte := new(bytes.Buffer)
		for _, value := range pids {
			_, err := fmt.Fprintf(pidsByte, "'%s',", strconv.FormatInt(value, 10))
			if err != nil {
				return
			}
		}
		pidsString := pidsByte.String()
		pidsString = pidsString[0 : len(pidsString)-2]
		pidsString = pidsString[1:]
		resErr = DB.Scopes(Paginate.Paginate(params.Page, params.PageSize)).Where("goods_cate in (?)", pidsString).Order("created_at desc").Find(&GoodsResult).Error
	} else {
		resErr = DB.Scopes(Paginate.Paginate(params.Page, params.PageSize)).Order("created_at desc").Find(&GoodsResult).Error
	}

	count := DB.Find(&model.Goods{}).RowsAffected

	if resErr != nil {
		utils.Fail(ctx, resErr.Error(), nil)
		return
	}
	utils.Success(ctx, "获取成功", gin.H{
		"count": count,
		"data":  GoodsResult,
	})

}
