package main

import (
	"fmt"
	"project1/config"
	"project1/model"
	"project1/router"
)

func main() {
	err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}
	model.InitDB()

	r := router.InitRouter()

	r.Run(fmt.Sprintf(":%d", config.AppConfig.App.Port))

}
