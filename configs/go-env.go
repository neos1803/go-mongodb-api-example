package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MongoEnvUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func DatabaseName() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("MONGODB_DB")
}
