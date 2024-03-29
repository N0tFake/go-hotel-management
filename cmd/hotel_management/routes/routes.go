package routes

import (
	"fmt"

	account_controller "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Account"
	client_controller "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Client"
	room_controllers "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Room"
	sale_controller "github.com/N0tFake/go-hotel-management/cmd/hotel_management/controllers/Sale"
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

	r.POST("/link/client/room/", room_controllers.LinkClientToRoom)
	r.GET("/room/:id/checkout/", room_controllers.CloseAccountCheckOut)

	r.GET("/clients", client_controller.GetAllClients)
	r.POST("/create/client", client_controller.CreateClient)
	r.POST("/client/cpf", client_controller.GetClientByCPF)

	r.POST("/account", account_controller.GetAccountByCPF)
	r.POST("/checkin", account_controller.CreateAccountCheckIn)

	r.POST("/add/sale", sale_controller.AddOrderOnAccount)

	return r

}
