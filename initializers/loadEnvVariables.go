package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
	}
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
