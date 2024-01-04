package routes

import (
	"fmt"

	client_controller "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Client"
	room_controllers "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Room"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		fmt.Println("Teste /")
	})

	r.GET("/rooms", room_controllers.GetAllRooms)
	r.GET("/room/:id", room_controllers.GetRoomByID)
	r.POST("/create/room", room_controllers.CreateRoom)

	r.GET("/clients", client_controller.GetAllClients)
	r.POST("/create/client", client_controller.CreateClient)
	r.POST("/client/cpf", client_controller.GetClientByCPF)

	return r

}
