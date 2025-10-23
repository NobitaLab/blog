package routes

import (
	"blog/server/controllers"
	"blog/server/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUploadRoutes 注册上传相关路由
func RegisterUploadRoutes(router *gin.Engine, db *gorm.DB) {
	// 创建上传控制器
	uploadController := controllers.NewUploadController(db)

	// 上传路由组
	uploadGroup := router.Group("/api/upload")
	{
		// 图片上传接口
		uploadGroup.POST("/image", middlewares.AuthMiddleware(), uploadController.UploadImage)
		// 图片删除接口
		uploadGroup.POST("/image/delete", middlewares.AuthMiddleware(), uploadController.DeleteImage)
	}
}
