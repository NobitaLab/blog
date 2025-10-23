package main

import (
	"log"
	"net/http"
	"time"

	"blog/server/config"
	"blog/server/middlewares"
	"blog/server/models"
	"blog/server/routes"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	err = db.AutoMigrate(&models.Blog{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化管理员账户
	initAdminAccount(db)

	// 创建 Gin 路由器
	router := gin.Default()

	// 注册中间件
	router.Use(middlewares.LoggerMiddleware())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// 静态文件服务 - 提供上传的图片访问
	router.Static("/uploads", "./uploads")

	// 注册路由
	routes.RegisterBlogRoutes(router, db)
	routes.RegisterUserRoutes(router, db)
	routes.RegisterUploadRoutes(router, db)

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

// initAdminAccount 初始化管理员账户
func initAdminAccount(db *gorm.DB) {
	// 检查管理员账户是否已存在
	var admin models.User
	result := db.Where("username = ?", "admin").First(&admin)
	if result.Error == nil {
		// 管理员账户已存在，不需要创建
		return
	}

	// 创建管理员账户
	hashedPassword, err := utils.HashPassword("admin@123") // 默认密码，实际项目中应改为更安全的密码
	if err != nil {
		log.Printf("Failed to hash admin password: %v", err)
		return
	}

	admin = models.User{
		Username: "admin",
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("Failed to create admin account: %v", err)
		return
	}

	log.Println("Admin account created successfully")
}

// corsMiddleware 处理跨域请求 - 安全版本
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限制允许的来源，只允许前端应用的域名
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := []string{
			"http://localhost:5173", // 开发环境前端地址
			"http://localhost:3000", // http-server静态文件服务地址
			// 生产环境可以添加实际的域名
			// "https://your-frontend-domain.com",
		}

		// 检查请求来源是否在允许列表中
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				// 只有在指定具体来源时才能启用凭证支持
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		// 设置允许的头部
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 设置允许的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

/* 旧的CORS中间件实现（安全性较低）
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
*/
