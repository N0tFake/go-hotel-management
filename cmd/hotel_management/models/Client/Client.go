package model_client

import "gorm.io/gorm"

type Client struct {
	gorm.Model
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
