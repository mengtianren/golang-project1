package main

import (
	"project1/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":8088")

}
