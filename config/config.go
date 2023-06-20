package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

// GetEnv returns the value of the specified environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
