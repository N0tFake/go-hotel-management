package model_sale

type Sale struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Value   float64
	RoomRef int
}
