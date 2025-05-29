package mocks

import (
	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/models"
)

type MockUserService struct {
	MockGet    func(id uint64) (*models.User, error)
	MockGetAll func(roles []enums.Role) ([]models.User, error)
}

func (m *MockUserService) Get(id uint64) (*models.User, error) {
	return m.MockGet(id)
}

func (m *MockUserService) GetAll(roles []enums.Role) ([]models.User, error) {
	return m.MockGetAll(roles)
}
