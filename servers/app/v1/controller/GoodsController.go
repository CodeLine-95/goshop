package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
	"goshop/utils/Paginate"
	"strconv"
)

// 获取商品列表信息
func GetGoodsList(ctx *gin.Context) {
	DB := config.InitDB()

	// 获取参数
	//params, _ := utils.DataMapByRequest(ctx)
	// 查询数据
	var goods model.Goods
	var GoodResult []model.GoodResult
	DB.Table(goods.TableName()).Scopes(Paginate.Paginate(ctx)).Order("created_at desc").Find(&GoodResult)
	// struct 转 map  (反射 reflect包)
	//data := make(map[string]interface{})
	//elem := reflect.ValueOf(&goods).Elem()
	//var relType reflect.Type
	//for i := 0; i < relType.NumField(); i++ {
	//	data[relType.Field(i).Name] = elem.Field(i).Interface()
	//}
	// 返回值
	utils.Success(ctx, "获取成功", gin.H{
		"coount": len(GoodResult),
		"data":   GoodResult,
	})
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

// 编辑商品
func EditGoods(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	int_id, _ := strconv.Atoi(params["id"])
	GoodsCate, _ := strconv.ParseUint(params["GoodsCate"], 0, 64)
	UnitPrice, _ := strconv.ParseFloat(params["UnitPrice"], 64)
	FavorablePrice, _ := strconv.ParseFloat(params["FavorablePrice"], 64)
	GoodsStock, _ := strconv.ParseUint(params["GoodsStock"], 0, 64)
	GoodsStatus, _ := strconv.ParseUint(params["GoodsStatus"], 0, 64)
	// 获取数据库句柄
	DB := config.GetDB()
	var goods model.Goods
	// 查询当前数据
	DB.First(&goods, uint(int_id))
	// 更新数据
	goods.GoodsCate = GoodsCate
	goods.GoodsName = params["GoodsName"]
	goods.GoodsProperty = params["GoodsProperty"]
	goods.GoodsDesc = params["GoodsDesc"]
	goods.GoodsContent = params["GoodsContent"]
	goods.UnitPrice = UnitPrice
	goods.FavorablePrice = FavorablePrice
	goods.GoodsStock = GoodsStock
	goods.GoodsCover = params["GoodsCover"]
	goods.GoodsSlides = params["GoodsSlides"]
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
	int_id, _ := strconv.Atoi(params["id"])
	GoodsStatus, _ := strconv.ParseUint(params["GoodsStatus"], 0, 64)
	// 获取数据库句柄
	DB := config.GetDB()
	var goods model.Goods
	// 查询当前数据
	DB.First(&goods, uint(int_id))
	goods.GoodsStatus = GoodsStatus
	result := DB.Save(&goods)
	// 返回值
	if result.Error != nil {
		utils.Fail(ctx, result.Error.Error(), nil)
	} else {
		utils.Success(ctx, "更新成功", nil)
	}
}
