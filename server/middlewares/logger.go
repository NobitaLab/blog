package middlewares

import (

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
	

		// 设置请求ID
		// requestID := uuid.New().String()
		// c.Set("requestID", requestID)

		// 处理请求
		c.Next()

		// 结束时间
	
		// 记录日志
		// 这里可以根据需要调整日志格式和内容
		// 实际项目中可能需要使用专门的日志库
		// fmt.Printf("[GIN] %s - [%s] %s %s %d %s\n",
		// 	requestID,
		// 	t.Format(time.RFC3339),
		// 	c.Request.Method,
		// 	c.Request.RequestURI,
		// 	c.Writer.Status(),
		// 	duration,
		// )
	}
}