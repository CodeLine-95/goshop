package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"goshop/model"
	"goshop/utils/logger"
	"net/url"
)

var DB *gorm.DB

// 初始化db
func InitDB() *gorm.DB {
	driver := viper.GetString("db.driver")
	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	port := viper.GetString("db.port")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	loc := viper.GetString("db.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", user, pass, host, port, dbname, charset, url.QueryEscape(loc))
	db, err := gorm.Open(driver, args)
	if err != nil {
		logger.PanicError(err, "链接数据库错误", true)
	}
	db.AutoMigrate(&model.Admin{})     // 自动创建 User 表
	db.AutoMigrate(&model.Roles{})     // 自动创建 Roles 表
	db.AutoMigrate(&model.Menu{})      // 自动创建 Menu 表
	db.AutoMigrate(&model.RoleMenu{})  // 自动创建 RoleMenu 表
	db.AutoMigrate(&model.Goods{})     // 自动创建 Goods 表
	db.AutoMigrate(&model.RoleRules{}) // 自动创建 RoleRules 表
	db.AutoMigrate(&model.Users{})     // 自动创建 Users 表
	db.AutoMigrate(&model.Category{})  // 自动创建 Category 表
	// db.SingularTable(true)            // 支持单数创建数据表
	db.DB().SetMaxIdleConns(10)  // 用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100) // 用于设置最大打开的连接数，默认值为0表示不限制
	DB = db
	return db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
