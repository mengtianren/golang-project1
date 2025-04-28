package controller

import (
	"fmt"
	"project1/model"
	"project1/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct{}

type Login struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (User) Login(c *gin.Context) {
	var data Login
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, "参数错误")
		return
	}
	var user model.User
	err := model.DB.Where("phone = ? AND password = ?", data.Phone, data.Password).Preload("Roles").First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseError(c, 0, "用户不存在")
			return
		} else {
			utils.ResponseError(c, 0, "查询失败")
			return
		}
	}

	var roles = []int{}
	for _, role := range user.Roles {
		roles = append(roles, int(role.ID))

	}
	if user.ID == 0 {
		utils.ResponseError(c, 0, "系统错误")
		return
	}

	fmt.Println(roles)
	token, err1 := utils.GenerateJWT(user.ID, user.Name, user.Phone, roles)
	if err1 != nil {
		utils.ResponseError(c, 0, "登录失败")
		return

	}

	fmt.Println(user)
	utils.ResponseSuccess(c, gin.H{
		"userInfo": user,
		"token":    token,
	})

}

func (User) GetUser(c *gin.Context) {
	userList := []model.User{}

	model.DB.Select("*").Preload("Roles").Find(&userList)
	fmt.Println(userList)

	utils.ResponseSuccess(c, userList)

}

func (User) AddUser(c *gin.Context) {
	// user := model.User{}

}
