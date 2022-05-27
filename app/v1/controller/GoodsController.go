package controller

import (
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

// 获取商品列表信息
func GetGoodsList(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	// 初始化模型
	var goods model.Goods
	// 获取数据库句柄
	DB := config.GetDB()
	// 查询数据
	GoodResult, count := goods.FindAll(DB, params)
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
	// 获取数据库句柄
	DB := config.GetDB()
	// 初始化模型
	var goods model.Goods
	err := goods.AddGoods(DB, params)
	// 返回值
	if err != nil {
		utils.Fail(ctx, err.Error(), nil)
	} else {
		utils.Success(ctx, "创建成功", nil)
	}
}

// 编辑商品
func EditGoods(ctx *gin.Context) {
	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	// 获取数据库句柄
	DB := config.GetDB()
	var goods model.Goods
	err := goods.EditGoods(DB, params)
	// 返回值
	if err != nil {
		utils.Fail(ctx, err.Error(), nil)
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
