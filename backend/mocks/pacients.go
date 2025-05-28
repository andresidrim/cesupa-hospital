package mocks

import "github.com/andresidrim/cesupa-hospital/models"

type MockPacientService struct {
	MockCreate func(pacient *models.Pacient) error
	MockGet    func(id uint64) (*models.Pacient, error)
	MockGetAll func(name string, ageStr string) ([]models.Pacient, error)
	MockUpdate func(id uint64, pacient *models.Pacient) error
}

func (m *MockPacientService) GetAll(name, ageStr string) ([]models.Pacient, error) {
	return m.MockGetAll(name, ageStr)
}

func (m *MockPacientService) Get(id uint64) (*models.Pacient, error) {
	if m.MockGet != nil {
		return m.MockGet(id)
	}
	return nil, nil
}

func (m *MockPacientService) Create(pacient *models.Pacient) error {
	if m.MockCreate != nil {
		return m.MockCreate(pacient)
	}
	return nil
}

func (m *MockPacientService) Update(id uint64, pacient *models.Pacient) error {
	if m.MockUpdate != nil {
		return m.MockUpdate(id, pacient)
	}
	return nil
}
