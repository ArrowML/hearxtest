package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	cwd, _ := os.Getwd()
	fmt.Print(cwd)

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(cwd + "/" + env + ".env")
	fmt.Print(err)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}
