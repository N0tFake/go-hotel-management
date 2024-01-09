package account_controller

import (
	"net/http"

	model_account "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Account"
	model_client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	model_room "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Room"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

func GetAccountByCPF(c *gin.Context) {
	var input model_account.FindAccountByCPF

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account model_account.Account
	var client model_client.Client
	if err := service.DB.Where("cpf = ?", input.CPF).Find(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not Found"})
		return
	}

	if err := service.DB.Where("payer_id = ?", client.ID).Find(&account).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not Found"})
		return
	}

	service.DB.Preload("Payer").First(&account)
	service.DB.Preload("Orders").Find(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})

}

func CreateAccount(c *gin.Context) {
	var input model_account.InputCreateAccount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client model_client.Client
	result_client := service.DB.First(&client, input.Payer_id)
	if result_client.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	var room model_room.Room
	result_room := service.DB.First(&room, "id = ?", input.Room_id)
	if result_room.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	account := model_account.Account{
		Payer:   &client,
		Room_id: input.Room_id,
	}

	service.DB.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})
}
