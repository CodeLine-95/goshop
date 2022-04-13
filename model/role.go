package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name     string `gorm:"size:100;index;default:'';not null;"` // 角色名称
	Sequence int    `gorm:"index;default:0;"`                    // 排序值
	Remark   string `gorm:"size:255;"`                           // 备注
	Status   int    `gorm:"index;default:0;"`                    // 状态（1:启用   2:禁用）
}
