package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/controllers"
	"blog/server/middlewares"
)

// RegisterUserRoutes 注册用户相关路由
func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	// 创建用户控制器
	userController := controllers.NewUserController(db)

	// 用户路由组
	userRoutes := router.Group("/api/users")
	{
		// 无需认证的路由
		userRoutes.POST("/register", userController.Register) // 用户注册
		userRoutes.POST("/login", userController.Login)       // 用户登录

		// 需要认证的路由
		auth := userRoutes.Group("")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.POST("/logout", userController.Logout)       // 用户注销
			auth.GET("/me", userController.GetCurrentUser)    // 获取当前用户信息
		}
	}

	// 管理员路由组 - 仅管理员可访问
	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(middlewares.AuthMiddleware())
	adminRoutes.Use(middlewares.AdminMiddleware())
	{
		// 用户管理
		adminRoutes.GET("/users", userController.GetAllUsers)        // 获取所有用户
		adminRoutes.PUT("/users/:id", userController.UpdateUserInfo) // 更新用户信息
	}
}