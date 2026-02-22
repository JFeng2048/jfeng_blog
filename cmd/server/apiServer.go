package main

import (
	"fmt"
	"os"

	"jfeng_blog/internal/router"
	"jfeng_blog/internal/utils/config"
	"jfeng_blog/internal/utils/logger"
)

// @title jfeng_blog API
// @version 1.0
// @description jfeng_blog API 文档
// @host localhost:8080
// @Schemes http https
// @BasePath /api/v1

func main() {
	fmt.Println("______________________________________________")
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("配置加载失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("______________________________________________")
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("日志初始化失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("______________________________________________")
	if _, err := config.ConnectDB(); err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("______________________________________________")
	router.Run()
	fmt.Println("______________________________________________")
}
