package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"goshop/config"
	"goshop/model"
	"goshop/utils"
)

// UserBindPhone 绑定手机号
func UserBindPhone(ctx *gin.Context) {
	var params model.BindPhone
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误"+err.Error(), nil)
		return
	}
	var user model.Users
	DB := config.GetDB()
	err := DB.Where("user_id = ?", params.UserID).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "该用户不存在："+err.Error(), nil)
		return
	}
	if len(user.Phone) > 0 {
		utils.Fail(ctx, "已绑定手机号", nil)
		return
	}
	user.Phone = params.Phone
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "绑定失败："+resErr.Error(), nil)
		return
	}
	utils.Success(ctx, "绑定成功", nil)
	return
}

// UnUserBindPhone 手机号解绑
func UnUserBindPhone(ctx *gin.Context) {
	var params model.BindPhone
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误"+err.Error(), nil)
		return
	}
	var user model.Users
	DB := config.GetDB()
	err := DB.Where("user_id = ?", params.UserID).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "该用户不存在："+err.Error(), nil)
		return
	}
	if len(user.Phone) == 0 {
		utils.Fail(ctx, "未绑定手机号", nil)
		return
	}
	if user.Phone != params.Phone {
		utils.Fail(ctx, "手机号不正确", nil)
		return
	}
	var null sql.NullString
	user.Phone = null.String
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "解绑失败："+resErr.Error(), nil)
		return
	}
	utils.Success(ctx, "解绑成功", nil)
	return
}
