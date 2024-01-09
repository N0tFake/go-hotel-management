package model_account

import (
	model_client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	model_sale "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Sale"
)

type Account struct {
	ID      uint `gorm:"primaryKey"`
	PayerID *int
	Payer   *model_client.Client `gorm:"foreignKey:PayerID"`
	Room_id int
	Orders  []model_sale.Sale `gorm:"foreignKey:AccountRef"`
	PaidOut bool              `gorm:"default:false"`
	Total   float64
}

type InputCreateAccount struct {
	Payer_id int `json:"payer_id" binding:"required"`
	Room_id  int `json:"room_id" binding:"required"`
}

type FindAccountByCPF struct {
	CPF string `json:"cpf" binding:"required"`
}
