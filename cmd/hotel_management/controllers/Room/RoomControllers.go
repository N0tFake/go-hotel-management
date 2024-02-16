package room_controller

import (
	"net/http"

	model_client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	model_room "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Room"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

// Link a client to a room
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

// Get All rooms
func GetAllRooms(c *gin.Context) {
	var rooms []model_room.Room
	service.DB.Find(&rooms)

	var output []map[string]interface{}
	for _, room := range rooms {
		output = append(output, map[string]interface{}{
			"code":              room.Code,
			"type-room":         room.Tipo,
			"number-bed":        room.Number_bed,
			"client":            room.ClientID,
			"account":           room.AccountID,
			"reservation-start": room.Reservation_start,
			"reservation-end":   room.Reservation_end,
			"reserved":          room.Reserved,
		})
	}

	c.JSON(http.StatusOK, gin.H{"rooms": output})
}

// Get Room by ID
func GetRoomByID(c *gin.Context) {
	var room model_room.Room

	if err := service.DB.Where("id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// service.DB.Preload("Clients").First(&room)

	// var client model_client.Client
	// if err := service.DB.Where("id = ?", room.ClientID).First(&client).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Client not found"})
	// 	return
	// }
	// client_data := map[string]interface{}{
	// 	"ID":   client.ID,
	// 	"name": client.Name,
	// 	"cpf":  client.CPF,
	// }

	// var account model_account.Account
	// if err := service.DB.Where("id = ?", room.AccountID).First(&account).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Account not found"})
	// 	return
	// }
	// service.DB.Preload("Orders").Find(&account)
	// account_data := map[string]interface{}{
	// 	"ID":      account.ID,
	// 	"Orders":  account.Orders,
	// 	"PaidOut": account.PaidOut,
	// 	"Total":   account.Total,
	// }

	output := map[string]interface{}{
		"code":              room.Code,
		"type-room":         room.Tipo,
		"number-bed":        room.Number_bed,
		"client":            room.ClientID,
		"account":           room.AccountID,
		"reservation-start": room.Reservation_start,
		"reservation-end":   room.Reservation_end,
		"reserved":          room.Reserved,
	}

	c.JSON(http.StatusOK, gin.H{"room": output})
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

// Check Out
// Close the account
// ! Terminar o Checkout
// ! Remover o Client, as Data de Reserva e atribuir como False o reserved
func CloseAccountCheckOut(c *gin.Context) {
	var room model_room.Room

	if err := service.DB.Preload("Account").First(&room, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	room.Account.PaidOut = true
	if err := service.DB.Model(&room).Updates(room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	service.DB.Save(room.Account)

	room.AccountID = nil
	room.Account = nil
	if err := service.DB.Model(&room).Updates(room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	service.DB.Save(&room)

	// room.Account = nil
	c.JSON(http.StatusOK, gin.H{"data": room})
}
