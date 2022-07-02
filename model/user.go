package model

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UserID   string `sql:"index" gorm:"type:varchar(255);not null;unique"`                            // 用户唯一标识
	Username string `form:"username" binding:"required" sql:"index" gorm:"type:varchar(20);not null"` // 用户名
	Avatar   string `form:"avatar" gorm:"type:varchar(255);not null"`                                 // 头像
	Usernick string `form:"usernick" gorm:"type:varchar(50);not null"`                                // 昵称
	Phone    string `form:"phone" sql:"index" gorm:"type:varchar(11);not null;unique"`                // 手机号
	Password string `form:"password" binding:"required" gorm:"size:255;not null"`                     // 密码
	Loginip  string `gorm:"type:varchar(20);not null"`                                                // 登录IP
	Email    string `form:"email" sql:"index" gorm:"size:255;not null"`                               // 邮箱
}
