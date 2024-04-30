package server

import (
	"auth_service/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	return router
}

func LoadConfig() Config {
	err := godotenv.Load("deployments/.env")
	if err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		Port: os.Getenv("SERVER_PORT"),
		Database: database.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}
}
