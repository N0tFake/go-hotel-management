package main

import (
	initializers "github.com/N0tFake/go-hotel-management/configs/initializations"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic("Error loading environment variables")
	}

	service.ConnectDatabase(&config)

	err = r.Run()
	if err != nil {
		return
	}
}
