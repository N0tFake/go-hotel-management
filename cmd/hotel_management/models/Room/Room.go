package model_room

import (
	"time"

	model_account "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Account"
	client "github.com/N0tFake/go-hotel-management/cmd/hotel_management/models/Client"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Code              string `gorm:"unique"`
	Tipo              string
	Number_bed        int
	ClientID          *int
	Client            *client.Client         `gorm:"foreignKey:ClientID"`
	Account           *model_account.Account `gorm:"foreignKey:AccountID"`
	AccountID         *int
	Reservation_start time.Time
	Reservation_end   time.Time
	Reserved          bool `gorm:"default:false"`
}

type InputRoom struct {
	Code       string `json:"code" binding:"required"`
	Tipo       string `json:"tipo" binding:"required"`
	Number_bed int    `json:"number_bed" binding:"required"`
}

type InputLinkClient struct {
	CodeRoom         string    `json:"code_room" binding:"required"`
	CpfClient        string    `json:"cpf_client" binding:"required"`
	ReservationStart time.Time `json:"reservation_start" binding:"required"`
	ReservationEnd   time.Time `json:"reservation_end" binding:"required"`
}
