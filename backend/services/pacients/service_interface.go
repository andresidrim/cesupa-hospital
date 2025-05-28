package pacients

import "github.com/andresidrim/cesupa-hospital/models"

type PacientService interface {
	Create(pacient *models.Pacient) error
	Get(id uint64) (*models.Pacient, error)
	GetAll(name string, ageStr string) ([]models.Pacient, error)
	Update(id uint64, pacient *models.Pacient) error
	Delete(id uint64) error
}
