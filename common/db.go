package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"goshop/model"
	"net/url"
)

var DB *gorm.DB // 全局db句柄
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
		fmt.Println("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
