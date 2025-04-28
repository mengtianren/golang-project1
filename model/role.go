package model

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (Role) TableName() string {
	return "t_role"
}
