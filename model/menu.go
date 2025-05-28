package model

import (
	"time"
)

// 字典相关
type Menu struct {
	// 重新定义 ID 字段并添加 JSON 标签
	ID        uint      `gorm:"primaryKey" json:"id"`                // 自定义 JSON 字段名为 "id"
	Name      string    `json:"name"`                                // 自定义 JSON 字段名为 "name"
	Path      string    `json:"path"`                                // 自定义 JSON 字段名为 "path"
	Icon      string    `json:"icon"`                                // 自定义 JSON 字段名为 "icon"
	ParentID  uint      `json:"parent_id"`                           // 自定义 JSON 字段名为 "parent_id"
	Component string    `json:"component"`                           // 自定义 JSON 字段名为 "component"
	Sort      int       `json:"sort"`                                // 自定义 JSON 字段名为 "sort"
	CreatedAt time.Time `json:"created_at"`                          // 自定义为 "created_at"
	UpdatedAt time.Time `json:"updated_at"`                          // 自定义为 "updated_at"
	Children  []*Menu   `gorm:"foreignKey:ParentID" json:"children"` // 自定义为 "children"

}

func (Menu) TableName() string {
	return "t_menu"
}
