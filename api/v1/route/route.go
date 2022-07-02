package route

import (
	"github.com/gin-gonic/gin"
	"goshop/api/v1/controller"
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
		categoryGroup := v1.Group("/category")
		{
			categoryGroup.GET("/get-lists", controller.GetCategoryLists)
		}

	}
}
