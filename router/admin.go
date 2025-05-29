package router

import (
	"project1/controller"
	"project1/middleware"

	"github.com/gin-gonic/gin"
)

func InitAdminRoute(r *gin.Engine) {
	router := r.Group("/admin")
	router.Use(middleware.AdminMiddleware())
	var userController controller.User
	{
		userR := router.Group("/user")
		{
			userR.POST("/page", userController.GetPage)
			userR.POST("/add", userController.AddItem)
			userR.PUT("/edit", userController.EditItem)
			userR.DELETE("/del", userController.DelItem)

		}
		dictR := router.Group("/dict")
		var dictController controller.Dict
		{
			dictR.POST("/page", dictController.GetPage)
			dictR.POST("/add", dictController.AddItem)
			dictR.PUT("/edit", dictController.EditItem)
			dictR.DELETE("/del", dictController.DelItem)
		}

	}
}
