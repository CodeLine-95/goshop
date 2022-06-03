package model

import "github.com/jinzhu/gorm"

type RoleRules struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	RoleID  uint   `json:"role_id" gorm:"size:10;index;default:'';not null;"`   // 角色编号
	RuleStr string `json:"rule_str" grom:"size:255;index;default:'';not nill;"` // 权限字符串编号
}

// 添加权限规则
func (RoleRules) AddRoleRule(DB *gorm.DB, params map[string]any) error {
	roleId := params["roleId"].(uint)
	ruleStr := params["ruleStr"].(string)
	roleRule := RoleRules{
		RoleID:  roleId,
		RuleStr: ruleStr,
	}

	// 写入数据库
	result := DB.Create(&roleRule)
	// 返回值
	return result.Error
}

// 修改权限规则
func (RoleRules) EditRoleRule(DB *gorm.DB, params map[string]any) error {
	id := params["id"].(uint)
	roleId := params["roleId"].(uint)
	ruleStr := params["ruleStr"].(string)
	var roleRules RoleRules
	// 查询当前数据
	DB.First(&roleRules, id)
	roleRules.RoleID = roleId
	roleRules.RuleStr = ruleStr
	result := DB.Save(&roleRules)
	// 返回值
	return result.Error
}
