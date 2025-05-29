package models

import (
	"time"

	"github.com/andresidrim/cesupa-hospital/enums"
	"gorm.io/gorm"
)

type Pacient struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string    `gorm:"not null" json:"name"`
	BirthDate   time.Time `gorm:"type:date;not null" json:"birthDate"`
	CPF         string    `gorm:"unique;not null" json:"cpf"`
	Sex         enums.Sex `gorm:"not null" json:"sex"`
	PhoneNumber string    `gorm:"not null" json:"phoneNumber"`
	Address     string    `gorm:"not null" json:"address"`

	Email     *string          `json:"email"`
	BloodType *enums.BloodType `json:"bloodType"`
	Allergies *string          `json:"allergies"`

	Appointments []Appointment `gorm:"foreignKey=PacientID;constraint:OnDelete:CASCADE" json:"appointments" swaggerignore:"true"`
}
