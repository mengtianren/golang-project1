package router

import (
	"project1/controller"

	"github.com/gin-gonic/gin"
)

func InitDictRoute(r *gin.Engine) {
	router := r.Group("/dict")
	var dict controller.Dict
	{
		router.GET("/get/type", dict.GetList)
	}
}
