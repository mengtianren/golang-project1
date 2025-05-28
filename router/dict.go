package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitDictRoute(r *gin.Engine) {
	router := r.Group("/dict")
	var dict controller.Dict
	{
		router.POST("/get/all", dict.GetAll)
		router.GET("/get/type", dict.GetByType)
		router.POST("/add", dict.AddDict)
		router.POST("/edit", dict.UpdateDict)
		router.DELETE("/del", dict.DeleteDict)
	}
}
