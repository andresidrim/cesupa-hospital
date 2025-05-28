package pacients

import (
	"testing"
	"time"

	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Auto-migrate the Pacient model
	err = db.AutoMigrate(&models.Pacient{})
	assert.NoError(t, err)

	t.Log("Test DB setup complete")

	return db
}

func TestServiceGetAll(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Arrange: create test pacients
	pacients := []models.Pacient{
		{Name: "John Doe", BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), CPF: "111", Sex: "male", PhoneNumber: "123", Address: "123 Street"},
		{Name: "Jane Smith", BirthDate: time.Date(1995, 5, 10, 0, 0, 0, 0, time.UTC), CPF: "222", Sex: "female", PhoneNumber: "456", Address: "456 Avenue"},
	}

	for _, p := range pacients {
		err := service.Create(&p)
		assert.NoError(t, err)
		t.Logf("Created pacient: %s, BirthDate: %s", p.Name, p.BirthDate.Format("2006-01-02"))
	}

	// Define test cases
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

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetAll(tt.filterName, tt.filterAge)
			assert.NoError(t, err)
			t.Logf("GetAll with name '%s' and age '%s' returned %d pacients", tt.filterName, tt.filterAge, len(result))

			for _, r := range result {
				t.Logf("Pacient: %s, BirthDate: %s", r.Name, r.BirthDate.Format("2006-01-02"))
			}

			assert.Len(t, result, tt.wantCount)

			// If expecting at least 1 result, check first's name
			if tt.wantCount > 0 {
				assert.Equal(t, tt.wantFirst, result[0].Name)
			}
		})
	}
}
