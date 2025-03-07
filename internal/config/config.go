package config

import (
	"os"
	"test_project_2/internal/models"

	"github.com/joho/godotenv"
)

func LoadConfig() (*models.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &models.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}, nil
}
