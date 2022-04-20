package model

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	GoodsCate      uint    `gorm:"size:10;not null;default:0;commit:'商品类别'"`                        // 商品类别
	GoodsName      string  `gorm:"size:255;not null;default:'';index:GoodsNameIndex;commit:'商品名称'"` // 商品名称
	GoodsProperty  string  `gorm:"size:255;not null;default:'';index:GoodsProperty;commit:'商品属性'"`  // 商品属性
	GoodsDesc      string  `gorm:"size:255;not null;default:'';index:GoodsDesc;commit:'商品单价'"`      // 商品简介
	GoodsContent   string  `gorm:"longtext;not null;default:'';commit:'商品信息'"`                      // 商品信息
	UnitPrice      float64 `gorm:"decimal(18,2);not null;default:0;commit:'商品单价'"`                  // 商品单价
	FavorablePrice float64 `gorm:"decimal(18,2);not null;default:0;commit:'优惠价格'"`                  // 优惠价格
	GoodsStock     uint    `gorm:"size:10;not null;default:0;commit:'商品库存'"`                        // 商品库存
	GoodsCover     string  `gorm:"size:255;not null;default:'';commit:'商品封面图'"`                     // 商品封面图
	GoodsSlides    string  `gorm:"size:255;not null;default:'';commit:'商品幻灯片'"`                     // 商品幻灯片
	GoodsStatus    uint    `gorm:"size:10;not null;default:0;commit:'商品状态'"`                        // 商品状态
}
