package schema

import (
	"fmt"
	"time"
)

// Role 角色对象
type Role struct {
	ID        uint      `json:"id,string"`                             // 唯一标识
	Name      string    `json:"name" binding:"required"`               // 角色名称
	Sequence  int       `json:"sequence"`                              // 排序值
	Memo      string    `json:"memo"`                                  // 备注
	Status    int       `json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:禁用)
	Creator   uint      `json:"creator"`                               // 创建者
	CreatedAt time.Time `json:"created_at"`                            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                            // 更新时间
	RoleMenus RoleMenus `json:"role_menus" binding:"required,gt=0"`    // 角色菜单列表
}

// Roles 角色对象列表
type Roles []*Role

// ToNames 获取角色名称列表
func (a Roles) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Roles) ToMap() map[uint]*Role {
	m := make(map[uint]*Role)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// ----------------------------------------RoleMenu--------------------------------------

// RoleMenu 角色菜单对象
type RoleMenu struct {
	ID       uint `json:"id,string"`                           // 唯一标识
	RoleID   uint `json:"role_id,string" binding:"required"`   // 角色ID
	MenuID   uint `json:"menu_id,string" binding:"required"`   // 菜单ID
	ActionID uint `json:"action_id,string" binding:"required"` // 动作ID
}

// RoleMenus 角色菜单列表
type RoleMenus []*RoleMenu

// ToMap 转换为map
func (a RoleMenus) ToMap() map[string]*RoleMenu {
	m := make(map[string]*RoleMenu)
	for _, item := range a {
		m[fmt.Sprintf("%d-%d", item.MenuID, item.ActionID)] = item
	}
	return m
}

// ToRoleIDMap 转换为角色ID映射
func (a RoleMenus) ToRoleIDMap() map[uint]RoleMenus {
	m := make(map[uint]RoleMenus)
	for _, item := range a {
		m[item.RoleID] = append(m[item.RoleID], item)
	}
	return m
}

// ToMenuIDs 转换为菜单ID列表
func (a RoleMenus) ToMenuIDs() []uint {
	var idList []uint
	m := make(map[uint]struct{})

	for _, item := range a {
		if _, ok := m[item.MenuID]; ok {
			continue
		}
		idList = append(idList, item.MenuID)
		m[item.MenuID] = struct{}{}
	}

	return idList
}

// ToActionIDs 转换为动作ID列表
func (a RoleMenus) ToActionIDs() []uint {
	idList := make([]uint, len(a))
	m := make(map[uint]struct{})
	for i, item := range a {
		if _, ok := m[item.ActionID]; ok {
			continue
		}
		idList[i] = item.ActionID
		m[item.ActionID] = struct{}{}
	}
	return idList
}
