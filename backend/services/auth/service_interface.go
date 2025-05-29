package auth

import "github.com/andresidrim/cesupa-hospital/models"

type AuthService interface {
	Login(cpf, password string) (string, error)
	Register(user *models.User) error
}
