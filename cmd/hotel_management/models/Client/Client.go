package client

type Client struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	CPF  string
}
