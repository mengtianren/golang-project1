package router

import (
	"fmt"
	"project1/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	router := gin.Default()
	router.Use(middleware.JWTAuthMiddleware()) //认证

	InitPubicRoute(router)
	InitUserRoute(router)
	InitAdminRoute(router)
	InitDictRoute(router)
	fmt.Println("路由初始化成功")

	router.Run(":8088")
}
