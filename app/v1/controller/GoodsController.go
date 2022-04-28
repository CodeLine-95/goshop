package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

// 获取商品列表信息
func GetGoodsList(ctx *gin.Context) {
	DB := config.InitDB()
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	// 查询数据
	var goods model.Goods
	GoodResult, count := goods.FindAll(DB, params)
	// struct 转 map  (反射 reflect包)
	//data := make(map[string]interface{})
	//elem := reflect.ValueOf(&goods).Elem()
	//var relType reflect.Type
	//for i := 0; i < relType.NumField(); i++ {
	//	data[relType.Field(i).Name] = elem.Field(i).Interface()
	//}
	// 返回值
	utils.Success(ctx, "获取成功", gin.H{
		"count": count,
		"data":  GoodResult,
	})
}

// 插入商品
func AddGoods(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	GoodsCate, _ := params["GoodsCate"].(uint64)
	UnitPrice, _ := params["UnitPrice"].(float64)
	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)
	GoodsStatus, _ := params["GoodsStatus"].(uint64)
	// 拼装数据
	goods := model.Goods{
		GoodsCate:      GoodsCate,
		GoodsName:      params["GoodsName"].(string),
		GoodsProperty:  params["GoodsProperty"].(string),
		GoodsDesc:      params["GoodsDesc"].(string),
		GoodsContent:   params["GoodsContent"].(string),
		UnitPrice:      UnitPrice,
		FavorablePrice: FavorablePrice,
		GoodsStock:     GoodsStock,
		GoodsCover:     params["GoodsCover"].(string),
		GoodsSlides:    params["GoodsSlides"].(string),
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

// 编辑商品
func EditGoods(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	int_id, _ := params["id"].(uint)
	GoodsCate, _ := params["GoodsCate"].(uint64)
	UnitPrice, _ := params["UnitPrice"].(float64)
	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)
	GoodsStatus, _ := params["GoodsStatus"].(uint64)
	// 获取数据库句柄
	DB := config.GetDB()
	var goods model.Goods
	// 查询当前数据
	DB.First(&goods, int_id)
	// 更新数据
	goods.GoodsCate = GoodsCate
	goods.GoodsName = params["GoodsName"].(string)
	goods.GoodsProperty = params["GoodsProperty"].(string)
	goods.GoodsDesc = params["GoodsDesc"].(string)
	goods.GoodsContent = params["GoodsContent"].(string)
	goods.UnitPrice = UnitPrice
	goods.FavorablePrice = FavorablePrice
	goods.GoodsStock = GoodsStock
	goods.GoodsCover = params["GoodsCover"].(string)
	goods.GoodsSlides = params["GoodsSlides"].(string)
	goods.GoodsStatus = GoodsStatus
	result := DB.Save(&goods)
	// 返回值
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "更新成功", nil)
	}
}

// 是否上下架
func IsUperOrLower(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	int_id, _ := params["id"].(uint)
	GoodsStatus, _ := params["GoodsStatus"].(uint64)
	// 获取数据库句柄
	DB := config.GetDB()
	var goods model.Goods
	// 查询当前数据
	DB.First(&goods, int_id)
	goods.GoodsStatus = GoodsStatus
	result := DB.Save(&goods)
	// 返回值
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "上架成功", nil)
	}
}
