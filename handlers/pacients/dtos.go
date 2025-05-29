package pacients

import (
	"time"

	"github.com/andresidrim/cesupa-hospital/enums"
)

type AddPacientDTO struct {
	Name        string           `json:"name" binding:"required"`
	BirthDate   time.Time        `json:"birthDate" binding:"required"`
	CPF         string           `json:"cpf" binding:"required"`
	Sex         enums.Sex        `json:"sex" binding:"required"`
	PhoneNumber string           `json:"phoneNumber" binding:"required"`
	Address     string           `json:"address" binding:"required"`
	Email       *string          `json:"email" binding:"omitempty,email"`
	BloodType   *enums.BloodType `json:"bloodType"`
	Allergies   *string          `json:"allergies"`
}

type UpdatePacientDTO struct {
	Name        string           `json:"name" binding:"required"`
	BirthDate   time.Time        `json:"birthDate" binding:"required"`
	CPF         string           `json:"cpf" binding:"required"`
	Sex         enums.Sex        `json:"sex" binding:"required"`
	PhoneNumber string           `json:"phoneNumber" binding:"required"`
	Address     string           `json:"address" binding:"required"`
	Email       *string          `json:"email" binding:"omitempty,email"`
	BloodType   *enums.BloodType `json:"bloodType"`
	Allergies   *string          `json:"allergies"`
}

type ScheduleAppointmentDTO struct {
	DoctorID uint      `json:"doctorId" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
}
