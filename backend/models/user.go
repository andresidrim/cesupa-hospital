package models

import (
	"github.com/andresidrim/cesupa-hospital/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string     `gorm:"not null"`
	CPF  string     `gorm:"unique;not null"`
	Role enums.Role `gorm:"not null"`
}

type Doctor struct {
	gorm.Model
	UserID       uint `gorm:"unique;not null"`
	User         User
	Appointments []Appointment `gorm:"foreignKey=DoctorID;constraint:OnDelete:CASCADE"`
}
