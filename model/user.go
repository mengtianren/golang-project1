package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles;"` // 多对多
}

func (User) TableName() string {
	return "t_user"
}
