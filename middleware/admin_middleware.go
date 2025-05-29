package middleware

import (
	"fmt"
	"project1/model"
	"project1/utils"

	"github.com/gin-gonic/gin"
)

/* *
 * @Description: 管理员权限验证
 * @return gin.HandlerFunc
 */
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			fmt.Println("获取到的用户信息:", user)
			utils.ResponseError(c, 0, "权限不足")
			return
		}
		var roles []model.Role
		err := model.DB.Where("id in ?", user.(*utils.Claims).Roles).Find(&roles).Error
		if err != nil {
			utils.ResponseError(c, 0, "权限不足")
			return
		}
		var isAdmin bool
		for _, role := range roles {
			if role.Type == "admin" {
				isAdmin = true
				break
			}

		}
		if !isAdmin {
			utils.ResponseError(c, 0, "权限不足")
			return
		}
		fmt.Println(user, isAdmin)
		c.Next() // 继续处理请求

	}
}
