package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goshop/common"
	"goshop/routes"
	"goshop/v1/route"
	"os"
)

// 初始化配置文件
func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to load configuration file")
	}
}

func initGin() {
	// 加载多路由
	routes.Include(route.Routers)
	// 初始化gin框架
	r := routes.Init()
	// 读取yml配置文件
	port := viper.GetString("server.port")
	if port != "" {
		err := r.Run(":" + port)
		if err != nil {
			fmt.Println("Service startup failed ！err: %v \n", err)
		}
	}

	err := r.Run()
	if err != nil {
		fmt.Println("Service startup failed ! err: %v \n", err)
	}
}

func main() {
	initConfig()          // 初始化配置
	db := common.InitDB() // 初始化数据库
	defer db.Close()
	initGin() //初始化Gin框架并启动
}
