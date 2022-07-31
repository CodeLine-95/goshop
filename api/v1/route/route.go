package route

import (
	"github.com/gin-gonic/gin"
	"goshop/api/v1/controller"
	"goshop/middleware"
)

func Routers(r *gin.Engine) {
	api := r.Group("/api")
	{
		// v1
		v1 := api.Group("/v1")
		// 登录|注册
		authGroup := v1.Group("/account")
		{
			authGroup.POST("/resgister", controller.Register)
			authGroup.POST("/login", controller.Login)
		}
		// 分类
		categoryGroup := v1.Group("/category", middleware.CheckToken())
		{
			categoryGroup.GET("/lists", controller.GetCategoryLists)
			categoryGroup.GET("/goods-lists", controller.GetCategoryGoodsLists)
		}
		// 商品
		goodsGroup := v1.Group("/goods", middleware.CheckToken())
		{
			goodsGroup.GET("/detail", controller.GoodsDetails)
		}
		// 用户
		userGroup := v1.Group("/user", middleware.CheckToken())
		{
			userGroup.POST("/bind-phone", controller.UserBindPhone)
			userGroup.POST("/un-bind-phone", controller.UnUserBindPhone)
		}

	}
}
