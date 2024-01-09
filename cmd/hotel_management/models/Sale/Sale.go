package model_sale

import "time"

type Sale struct {
	ID         uint `gorm:"primaryKey"`
	Product    string
	Date       time.Time
	Value      float64
	Total      float64
	Quant      int
	AccountRef int
}

type AddOrder struct {
	Product    string  `json:"product" binding:"required"`
	Value      float64 `json:"value" binding:"required"`
	Quant      int     `json:"quant" binding:"required"`
	AccountRef int     `json:"account_ref" binding:"required"`
}
