package controller

import (
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
	ID     uint   `json:"id" binding:"required" min:"1"` // 必填字段
	Keyy   string `json:"keyy" `
	Valuee string `json:"valuee"`
	Type   string `json:"type" `
}

type SearchDictRequest struct {
	Type string `json:"type" form:"type"` // 必填字段
}

type Dict struct{}

// 获取字典分页
func (Dict) GetPage(c *gin.Context) {
	tx := model.DB.Model(&model.Dict{})
	result, err := utils.NewPagedResult(c, tx, []model.Dict{})
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	utils.ResponseSuccess(c, result)

}

// 根据type 获取字典列表
func (Dict) GetList(c *gin.Context) {
	var query SearchDictRequest
	err := c.ShouldBind(&query)
	if err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	tx := model.DB.Model(&model.Dict{})
	if query.Type != "" {
		tx = tx.Where("type =?", query.Type)
	}
	dictList := []model.Dict{}
	tx.Find(&dictList)
	utils.ResponseSuccess(c, dictList)
}

// 更新字典数据
func (Dict) EditItem(c *gin.Context) {
	var data UpdateDictRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	dict := model.Dict{}
	err := model.DB.Where("id =?", data.ID).First(&dict).Error
	if err != nil {
		utils.ResponseError(c, 0, "字典不存在")
		return
	}
	if data.Keyy != "" {
		dict.Keyy = data.Keyy
	}
	if data.Valuee != "" {
		dict.Valuee = data.Valuee
	}
	if data.Type != "" {
		dict.Type = data.Type
	}

	model.DB.Save(&dict)
	utils.ResponseSuccess(c, dict)

}

func (Dict) AddItem(c *gin.Context) {
	var data AddDictRequest
	if err := c.ShouldBind(&data); err != nil {
		utils.ResponseError(c, 0, err.Error())
		return
	}
	dict := model.Dict{
		Keyy:   data.Keyy,
		Valuee: data.Valuee,
		Type:   data.Type,
	}

	err1 := model.DB.Where("keyy =? AND valuee=? AND type=?", data.Keyy, data.Valuee, data.Type).First(&dict).Error
	if err1 == nil {
		utils.ResponseError(c, 0, "字典已存在")
		return
	}

	err := model.DB.Create(&dict).Error
	if err != nil {
		utils.ResponseError(c, 0, "添加失败")
	} else {
		utils.ResponseSuccess(c, dict)

	}

}

// 删除字典数据
func (Dict) DelItem(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.ResponseError(c, 0, "id不能为空")
		return
	}
	dict := model.Dict{}
	model.DB.Where("id = ?", id).Delete(&dict)
	utils.ResponseSuccess(c, true)
}
