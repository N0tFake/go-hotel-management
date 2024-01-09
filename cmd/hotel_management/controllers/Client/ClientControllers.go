package client_controller

import (
	"log"
	"net/http"

	model_client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

func GetAllClients(c *gin.Context) {
	var clients []model_client.Client
	service.DB.Find(&clients)
	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

func GetClientByCPF(c *gin.Context) {
	var input model_client.InputFindClientByCPF

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client model_client.Client
	if err := service.DB.Where("cpf = ?", input.CPF).Find(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client": client})
}

func CreateClient(c *gin.Context) {
	var input model_client.InputClient

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := model_client.Client{
		Name: input.Name,
		CPF:  input.CPF,
	}

	log.Println(client)

	resutl := service.DB.Create(&client)
	if resutl.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resutl.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}
