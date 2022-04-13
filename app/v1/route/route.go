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
