package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回值正确统一返回
func ResponseSuccess(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "操作成功",
		"data": data,
	})
}

// 返回值错误统一返回
func ResponseError(c *gin.Context, code int, msg string) {
	if code == 0 {
		code = 400 // 默认错误码
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
	c.Abort() // 阻断后续处理
}
