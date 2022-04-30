package model

import (
	"github.com/jinzhu/gorm"
	"goshop/utils"
	"goshop/utils/Paginate"
)

type Roles struct {
	Model
	Name     string     `json:"name" gorm:"size:100;index;default:'';not null;"`    // 角色名称
	Alias    string     `json:"alias" grom:"size:255;index;default:'';not nill;"`   // 别名
	ParentID uint       `json:"parent_id" gorm:"size:10;index;default:0;not null;"` // 父级ID
	Sort     uint       `json:"sort" gorm:"size:1;index;default:0;"`                // 排序值
	Remark   string     `json:"remark" gorm:"size:255;"`                            // 备注
	Status   uint       `json:"status" gorm:"size:1;index;default:0;"`              // 状态（1:启用   2:禁用）
	Children *RoleTrees `json:"children"`
}

// 获取表名
func (Roles) TableName() string {
	return "roles"
}

// 根据检索条件，获取记录行，并获取总记录条数
func (Roles) FindAll(DB *gorm.DB, params map[string]interface{}) ([]*Roles, int64) {
	var Result []*Roles
	page := params["page"].(string)
	pageSize := params["pageSize"].(string)
	ParamsFilter := utils.ParamsFilter(params, "page,pageSize")
	DB.Scopes(Paginate.Paginate(page, pageSize)).Where(ParamsFilter).Order("created_at desc").Find(&Result)
	TotalCount := DB.Find(&Roles{})
	return Result, TotalCount.RowsAffected
}

// RoleTrees 二叉树列表
type RoleTrees []*Roles

// ToTree 转换为树形结构
func (Roles) ToTree(data RoleTrees) RoleTrees {
	// 定义 HashMap 的变量，并初始化
	TreeData := make(map[uint]*Roles)
	// 先重组数据：以数据的ID作为外层的key编号，以便下面进行子树的数据组合
	for _, item := range data {
		TreeData[item.ID] = item
	}
	// 定义 RoleTrees 结构体
	var TreeDataList RoleTrees
	// 开始生成树形
	for _, item := range TreeData {
		// 如果没有根节点树，则为根节点
		if item.ParentID == 0 {
			// 追加到 TreeDataList 结构体中
			TreeDataList = append(TreeDataList, item)
			// 跳过该次循环
			continue
		}
		// 通过 上面的 TreeData HashMap的组合，进行判断是否存在根节点
		// 如果存在根节点，则对应该节点进行处理
		if p_item, ok := TreeData[item.ParentID]; ok {
			// 判断当次循环是否存在子节点，如果没有则作为子节点进行组合
			if p_item.Children == nil {
				// 写入子节点
				children := RoleTrees{item}
				// 插入到 当次结构体的子节点字段中，以指针的方式
				p_item.Children = &children
				// 跳过当前循环
				continue
			}
			// 以指针地址的形式进行追加到结构体中
			*p_item.Children = append(*p_item.Children, item)
		}
	}
	return TreeDataList
}
