package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitAdminRoute(r *gin.Engine) {
	router := r.Group("/admin")
	{
		router.GET("/get", func(c *gin.Context) {
			c.String(200, "get admin")
		})
		router.POST("/add", controller.Admin{}.AddUser)
		router.DELETE("/del", controller.Admin{}.DeleteUser)
		router.GET("/list", controller.Admin{}.GetUserList)
	}
}
