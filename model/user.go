package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`        // 用户名
	Usernick string `gorm:"type:varchar(50);not null"`        // 昵称
	Phone    string `gorm:"type:varchar(11);not null;unique"` // 手机号
	Password string `gorm:"size:255;not null"`                // 密码
	Loginip  string `gorm:"type:varchar(20);not null"`        // 登录IP
	Email    string `gorm:"size:255;not null"`                // 邮箱
	Isadmin  uint   `gorm:"type:tinyint(1)"`                  // 是否管理员
}
