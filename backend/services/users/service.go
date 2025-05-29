package users

import (
	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Get(id uint64) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Appointments").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) GetAll(filterRoles []enums.Role) ([]models.User, error) {
	var users []models.User
	query := s.db.Model(&models.User{}).Preload("Appointments")

	if len(filterRoles) > 0 {
		query = query.Where("role IN ?", filterRoles)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
