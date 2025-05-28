package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PacientID uint `gorm:"not null"`
	DoctorID  uint `gorm:"not null;constraint:OnDelete:CASCADE"`
	Doctor    Doctor
	Date      time.Time `gorm:"not null"`
}
