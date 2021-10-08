package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	var err error = godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error lodaing .env file")
	}

	return os.Getenv(key)
}
