package pacients

import (
	"testing"
	"time"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.Exec("PRAGMA foreign_keys = ON;")

	err = db.AutoMigrate(
		&models.User{},
		&models.Appointment{},
		&models.Pacient{},
	)
	assert.NoError(t, err)

	t.Log("Test DB setup complete")

	return db
}

func TestServiceGet(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Arrange: crie um pacient válido para o cenário de sucesso
	existing := models.Pacient{
		Name:        "Fetch Me",
		BirthDate:   time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC),
		CPF:         "55566677788",
		Sex:         "male",
		PhoneNumber: "+5511999999999",
		Address:     "Rua Teste, 123",
	}
	assert.NoError(t, service.Create(&existing))

	tests := []struct {
		name      string
		id        uint64
		wantError bool
		wantName  string
	}{
		{
			name:      "found",
			id:        uint64(existing.ID),
			wantError: false,
			wantName:  "Fetch Me",
		},
		{
			name:      "not_found",
			id:        9999,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Get(tt.id)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantName, got.Name)
			}
		})
	}
}

func TestServiceGetAll(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	pacients := []models.Pacient{
		{Name: "John Doe", BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), CPF: "111", Sex: "male", PhoneNumber: "123", Address: "123 Street"},
		{Name: "Jane Smith", BirthDate: time.Date(1995, 5, 10, 0, 0, 0, 0, time.UTC), CPF: "222", Sex: "female", PhoneNumber: "456", Address: "456 Avenue"},
	}

	for _, p := range pacients {
		err := service.Create(&p)
		assert.NoError(t, err)
		t.Logf("Created pacient: %s, BirthDate: %s", p.Name, p.BirthDate.Format("2006-01-02"))
	}

	tests := []struct {
		name       string
		filterName string
		filterAge  string
		wantCount  int
		wantFirst  string
	}{
		{
			name:       "no filters",
			filterName: "",
			filterAge:  "",
			wantCount:  2,
			wantFirst:  "John Doe",
		},
		{
			name:       "filter by name Jane",
			filterName: "Jane",
			filterAge:  "",
			wantCount:  1,
			wantFirst:  "Jane Smith",
		},
		{
			name:       "filter by non-matching age",
			filterName: "",
			filterAge:  "1000",
			wantCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetAll(tt.filterName, tt.filterAge)
			assert.NoError(t, err)
			t.Logf("GetAll with name '%s' and age '%s' returned %d pacients", tt.filterName, tt.filterAge, len(result))

			for _, r := range result {
				t.Logf("Pacient: %s, BirthDate: %s", r.Name, r.BirthDate.Format("2006-01-02"))
			}

			assert.Len(t, result, tt.wantCount)

			if tt.wantCount > 0 {
				assert.Equal(t, tt.wantFirst, result[0].Name)
			}
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	pacient := models.Pacient{
		Name:        "John Doe",
		BirthDate:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		CPF:         "123",
		Sex:         "male",
		PhoneNumber: "+123456789",
		Address:     "123 Street",
	}
	err := service.Create(&pacient)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		id            uint64
		updateData    models.Pacient
		expectedError bool
	}{
		{
			name: "successful update",
			id:   uint64(pacient.ID),
			updateData: models.Pacient{
				Name: "Updated Name",
			},
			expectedError: false,
		},
		{
			name: "pacient not found",
			id:   9999,
			updateData: models.Pacient{
				Name: "Nonexistent",
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Update(tt.id, &tt.updateData)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.id == uint64(pacient.ID) {
				updated, err := service.Get(tt.id)
				assert.NoError(t, err)
				assert.Equal(t, "Updated Name", updated.Name)
			}
		})
	}
}

func TestServiceDelete(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	pacient := models.Pacient{
		Name:        "John Doe",
		BirthDate:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		CPF:         "123",
		Sex:         "male",
		PhoneNumber: "+123456789",
		Address:     "123 Street",
	}
	err := service.Create(&pacient)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		id            uint64
		expectedError error
	}{
		{
			name:          "successful delete",
			id:            uint64(pacient.ID),
			expectedError: nil,
		},
		{
			name:          "non-existent pacient",
			id:            9999,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Delete(tt.id)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedError == nil {
				_, getErr := service.Get(tt.id)
				assert.ErrorIs(t, getErr, gorm.ErrRecordNotFound)
			}
		})
	}
}

func TestServiceScheduleAppointment(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	pacient := models.Pacient{
		Name:        "John Doe",
		BirthDate:   time.Now().AddDate(-30, 0, 0),
		CPF:         "12345678900",
		Sex:         "male",
		PhoneNumber: "+123456789",
		Address:     "123 Street",
	}
	assert.NoError(t, service.Create(&pacient))

	doctor := models.User{
		Name: "Dr. Smith",
		CPF:  "98765432100",
		Role: enums.Doctor,
	}
	assert.NoError(t, db.Create(&doctor).Error)

	tests := []struct {
		name          string
		appointment   models.Appointment
		expectedError bool
	}{
		{
			name: "valid appointment",
			appointment: models.Appointment{
				PacientID: pacient.ID,
				UserID:    doctor.ID,
				Date:      time.Now().AddDate(0, 0, 1),
			},
			expectedError: false,
		},
		{
			name: "missing DoctorID",
			appointment: models.Appointment{
				PacientID: pacient.ID,
				Date:      time.Now().AddDate(0, 0, 1),
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.ScheduleAppointment(&tt.appointment)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				var appt models.Appointment
				err = db.First(&appt, tt.appointment.ID).Error
				assert.NoError(t, err)
				assert.Equal(t, tt.appointment.PacientID, appt.PacientID)
			}
		})
	}
}
