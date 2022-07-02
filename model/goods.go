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
}

// 获取表名
func (Goods) TableName() string {
	return "goods"
}

// 根据检索条件，获取记录行，并获取总记录条数
func (Goods) FindAll(DB *gorm.DB, params map[string]any) ([]Goods, int64) {
	var GoodResult []Goods
	page := params["page"].(string)
	pageSize := params["pageSize"].(string)
	ParamsFilter := utils.ParamsFilter("page,pageSize", params)
	DB.Scopes(Paginate.Paginate(page, pageSize)).Where(ParamsFilter).Order("created_at desc").Find(&GoodResult)
	GoodCount := DB.Find(&Goods{})
	return GoodResult, GoodCount.RowsAffected
}

// 插入商品操作
func (Goods) AddGoods(DB *gorm.DB, params map[string]any) error {
	GoodsCate, _ := params["GoodsCate"].(uint64)
	UnitPrice, _ := params["UnitPrice"].(float64)
	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)
	GoodsStatus, _ := params["GoodsStatus"].(uint64)
	// 拼装数据
	goods := Goods{
		GoodsCate:      GoodsCate,
		GoodsName:      params["GoodsName"].(string),
		GoodsProperty:  params["GoodsProperty"].(string),
		GoodsDesc:      params["GoodsDesc"].(string),
		GoodsContent:   params["GoodsContent"].(string),
		UnitPrice:      UnitPrice,
		FavorablePrice: FavorablePrice,
		GoodsStock:     GoodsStock,
		GoodsCover:     params["GoodsCover"].(string),
		GoodsSlides:    params["GoodsSlides"].(string),
		GoodsStatus:    GoodsStatus,
	}
	// 写入数据库
	result := DB.Create(&goods)
	return result.Error
}

// 编辑商品
func (Goods) EditGoods(DB *gorm.DB, params map[string]any) error {
	int_id, _ := params["id"].(uint)
	GoodsCate, _ := params["GoodsCate"].(uint64)
	UnitPrice, _ := params["UnitPrice"].(float64)
	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)
	GoodsStatus, _ := params["GoodsStatus"].(uint64)
	var goods Goods
	// 查询当前数据
	DB.First(&goods, int_id)
	// 更新数据
	goods.GoodsCate = GoodsCate
	goods.GoodsName = params["GoodsName"].(string)
	goods.GoodsProperty = params["GoodsProperty"].(string)
	goods.GoodsDesc = params["GoodsDesc"].(string)
	goods.GoodsContent = params["GoodsContent"].(string)
	goods.UnitPrice = UnitPrice
	goods.FavorablePrice = FavorablePrice
	goods.GoodsStock = GoodsStock
	goods.GoodsCover = params["GoodsCover"].(string)
	goods.GoodsSlides = params["GoodsSlides"].(string)
	goods.GoodsStatus = GoodsStatus
	result := DB.Save(&goods)
	return result.Error
}
