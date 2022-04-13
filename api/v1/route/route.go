package route

import (
	"github.com/gin-gonic/gin"
	"goshop/app/v1/controller"
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

		// 用户信息
		userGroup := v1.Group("/user", middleware.CheckToken())
		{
			userGroup.POST("/userinfo", controller.UserInfo)
		}
	}
}
