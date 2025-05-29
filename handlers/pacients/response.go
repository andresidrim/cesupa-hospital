package pacients

import "time"

// PacientResponse é o payload retornado em /pacients e /pacients/{id}
type PacientResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	BirthDate   time.Time `json:"birthDate"`
	CPF         string    `json:"cpf"`
	Sex         string    `json:"sex"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Email       *string   `json:"email,omitempty"`
	BloodType   *string   `json:"bloodType,omitempty"`
	Allergies   *string   `json:"allergies,omitempty"`
}

// AppointmentResponse é o payload retornado em POST /pacients/{id}/appointments
type AppointmentResponse struct {
	ID        uint      `json:"id"`
	PacientID uint      `json:"pacientId"`
	UserID    uint      `json:"doctorId"`
	Date      time.Time `json:"date"`
}

// ErrorResponse representa um erro comum
type ErrorResponse struct {
	Message string `json:"message"`
}
