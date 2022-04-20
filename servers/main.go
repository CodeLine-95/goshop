package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goshop/config"
	"goshop/routes"
	"goshop/utils/logger"
	"os"
)

// 初始化配置文件
func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		logger.PanicError(err, "读取YML配置", true)
	}
}

func initGin() {
	// 初始化gin框架
	r := routes.Init()
	// 读取yml配置文件
	port := viper.GetString("server.port")
	if port != "" {
		err := r.Run(":" + port)
		if err != nil {
			logger.PanicError(err, "Service startup failed！", true)
		}
	}

	err := r.Run()
	if err != nil {
		logger.PanicError(err, "Service startup failed！", true)
	}
}

func main() {
	initConfig()          // 初始化配置
	db := config.InitDB() // 初始化数据库
	defer db.Close()
	initGin() //初始化Gin框架并启动
	config.InitCasbin()
}
