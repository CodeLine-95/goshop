package controller

//// 获取用户信息
//func UserInfo(ctx *gin.Context) {
//	user, _ := ctx.Get("user")
//	UserResultData := make(map[string]interface{})
//	userModel := user.(model.Admin)
//	UserResultData["username"] = userModel.Username
//	UserResultData["usernick"] = userModel.Usernick
//	UserResultData["phone"] = userModel.Phone
//	UserResultData["email"] = userModel.Email
//	UserResultData["loginip"] = userModel.Loginip
//	UserResultData["logintime"] = userModel.LoginAt
//	// 返回值
//	utils.Success(ctx, "获取成功", UserResultData)
//}
