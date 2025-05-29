package auth

import "github.com/andresidrim/cesupa-hospital/enums"

type RegisterDTO struct {
	Name     string     `json:"name" binding:"required"`
	CPF      string     `json:"cpf" binding:"required"`
	Password string     `json:"password" binding:"required,min=6"`
	Role     enums.Role `json:"role" binding:"required"`
}

type LoginDTO struct {
	CPF      string `json:"cpf" binding:"required"`
	Password string `json:"password" binding:"required"`
}
