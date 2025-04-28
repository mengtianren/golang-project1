package controller

import (
	"project1/model"
	"project1/utils"

	"github.com/gin-gonic/gin"
)

type AddUserRequest struct {
	Name     string `json:"name" binding:"required"`     // 必填字段
	Password string `json:"password" binding:"required"` // 必填字段
	Phone    string `json:"phone" binding:"required"`
	Roles    []uint `json:"roles" binding:"required"` // 角色列表
}

type Admin struct{}

// 管理员添加用户
func (Admin) AddUser(c *gin.Context) {
	var data AddUserRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	user := model.User{
		Name:     data.Name,
		Password: data.Password,
		Phone:    data.Phone,
	}

	roles := []model.Role{}

	if err := model.DB.Where("id in ?", data.Roles).Find(&roles).Error; err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	user.Roles = roles

	if err := model.DB.Create(&user).Error; err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}

	utils.ResponseSuccess(c, user)
}

// 删除用户
func (Admin) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.ResponseError(c, 0, "id不能为空")
		return
	}
	var user model.User
	model.DB.Where("id = ?", id).Delete(&user)

	utils.ResponseSuccess(c, user)
}
