package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"strconv"
)

// 获取商品列表信息
func GetGoodsList(ctx *gin.Context) {
	DB := config.InitDB()

	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	page, _ := strconv.Atoi(params["page"])
	pageSize, _ := strconv.Atoi(params["pageSize"])
	Offset := (page - 1) * pageSize
	// 查询数据
	var goods model.Goods
	DB.Limit(page).Offset(Offset).First(&goods)

	fmt.Println(goods)

	// 返回值
	utils.Success(ctx, "获取成功", nil)
}

// 插入商品
func AddGoods(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	GoodsCate, _ := strconv.ParseUint(params["GoodsCate"], 0, 64)
	UnitPrice, _ := strconv.ParseFloat(params["UnitPrice"], 64)
	FavorablePrice, _ := strconv.ParseFloat(params["FavorablePrice"], 64)
	GoodsStock, _ := strconv.ParseUint(params["GoodsStock"], 0, 64)
	GoodsStatus, _ := strconv.ParseUint(params["GoodsStatus"], 0, 64)
	// 拼装数据
	goods := model.Goods{
		GoodsCate:      GoodsCate,
		GoodsName:      params["GoodsName"],
		GoodsProperty:  params["GoodsProperty"],
		GoodsDesc:      params["GoodsDesc"],
		GoodsContent:   params["GoodsContent"],
		UnitPrice:      UnitPrice,
		FavorablePrice: FavorablePrice,
		GoodsStock:     GoodsStock,
		GoodsCover:     params["GoodsCover"],
		GoodsSlides:    params["GoodsSlides"],
		GoodsStatus:    GoodsStatus,
	}
	// 获取数据库句柄
	DB := config.GetDB()
	// 写入数据库
	result := DB.Create(&goods)
	// 返回值
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "创建成功", nil)
	}
}
