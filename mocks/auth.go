package mocks

import (
	"github.com/andresidrim/cesupa-hospital/models"
)

type MockAuthService struct {
	MockLogin    func(cpf, password string) (string, error)
	MockRegister func(user *models.User) error
}

func (m *MockAuthService) Login(cpf, password string) (string, error) {
	if m.MockLogin != nil {
		return m.MockLogin(cpf, password)
	}
	return "", nil
}

func (m *MockAuthService) Register(user *models.User) error {
	if m.MockRegister != nil {
		return m.MockRegister(user)
	}
	return nil
}
