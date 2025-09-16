package initializers

import (
	"authnz/internal/db"
	"authnz/internal/services"
	"fmt"
)

func SyncDB() {
	err := db.DB.AutoMigrate(&services.User{})
	if err != nil {
		fmt.Errorf("error while automirating: %v", err)
	}
}
