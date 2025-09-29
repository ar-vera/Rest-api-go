package db

import (
	"example.com/mod/internal/calculationService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5433 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&calculationService.Calculation{}); err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	return db, nil
}
