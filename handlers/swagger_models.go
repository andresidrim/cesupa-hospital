package handlers

import "github.com/andresidrim/cesupa-hospital/enums"

// ErrorResponse é usado para todas as falhas
type ErrorResponse struct {
	Message string `json:"message"`
}

// RegisterResponse é o payload de sucesso de /register
type RegisterResponse struct {
	ID   uint       `json:"id"`
	Name string     `json:"name"`
	CPF  string     `json:"cpf"`
	Role enums.Role `json:"role"`
}

// TokenResponse é o payload de sucesso de /login
type TokenResponse struct {
	Token string `json:"token"`
}
