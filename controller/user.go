package controller

import (
	"fmt"
	"net/http"
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
type UserInfo struct {
	ID    uint         `json:"id"`
	Name  string       `json:"name"`
	Phone string       `json:"phone"`
	Roles []model.Role `json:"roles" gorm:"many2many:user_roles;"` // 多对多
}

func (UserInfo) TableName() string {
	return "t_user"
}

type Register struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Roles    []uint `json:"roles"` // 角色列表
}

/*
* @api {post} /api/user/login 登录
* @apiName Login
* @apiGroup User
* @apiParam {String} phone 手机号
* @apiParam {String} password 密码
* @apiSuccess {String} token 登录凭证
* @apiSuccessExample {json} Success-Response:
*     HTTP/1.1 200 OK
*     {
*       "code": 200,
}
*/
func (User) Login(c *gin.Context) {
	var data Login
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, "参数错误")
		return
	}
	var user model.User
	err := model.DB.Preload("Roles").Where("phone = ? AND password = ?", data.Phone, data.Password).First(&user).Error

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
	token, err1 := utils.GenerateJWT(int(user.ID), user.Name, user.Phone, roles)
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

/*
* @api {post} /api/user/register 注册
* @apiName Register
* @apiGroup User
* @apiParam {String} name 用户名
* @apiParam {String} password 密码
* @apiParam {String} phone 手机号
* @apiParam {String} roles 角色
* @apiSuccessExample {json} Success-Response:
*     HTTP/1.1 200 OK
*     {
*       "code": 200,
}
*/
func (User) Register(c *gin.Context) {
	var data Register
	err := c.ShouldBind(&data)
	if err != nil {
		utils.ResponseError(c, 0, "参数错误")
		return
	}
	var roles []model.Role

	err2 := model.DB.Where("id in ?", data.Roles).Find(&roles).Error
	fmt.Println(roles, data.Roles)
	if err2 != nil {
		utils.ResponseError(c, 0, err2.Error())
		return

	}

	var user = model.User{Name: data.Name, Phone: data.Phone, Password: data.Password, Roles: roles}

	err1 := model.DB.Create(&user).Error
	if err1 != nil {
		utils.ResponseError(c, 0, err1.Error())
		return
	}
	utils.ResponseSuccess(c, "注册成功")

}

/*
* @api {get} /api/user/get 获取用户信息
* @apiName GetUserInfo
* @apiGroup User
 */
func (User) GetUser(c *gin.Context) {
	// user := model.User{}
	user, exists := c.Get("user")
	if exists {
		fmt.Println("获取到的用户信息:", user)
	} else {
		fmt.Println("未获取到用户信息")
	}
	fmt.Println(user)
	// 将 int 类型的 user.(utils.Claims).ID 转换为 uint 类型
	responseUser := model.User{ID: uint(user.(*utils.Claims).ID)}

	model.DB.Select("*").Preload("Roles").Find(&responseUser)
	fmt.Println(responseUser)

	utils.ResponseSuccess(c, gin.H{
		"id":    responseUser.ID,
		"name":  responseUser.Name,
		"phone": responseUser.Phone,
		"roles": responseUser.Roles, // 这里需要根据你的实际情况调整字段名和类型，这里假设 roles 是一个切片或数组，你可以根据需要调整类型
	})

}

/*
* 修改用户信息
 */
func (User) PutUser(c *gin.Context) {
	var updateData struct {
		ID       uint   `json:"id" required:"true" min:"1"` // 假设用户ID是必需的
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var existingUser model.User
	err := model.DB.Where("id = ?", updateData.ID).Find(&existingUser).Error
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "查询失败")
		return
	}
	if updateData.Name != "" {
		existingUser.Name = updateData.Name
	}
	if updateData.Phone != "" {
		existingUser.Phone = updateData.Phone
	}
	if updateData.Password != "" {
		existingUser.Password = updateData.Password
	}

	// 保存更新
	result := model.DB.Save(&existingUser)
	if result.Error != nil {
		utils.ResponseError(c, http.StatusBadRequest, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		utils.ResponseError(c, http.StatusBadRequest, "未更新任何记录")
		return
	}

	utils.ResponseSuccess(c, existingUser)
}

func (User) AddUser(c *gin.Context) {
	// user := model.User{}

}
