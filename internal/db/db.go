package db

import (
	"authnz/internal/userService"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectToDB() {
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connnect to db")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&userService.User{})
	if err != nil {
		fmt.Errorf("error while automirating: %v", err)
	}
}
