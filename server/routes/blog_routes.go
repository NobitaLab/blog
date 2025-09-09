package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/controllers"
)

// RegisterBlogRoutes 注册博客相关路由
func RegisterBlogRoutes(router *gin.Engine, db *gorm.DB) {
	// 创建博客控制器
	blogController := controllers.NewBlogController(db)

	// 博客路由组
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.GET("", blogController.GetAllBlogs)     // 获取所有博客文章
		blogRoutes.GET("/:id", blogController.GetBlogByID)  // 获取单篇博客文章
		blogRoutes.POST("", blogController.CreateBlog)      // 创建新的博客文章
		blogRoutes.PUT("/:id", blogController.UpdateBlog)   // 更新博客文章
		blogRoutes.DELETE("/:id", blogController.DeleteBlog) // 删除博客文章
	}
}