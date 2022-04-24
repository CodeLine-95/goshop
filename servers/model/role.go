package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name     string `gorm:"size:100;index;default:'';not null;"` // 角色名称
	Alias    string `grom:"size:255;index;default:'';not nill;"` // 别名
	ParentID uint   `gorm:"size:10;index;default:0;not null;"`   // 父级ID
	Sort     uint   `gorm:"size:1;index;default:0;"`             // 排序值
	Remark   string `gorm:"size:255;"`                           // 备注
	Status   uint   `gorm:"size:1;index;default:0;"`             // 状态（1:启用   2:禁用）
}
