package route

import (
	"github.com/gin-gonic/gin"
	"goshop/app/v1/controller"
	"goshop/middleware"
)

func Routers(r *gin.Engine) {
	app := r.Group("/app")
	{
		// v1
		v1 := app.Group("/v1")
		{
			// 登录|注册
			authGroup := v1.Group("/account")
			{
				authGroup.POST("/resgister", controller.Register)
				authGroup.POST("/login", controller.Login)
			}
			// 管理平台
			adminGroup := v1.Group("/admin", middleware.CheckToken())
			{
				// 用户信息
				userGroup := adminGroup.Group("/user")
				{
					userGroup.POST("/get-info", controller.UserInfo)
				}
				// 商品
				goodsGroup := adminGroup.Group("/goods")
				{
					goodsGroup.POST("/get-list", controller.GetGoodsList)
					goodsGroup.POST("/add", controller.AddGoods)
					goodsGroup.PUT("/edit", controller.EditGoods)
					goodsGroup.PUT("/is-uper-lower", controller.IsUperOrLower)
				}
				// 角色
				roleGroup := adminGroup.Group("/role")
				{
					roleGroup.POST("/add", controller.AddRole)
				}
			}

		}
	}
}
