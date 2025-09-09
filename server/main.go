package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/config"
	"blog/server/models"
	"blog/server/routes"
	"blog/server/middlewares"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移数据库表结构
	err = db.AutoMigrate(&models.Blog{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 添加一些测试数据
	addTestData(db)

	// 创建 Gin 路由器
	router := gin.Default()

	// 注册中间件
	router.Use(middlewares.LoggerMiddleware())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// 注册路由
	routes.RegisterBlogRoutes(router, db)

	// 启动服务器
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// corsMiddleware 处理跨域请求
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// addTestData 添加测试数据
func addTestData(db *gorm.DB) {
	// 检查是否已有测试数据
	var count int64
	db.Model(&models.Blog{}).Count(&count)
	if count > 0 {
		return
	}

	// 添加测试数据
	blogs := []models.Blog{
		{
			Title:   "Go语言入门教程",
			Content: "# Go语言入门教程\n\n## 什么是Go语言\nGo是一种开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。\n\n## 为什么学习Go\n- 高性能\n- 并发支持\n- 简洁的语法\n- 强大的标准库",
			Author: "张三",
		},
		{
			Title:   "Vue.js 3.0新特性",
			Content: "# Vue.js 3.0新特性\n\n## Composition API\nComposition API是Vue 3.0中引入的新API，它提供了一种更灵活的方式来组织和重用组件逻辑。\n\n## Teleport\nTeleport允许我们将组件的内容渲染到DOM中的任何位置，非常适合模态框等场景。",
			Author: "李四",
		},
	}

	for _, blog := range blogs {
		db.Create(&blog)
	}
}