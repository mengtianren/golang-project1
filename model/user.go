package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// 重新定义 ID 字段并添加 JSON 标签
	ID        uint           `gorm:"primaryKey" json:"id"`    // 自定义 JSON 字段名为 "id"
	CreatedAt time.Time      `json:"created_at"`              // 自定义为 "created_at"
	UpdatedAt time.Time      `json:"updated_at"`              // 自定义为 "updated_at"
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // 自定义为 "deleted_at"
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Roles     []Role         `json:"roles" gorm:"many2many:user_roles;"` // 多对多
}

func (User) TableName() string {
	return "t_user"
}
