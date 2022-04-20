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

func AddGoods(ctx *gin.Context) {
	//DB := config.InitDB()

	// 获取参数
	params, _ := utils.DataMapByRequest(ctx)
	fmt.Println(params)
	// 返回值
	utils.Success(ctx, "获取成功", nil)
}
