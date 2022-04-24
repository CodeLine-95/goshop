package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	Username string    `gorm:"type:varchar(20);not null"`        // 用户名
	Avatar   string    `gorm:"type:varchar(255);not null"`       // 头像
	Usernick string    `gorm:"type:varchar(50);not null"`        // 昵称
	Phone    string    `gorm:"type:varchar(11);not null;unique"` // 手机号
	Password string    `gorm:"size:255;not null"`                // 密码
	Loginip  string    `gorm:"type:varchar(20);not null"`        // 登录IP
	Email    string    `gorm:"size:255;not null"`                // 邮箱
	group    string    `gorm:"size:255;not null"`                // 角色组 1,2,3,4
	LoginAt  time.Time // 登录时间
}
