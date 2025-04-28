package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(r *gin.Engine) {
	router := r.Group("/user")
	{
		router.GET("/get", controller.User{}.GetUser)
	}
}
