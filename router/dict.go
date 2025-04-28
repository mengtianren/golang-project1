package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitDictRoute(r *gin.Engine) {
	router := r.Group("/dict")
	{
		router.GET("/get/all", controller.Dict{}.GetAll)
		router.GET("/get/type", controller.Dict{}.GetByType)
		router.POST("/add", controller.Dict{}.AddDict)
		router.POST("/edit", controller.Dict{}.UpdateDict)
		router.DELETE("/del", controller.Dict{}.DeleteDict)
	}
}
