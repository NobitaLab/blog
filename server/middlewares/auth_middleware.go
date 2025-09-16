package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"blog/server/utils"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头部
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "缺少Authorization头部")
		c.Abort()
		return
	}

	// 检查token格式
	tokenParts := strings.SplitN(authHeader, " ", 2)
	if !(len(tokenParts) == 2 && tokenParts[0] == "Bearer") {
		utils.ErrorResponse(c, http.StatusUnauthorized, "无效的Authorization格式")
		c.Abort()
		return
	}

	// 解析token
	claims, err := utils.ParseToken(tokenParts[1])
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "无效的令牌")
		c.Abort()
		return
	}

	// 将用户信息存入上下文
	c.Set("userID", claims.UserID)
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)

	// 继续处理请求
	c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.ErrorResponse(c, http.StatusForbidden, "需要管理员权限")
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next()
	}
}