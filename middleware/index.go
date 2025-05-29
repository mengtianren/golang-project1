package middleware

import (
	"net/http"
	"project1/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 需要跳过校验的路径
var noAuthPaths = []string{
	"/login",
	"/register",
	"/public",
}

// 检查当前请求路径是否在不需要校验的列表中
func isNoAuthPath(path string) bool {
	for _, p := range noAuthPaths {
		if strings.HasPrefix(path, p) { // 允许/public/xxx 这种情况
			return true
		}
	}
	return false
}

// JWT中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果是免认证路径，直接放行
		if isNoAuthPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 从请求头拿 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ResponseError(c, http.StatusUnauthorized, "请求头缺少Authorization")
			c.Abort()
			return
		}

		// 可能带了 Bearer 前缀，处理掉
		if fl := strings.HasPrefix(tokenString, "Bearer "); fl {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// 解析 token
		user, err := utils.GetJWT(tokenString)
		if err != nil {
			utils.ResponseError(c, http.StatusUnauthorized, "无效的Token，请重新登录")
		}
		c.Set("user", user) // 将用户信息存入上下文中，供后续处理使用
		// token校验通过，继续执行
		c.Next()
	}
}
