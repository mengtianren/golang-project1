package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(r *gin.Engine) {
	router := r.Group("/user")
	userController := controller.User{}
	{
		router.GET("/info", userController.GetUser)
		router.PUT("/info", userController.PutUser)
	}
}
