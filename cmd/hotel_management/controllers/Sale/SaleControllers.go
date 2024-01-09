package sale_controller

import (
	"net/http"
	"time"

	model_account "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Account"
	model_sale "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Sale"
	"github.com/N0tFake/go-hotel-management/configs/service"
	"github.com/gin-gonic/gin"
)

func AddOrderOnAccount(c *gin.Context) {
	var input model_sale.AddOrder

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account model_account.Account
	if err := service.DB.Where("id = ?", input.AccountRef).First(&account).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Account not found"})
		return
	}

	sale := model_sale.Sale{
		Product:    input.Product,
		Date:       time.Now(),
		Value:      input.Value,
		Quant:      input.Quant,
		Total:      float64(input.Quant) * float64(input.Value),
		AccountRef: input.AccountRef,
	}

	service.DB.Preload("Orders").Find(&account)

	service.DB.Model(&account).Association("Orders").Append(&sale)

	var total float64
	for _, order := range account.Orders {
		total += order.Total
	}

	account.Total = total
	service.DB.Save(&account)

	c.JSON(http.StatusCreated, gin.H{"data": account})
}
