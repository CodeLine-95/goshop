package model

import "github.com/jinzhu/gorm"

type RoleMenu struct {
	gorm.Model
	RoleID   uint `gorm:"index;not null;"` // 角色ID
	MenuID   uint `gorm:"index;not null;"` // 菜单ID
	ActionID uint `gorm:"index;not null;"` // 动作ID
}
