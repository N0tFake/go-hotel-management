package model_client

type Client struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	CPF  string
}

type InputClient struct {
	Name string `json:"name" binding:"required"`
	CPF  string `json:"cpf" binding:"required"`
}

type InputFindClientByCPF struct {
	CPF string `json:"cpf" binding:"required"`
}
