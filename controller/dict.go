package controller

import (
	"fmt"
	"project1/model"
	"project1/utils"

	"github.com/gin-gonic/gin"
)

type AddDictRequest struct {
	Keyy   string `json:"keyy" binding:"required"`   // 必填字段
	Valuee string `json:"valuee" binding:"required"` // 必填字段
	Type   string `json:"type" binding:"required"`   // 必填字段
}
type UpdateDictRequest struct {
	ID     uint   `json:"id" binding:"required"` // 必填字段
	Keyy   string `json:"keyy" `
	Valuee string `json:"valuee"`
	Type   string `json:"type" `
}

type dictType struct {
	Type string `json:"type" binding:"required"` // 必填字段
}

type Dict struct{}

// GetAll 获取字典数据
func (Dict) GetAll(c *gin.Context) {
	dictList := []model.Dict{}
	model.DB.Find(&dictList)
	user, _ := c.Get("user")
	fmt.Println(user)
	utils.ResponseSuccess(c, dictList)

}

// 根据type 获取字典列表
func (Dict) GetByType(c *gin.Context) {
	var query dictType

	err := c.ShouldBind(&query)
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	queryType := query.Type

	dictList := []model.Dict{}
	model.DB.Where("type = ?", queryType).Find(&dictList)
	utils.ResponseSuccess(c, dictList)
}

// 更新字典数据
func (Dict) UpdateDict(c *gin.Context) {
	var data UpdateDictRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	dict := model.Dict{
		ID:     data.ID,
		Type:   data.Type,
		Valuee: data.Valuee,
		Keyy:   data.Keyy,
	}
	model.DB.Save(&dict)
	utils.ResponseSuccess(c, dict)

}

func (Dict) AddDict(c *gin.Context) {
	var data AddDictRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	dict := model.Dict{
		Keyy:   data.Keyy,
		Valuee: data.Valuee,
		Type:   data.Type,
	}

	err := model.DB.Create(&dict)
	if err != nil {
		utils.ResponseError(c, 0, "添加失败")
	} else {
		utils.ResponseSuccess(c, err)

	}

}

// 删除字典数据
func (Dict) DeleteDict(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.ResponseError(c, 0, "id不能为空")
		return
	}
	dict := model.Dict{}
	model.DB.Where("id = ?", id).Delete(&dict)
	utils.ResponseSuccess(c, true)
}
