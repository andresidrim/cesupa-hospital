package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PacientID uint      `gorm:"not null" json:"pacientId"`
	Pacient   Pacient   `json:"pacient"`
	UserID    uint      `gorm:"not null" json:"userId"`
	User      User      `json:"user"`
	Date      time.Time `gorm:"not null" json:"date"`
}
