package database

import (
	"log"

	"github.com/andresidrim/cesupa-hospital/env"
	"github.com/andresidrim/cesupa-hospital/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(env.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected successfully")

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.Pacient{},
	); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	log.Println("Database auto migrated")
}
