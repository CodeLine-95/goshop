package route

import (
	"github.com/gin-gonic/gin"
	"goshop/middleware"
	"goshop/v1/controller"
)

func Routers(r *gin.Engine) {
	Api_v1 := r.Group("/v1")
	{
		// 登录|注册
		authGroup := Api_v1.Group("/account")
		{
			authGroup.POST("/resgister", controller.Register)
			authGroup.POST("/login", controller.Login)
		}

		// 用户信息
		userGroup := Api_v1.Group("/user", middleware.CheckToken())
		{
			userGroup.POST("/userinfo", controller.UserInfo)
		}
	}
}
