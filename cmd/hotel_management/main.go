package main

import (
	"github.com/N0tFake/go-hotel-management/cmd/hotel_management/routes"
	initializers "github.com/N0tFake/go-hotel-management/configs/initializations"
	"github.com/N0tFake/go-hotel-management/configs/service"
)

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic("Error loading environment variables")
	}

	service.ConnectDatabase(&config)

	r := routes.SetupRouter()

	err = r.Run()
	if err != nil {
		return
	}
}
