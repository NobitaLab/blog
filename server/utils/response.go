package utils

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse 返回成功响应
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

// ErrorResponse 返回错误响应
func ErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

// SuccessListResponse 返回成功的列表响应（带分页）
func SuccessListResponse(c *gin.Context, statusCode int, data interface{}, total int64, page int, limit int) {
	c.JSON(statusCode, gin.H{
		"data":  data,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}