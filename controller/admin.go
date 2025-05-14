package controller

import (
	"fmt"
	"project1/model"
	"project1/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type addUserRequest struct {
	Name     string       `json:"name" binding:"required"`     // 必填字段
	Password string       `json:"password" binding:"required"` // 必填字段
	Phone    string       `json:"phone" binding:"required"`
	Roles    []model.Role `json:"roles" binding:"required"` // 角色列表
}

type DeleteUserRequest struct {
	ID uint `form:"id" json:"id" binding:"required"`
}
type updateUserRequest struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"` // 必填字段
	Phone string `json:"phone" binding:"required"`
	Roles []uint `json:"roles"` // 角色列表
}
type RoleDTO struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type UserDTO struct {
	ID    uint      `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
	Roles []RoleDTO `json:"roles"`
}
type Admin struct{}

// 管理员添加用户
func (Admin) AddUser(c *gin.Context) {
	var data addUserRequest
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
	var query DeleteUserRequest
	err := c.ShouldBind(&query)
	fmt.Println(query)
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	var user = model.User{
		ID: query.ID,
	}
	err1 := model.DB.Delete(&user).Error
	if err1 != nil {
		utils.ResponseError(c, 0, err1.Error())
		return
	}
	fmt.Println(user)

	utils.ResponseSuccess(c, true)
}

// 获取用户列表
func (Admin) GetUserList(c *gin.Context) {

	var users []model.User

	err := model.DB.Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Select("name", "type", "id") // 只查这两列
	}).Find(&users).Error
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	// 转换为精简结构体
	userList := []UserDTO{}

	for _, user := range users {
		// 如果数据不存在返回 []
		roles := []RoleDTO{}
		// 如果数据不存在 返回 null
		// var roles []RoleDTO
		fmt.Println("roles", roles)
		for _, role := range user.Roles {
			roles = append(roles, RoleDTO{
				Name: role.Name,
				Type: role.Type,
			})
		}
		userList = append(userList, UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Roles: roles,
		})

	}

	utils.ResponseSuccess(c, userList)
}

func (Admin) UpdateUser(c *gin.Context) {
	var query updateUserRequest
	err := c.ShouldBind(&query)
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	var roles []model.Role
	err1 := model.DB.Where("id in ?", query.Roles).Find(&roles).Error
	if err1 != nil {
		utils.ResponseError(c, 0, err1.Error())
		return

	}

	var user = model.User{
		ID:    query.ID,
		Name:  query.Name,
		Phone: query.Phone,
		Roles: roles,
	}

	err2 := model.DB.Save(&user).Error
	if err2 != nil {
		utils.ResponseError(c, 0, err2.Error())
		return

	}
	utils.ResponseSuccess(c, true)
}
