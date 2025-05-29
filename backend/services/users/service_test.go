package users

import (
	"testing"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.User{}, &models.Appointment{})
	assert.NoError(t, err)

	return db
}

func createTestUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "Alice", CPF: "12345678901", Role: enums.Admin},
		{Name: "Bob", CPF: "12345678902", Role: enums.Doctor},
		{Name: "Carol", CPF: "12345678903", Role: enums.Receptionist},
	}
	for _, user := range users {
		db.Create(&user)
	}
}

func TestServiceGetAll(t *testing.T) {
	db := setupTestDB(t)
	createTestUsers(db)
	service := NewService(db)

	tests := []struct {
		name        string
		filterRoles []enums.Role
		expected    int
	}{
		{"no_filter", nil, 3},
		{"filter_admin", []enums.Role{enums.Admin}, 1},
		{"filter_doctor_and_admin", []enums.Role{enums.Doctor, enums.Admin}, 2},
		{"filter_none", []enums.Role{"unknown"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := service.GetAll(tt.filterRoles)
			assert.NoError(t, err)
			assert.Len(t, users, tt.expected)
		})
	}
}

func TestServiceGet(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Arrange: crie um usuário para o cenário de sucesso
	existing := models.User{
		Name: "Tester",
		CPF:  "99999999999",
		Role: enums.Doctor,
	}
	assert.NoError(t, db.Create(&existing).Error)

	tests := []struct {
		name      string
		id        uint64
		wantError bool
		wantName  string
		wantRole  enums.Role
	}{
		{
			name:      "found",
			id:        uint64(existing.ID),
			wantError: false,
			wantName:  "Tester",
			wantRole:  enums.Doctor,
		},
		{
			name:      "not_found",
			id:        9999,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.Get(tt.id)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantName, user.Name)
				assert.Equal(t, tt.wantRole, user.Role)
			}
		})
	}
}
