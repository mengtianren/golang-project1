package model

import (
	"fmt"
	"project1/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	fmt.Println("数据库连接中", config.AppConfig)
	// dsn := "root:mengtianren@mysql@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.DBName,
		config.AppConfig.Database.Charset,
		config.AppConfig.Database.ParseTime,
		config.AppConfig.Database.Loc)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	DB.AutoMigrate(&User{}, &Role{}, &Menu{}, &Dict{})

	fmt.Println("数据库连接成功")
}
