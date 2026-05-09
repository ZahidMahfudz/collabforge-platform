package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		Logger.Fatalf("Error loading .env file: %v", err)
	}

	Logger.Infof("Environment variables loaded successfully")
}

func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	Logger.Fatalf("Environment variable %s not set", key)
	return ""

}