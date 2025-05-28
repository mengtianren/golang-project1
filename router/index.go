package router

import (
	"fmt"
	"project1/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JWTAuthMiddleware()) //认证

	InitPubicRoute(r)
	InitAdminRoute(r)
	InitUserRoute(r)
	InitDictRoute(r)
	InitMenuRoute(r)
	fmt.Println("路由初始化成功")

	return r
}
