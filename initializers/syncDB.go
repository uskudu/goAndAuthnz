package initializers

import (
	"authnz/models"
	"fmt"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Errorf("error while automirating: %v", err)
	}
}
