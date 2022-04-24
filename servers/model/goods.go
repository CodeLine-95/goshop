package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 模型结构体
type Goods struct {
	gorm.Model
	GoodsCate      uint64  `gorm:"size:10;not null;default:0;comment:'商品类别'"`                        // 商品类别
	GoodsName      string  `gorm:"size:255;not null;default:'';index:GoodsNameIndex;comment:'商品名称'"` // 商品名称
	GoodsProperty  string  `gorm:"size:255;not null;default:'';index:GoodsProperty;comment:'商品属性'"`  // 商品属性
	GoodsDesc      string  `gorm:"size:255;not null;default:'';index:GoodsDesc;comment:'商品单价'"`      // 商品简介
	GoodsContent   string  `gorm:"longtext;not null;default:'';comment:'商品信息'"`                      // 商品信息
	UnitPrice      float64 `gorm:"decimal(18,2);not null;default:0;comment:'商品单价'"`                  // 商品单价
	FavorablePrice float64 `gorm:"decimal(18,2);not null;default:0;comment:'优惠价格'"`                  // 优惠价格
	GoodsStock     uint64  `gorm:"size:10;not null;default:0;comment:'商品库存'"`                        // 商品库存
	GoodsCover     string  `gorm:"size:255;not null;default:'';comment:'商品封面图'"`                     // 商品封面图
	GoodsSlides    string  `gorm:"size:255;not null;default:'';comment:'商品幻灯片'"`                     // 商品幻灯片
	GoodsStatus    uint64  `gorm:"size:10;not null;default:0;comment:'商品状态'"`                        // 商品状态
}

// 获取表名
func (Goods) TableName() string {
	return "goods"
}

// 数据返回结果集 json
type GoodResult struct {
	ID             uint       `json:"id"`              // 商品编号
	GoodsCate      uint64     `json:"goods_cate"`      // 商品类别
	GoodsName      string     `json:"goods_name"`      // 商品名称
	GoodsProperty  string     `json:"goods_property"`  // 商品属性
	GoodsDesc      string     `json:"goods_desc"`      // 商品简介
	GoodsContent   string     `json:"goods_content"`   // 商品信息
	UnitPrice      float64    `json:"unit_price"`      // 商品单价
	FavorablePrice float64    `json:"favorable_price"` // 优惠价格
	GoodsStock     uint64     `json:"goods_stock"`     // 商品库存
	GoodsCover     string     `json:"goods_cover"`     // 商品封面图
	GoodsSlides    string     `json:"goods_slides"`    // 商品幻灯片
	GoodsStatus    uint64     `json:"goods_status"`    // 商品状态
	CreatedAt      time.Time  `json:"created_at"`      // 创建时间
	UpdatedAt      time.Time  `json:"updated_at"`      // 更新时间
	DeletedAt      *time.Time `json:"deleted_at"`      // 软删除时间
}
