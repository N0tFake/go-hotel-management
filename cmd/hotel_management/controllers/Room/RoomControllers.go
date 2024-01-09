package room_controller

import (
	"net/http"

	model_client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	model_room "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Room"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

func LinkClientToRoom(c *gin.Context) {
	var input model_room.InputLinkClient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var room model_room.Room
	if err := service.DB.Where("code = ?", input.CodeRoom).First(&room).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	if room.Reserved {
		c.JSON(http.StatusConflict, gin.H{"error": "Room already reserved"})
		return
	}

	var client model_client.Client
	if err := service.DB.Where("cpf = ?", input.CpfClient).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	room.Client = &client
	room.Reservation_start = input.ReservationStart
	room.Reservation_end = input.ReservationEnd
	room.Reserved = true

	if err := service.DB.Model(&room).Updates(room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	service.DB.Save(&room)

	c.JSON(http.StatusOK, gin.H{"data": room})

}

func GetAllRooms(c *gin.Context) {
	var rooms []model_room.Room
	service.DB.Find(&rooms)
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func GetRoomByID(c *gin.Context) {
	var room model_room.Room

	if err := service.DB.Where("id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.DB.Preload("Clients").First(&room)

	c.JSON(http.StatusOK, gin.H{"room": room})
}

func CreateRoom(c *gin.Context) {
	var input model_room.InputRoom

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := model_room.Room{
		Code:       input.Code,
		Tipo:       input.Tipo,
		Number_bed: input.Number_bed,
	}

	if err := service.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}
