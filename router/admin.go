package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitAdminRoute(r *gin.Engine) {
	router := r.Group("/admin")
	adminController := controller.Admin{}
	{
		userR := router.Group("/user")
		{
			userR.POST("/add", adminController.AddUser)
			userR.DELETE("/del", adminController.DeleteUser)
			userR.POST("/list", adminController.GetUserList)
			userR.PUT("/edit", adminController.UpdateUser)
		}

	}
}
