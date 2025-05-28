package models

import (
	"time"

	"github.com/andresidrim/cesupa-hospital/enums"
	"gorm.io/gorm"
)

type Pacient struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	BirthDate   time.Time `gorm:"type:date;not null"`
	CPF         string    `gorm:"unique;not null"`
	Sex         enums.Sex `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	Address     string    `gorm:"not null"`

	Email     *string
	BloodType *enums.BloodType
	Allergies *string

	Appointments []Appointment `gorm:"foreignKey=PacientID;constraint:OnDelete:CASCADE"`
}
