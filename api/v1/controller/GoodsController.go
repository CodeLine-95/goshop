package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"strings"
)

type GoodsDetailsParamsRequest struct {
	ID int64 `form:"id" json:"id"`
}

type GoodsDetailsResult struct {
	model.Goods
	CateName          string   `json:"cate_name"`         // 分类名称
	GoodsPropertyList []string `json:"goodsPropertyList"` // 商品属性列表
}

// GoodsDetails 商品详情
func GoodsDetails(ctx *gin.Context) {
	var params GoodsDetailsParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}
	var GoodDetail GoodsDetailsResult
	DB := config.GetDB()
	resErr := DB.Model(&model.Goods{}).Select("goods.*, categories.name as cate_name").Joins("left join categories on categories.id = goods.goods_cate").Where("goods.id = ?", params.ID).Scan(&GoodDetail).Error
	if resErr != nil {
		utils.Fail(ctx, resErr.Error(), nil)
		return
	}

	GoodDetail.GoodsPropertyList = strings.Split(GoodDetail.GoodsProperty, ";")

	utils.Success(ctx, "获取成功", GoodDetail)
	return
}
