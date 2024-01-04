package model_room

import (
	"time"

	client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	sale "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Sale"
)

type Room struct {
	ID                uint `gorm:"primaryKey"`
	Code              string
	Tipo              string
	Number_bed        int
	Client            client.Client `gorm:"foreignKey:ID"`
	Conta             []sale.Sale   `gorm:"foreignKey:RoomRef"`
	Reservation_start time.Time
	Reservation_end   time.Time
}

type InputRoom struct {
	Code       string `json:"code" binding:"required"`
	Tipo       string `json:"tipo" binding:"required"`
	Number_bed int    `json:"number_bed" binding:"required"`
}
