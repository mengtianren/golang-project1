package model

import (
	"time"
)

// 字典相关
type Dict struct {
	// 重新定义 ID 字段并添加 JSON 标签
	ID        uint      `gorm:"primaryKey" json:"id"` // 自定义 JSON 字段名为 "id"
	CreatedAt time.Time `json:"created_at"`           // 自定义为 "created_at"
	UpdatedAt time.Time `json:"updated_at"`           // 自定义为 "updated_at"
	Keyy      string    `json:"keyy"`
	Valuee    string    `json:"valuee"`
	Type      string    `json:"type"`
}

func (Dict) TableName() string {
	return "t_dict"
}
