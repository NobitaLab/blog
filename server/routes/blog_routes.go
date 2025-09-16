package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/controllers"
	"blog/server/middlewares"
)

// RegisterBlogRoutes 注册博客相关路由
func RegisterBlogRoutes(router *gin.Engine, db *gorm.DB) {
	// 创建博客控制器
	blogController := controllers.NewBlogController(db)

	// 博客路由组 - 公开访问（获取博客列表和详情）
	router.GET("/api/blogs", blogController.GetAllBlogs)     // 获取所有博客文章
	router.GET("/api/blogs/:id", blogController.GetBlogByID)  // 获取单篇博客文章

	// 需要认证的博客路由
	authBlogRoutes := router.Group("/api/blogs")
	authBlogRoutes.Use(middlewares.AuthMiddleware())
	{
		authBlogRoutes.POST("", blogController.CreateBlog)      // 创建新的博客文章
		authBlogRoutes.PUT("/:id", blogController.UpdateBlog)   // 更新博客文章
		authBlogRoutes.DELETE("/:id", blogController.DeleteBlog) // 删除博客文章
	}
}