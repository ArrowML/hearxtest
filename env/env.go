package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(env + ".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}
