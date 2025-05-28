package pacients

import (
	"fmt"
	"strconv"
	"time"

	"github.com/andresidrim/cesupa-hospital/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(pacient *models.Pacient) error {
	return s.db.Create(pacient).Error
}

func (s *Service) Get(id uint64) (*models.Pacient, error) {
	var pacient models.Pacient
	if err := s.db.First(&pacient, id).Error; err != nil {
		return nil, err
	}

	return &pacient, nil
}

func (s *Service) GetAll(name string, ageStr string) ([]models.Pacient, error) {
	var pacients []models.Pacient
	query := s.db.Model(&models.Pacient{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid age: %v", err)
		}

		from, to := calculateAgeRange(age)
		query = query.Where("birth_date BETWEEN ? AND ?", from, to)
	}

	if err := query.Find(&pacients).Error; err != nil {
		return nil, err
	}

	return pacients, nil
}

func calculateAgeRange(age int) (time.Time, time.Time) {
	now := time.Now()
	from := now.AddDate(-age-1, 0, 1)
	to := now.AddDate(-age, 0, 0)
	return from, to
}
