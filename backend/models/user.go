package models

import (
	"github.com/andresidrim/cesupa-hospital/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string        `gorm:"not null" json:"name"`
	CPF          string        `gorm:"unique;not null" json:"cpf"`
	Password     string        `gorm:"not null" json:"-"`
	Role         enums.Role    `gorm:"not null" json:"role"`
	Appointments []Appointment `gorm:"foreignKey=UserID;constraint:OnDelete:CASCADE" json:"appointments"`
}
