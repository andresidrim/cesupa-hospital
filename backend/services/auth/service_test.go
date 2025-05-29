package auth

import (
	"os"
	"testing"

	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/andresidrim/cesupa-hospital/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{})
	assert.NoError(t, err)
	return db
}

func TestAuthServiceRegister(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	user := &models.User{
		Name:     "Jane Doe",
		CPF:      "99988877766",
		Password: "mypassword",
		Role:     "doctor",
	}

	err := service.Register(user)
	assert.NoError(t, err)

	var saved models.User
	err = db.Where("cpf = ?", "99988877766").First(&saved).Error
	assert.NoError(t, err)

	assert.NotEqual(t, "mypassword", saved.Password)
	assert.NoError(t, utils.CheckPassword(saved.Password, "mypassword"))
}

func TestAuthServiceLogin(t *testing.T) {
	os.Setenv("SECRET_KEY", "test-secret")

	db := setupTestDB(t)
	service := NewService(db)

	hashed, err := utils.HashPassword("secret123")
	assert.NoError(t, err)
	u := models.User{
		Name:     "John Doe",
		CPF:      "11122233344",
		Password: hashed,
		Role:     "admin",
	}
	assert.NoError(t, db.Create(&u).Error)

	tests := []struct {
		name        string
		cpf, pass   string
		wantErr     bool
		errContains string
	}{
		{
			name:        "user not found",
			cpf:         "doesnotexist",
			pass:        "whatever",
			wantErr:     true,
			errContains: "user not found",
		},
		{
			name:        "incorrect password",
			cpf:         "11122233344",
			pass:        "wrongpass",
			wantErr:     true,
			errContains: "incorrect password",
		},
		{
			name:    "success",
			cpf:     "11122233344",
			pass:    "secret123",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := service.Login(tt.cpf, tt.pass)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
