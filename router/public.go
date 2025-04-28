package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitPubicRoute(router *gin.Engine) {

	router.POST("/login", controller.User{}.Login)

}
