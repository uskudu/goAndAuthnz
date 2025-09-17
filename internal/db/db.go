package db

import (
	"authnz/internal/userService"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	err = DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatalf("failed while migrating database: %v", err)
	}
	return DB, nil
}
