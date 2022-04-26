package model

import (
	"github.com/jinzhu/gorm"
	"goshop/utils"
	"goshop/utils/Paginate"
)

// 模型结构体
type Goods struct {
	Model
	GoodsCate      uint64  `json:"goods_cate" gorm:"size:10;not null;default:0;comment:'商品类别'"`                           // 商品类别
	GoodsName      string  `json:"goods_name" gorm:"size:255;not null;default:'';index:GoodsNameIndex;comment:'商品名称'"`    // 商品名称
	GoodsProperty  string  `json:"goods_property" gorm:"size:255;not null;default:'';index:GoodsProperty;comment:'商品属性'"` // 商品属性
	GoodsDesc      string  `json:"goods_desc" gorm:"size:255;not null;default:'';index:GoodsDesc;comment:'商品单价'"`         // 商品简介
	GoodsContent   string  `json:"goods_content" gorm:"longtext;not null;default:'';comment:'商品信息'"`                      // 商品信息
	UnitPrice      float64 `json:"unit_price" gorm:"decimal(18,2);not null;default:0;comment:'商品单价'"`                     // 商品单价
	FavorablePrice float64 `json:"favorable_price" gorm:"decimal(18,2);not null;default:0;comment:'优惠价格'"`                // 优惠价格
	GoodsStock     uint64  `json:"goods_stock" gorm:"size:10;not null;default:0;comment:'商品库存'"`                          // 商品库存
	GoodsCover     string  `json:"goods_cover" gorm:"size:255;not null;default:'';comment:'商品封面图'"`                       // 商品封面图
	GoodsSlides    string  `json:"goods_slides" gorm:"size:255;not null;default:'';comment:'商品幻灯片'"`                      // 商品幻灯片
	GoodsStatus    uint64  `json:"goods_status" gorm:"size:10;not null;default:0;comment:'商品状态'"`                         // 商品状态
	LoginTime      string  `json:"login_time"`                                                                            // 登录时间
}

// 获取表名
func (Goods) TableName() string {
	return "goods"
}

// 根据检索条件，获取记录行，并获取总记录条数
func (Goods) FindAll(DB *gorm.DB, params map[string]interface{}) ([]Goods, int64) {
	var GoodResult []Goods
	page := params["page"].(string)
	pageSize := params["pageSize"].(string)
	ParamsFilter := utils.ParamsFilter(params, "page,pageSize")
	DB.Scopes(Paginate.Paginate(page, pageSize)).Where(ParamsFilter).Order("created_at desc").Find(&GoodResult)
	GoodCount := DB.Find(&Goods{})
	return GoodResult, GoodCount.RowsAffected
}
