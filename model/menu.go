package model

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name       string `gorm:"size:50;index;default:'';not null;"` // 菜单名称
	Icon       string `gorm:"size:255;"`                          // 菜单图标
	Router     string `gorm:"size:255;"`                          // 访问路由
	ParentID   uint   `gorm:"index;default:0;"`                   // 父级内码
	ParentPath string `gorm:"size:255;index;default:'';"`         // 父级路径
	IsShow     int    `gorm:"index;default:0;'"`                  // 是否显示(1:显示  2:隐藏)
	Status     int    `gorm:"index;default:0;"`                   // 状态(1:启用  2:禁用)
	Sequence   int    `gorm:"index;default:0;"`                   // 排序值
	Remark     string `gorm:"size:255;"`                          // 备注
}
