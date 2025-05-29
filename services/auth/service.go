package auth

import (
	"fmt"

	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/andresidrim/cesupa-hospital/utils"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Login(cpf, password string) (string, error) {
	var user models.User
	if err := s.db.Where("cpf = ?", cpf).First(&user).Error; err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", fmt.Errorf("incorrect password: %v", err)
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Register(user *models.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.db.Create(user).Error
}
