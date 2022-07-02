package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string         `form:"name" binding:"required" json:"name" gorm:"varchar(255);not null;default:'';comment:'分类名称'"` // 分类名称
	Pid      int64          `form:"pid" json:"pid" gorm:"size:11;not null;default:0;comment:'分类节点'"`                            // 分类节点：0根节点
	Icon     string         `form:"icon" json:"icon" gorm:"varchar(255);not null;default:'';comment:'分类图标'"`                    // 分类图标
	State    int64          `form:"state" json:"state" gorm:"size:1;not null;default:0;comment:'分类状态：0未启用，1已启用'"`               // 分类状态：0未启用，1已启用
	Sort     int64          `form:"sort" json:"sort" gorm:"size:11;not null;default:0;comment:'分类排序'"`                          // 分类排序
	Tag      string         `form:"tag" json:"tag" gorm:"varchar(255);not null;default:0;comment:'分类标签'"`                       // 分类标签
	Children *CategoryTrees `json:"children"`
}

// CategoryTrees 二叉树列表
type CategoryTrees []*Category

// ToTree 转换为树形结构
func (Category) ToTree(data CategoryTrees) CategoryTrees {
	// 定义 HashMap 的变量，并初始化
	TreeData := make(map[int64]*Category)
	// 先重组数据：以数据的ID作为外层的key编号，以便下面进行子树的数据组合
	for _, item := range data {
		TreeData[int64(item.ID)] = item
	}
	// 定义 RoleTrees 结构体
	var TreeDataList CategoryTrees
	// 开始生成树形
	for _, item := range TreeData {
		// 如果没有根节点树，则为根节点
		if item.Pid == 0 {
			// 追加到 TreeDataList 结构体中
			TreeDataList = append(TreeDataList, item)
			// 跳过该次循环
			continue
		}
		// 通过 上面的 TreeData HashMap的组合，进行判断是否存在根节点
		// 如果存在根节点，则对应该节点进行处理
		if pItem, ok := TreeData[item.Pid]; ok {
			// 判断当次循环是否存在子节点，如果没有则作为子节点进行组合
			if pItem.Children == nil {
				// 写入子节点
				children := CategoryTrees{item}
				// 插入到 当次结构体的子节点字段中，以指针的方式
				pItem.Children = &children
				// 跳过当前循环
				continue
			}
			// 以指针地址的形式进行追加到结构体中
			*pItem.Children = append(*pItem.Children, item)
		}
	}
	return TreeDataList
}
