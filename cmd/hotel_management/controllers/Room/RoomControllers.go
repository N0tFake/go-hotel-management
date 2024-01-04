package room_controller

import (
	"net/http"

	model_room "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Room"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

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

	service.DB.Create(&room)

	c.JSON(http.StatusOK, gin.H{"data": room})
}
