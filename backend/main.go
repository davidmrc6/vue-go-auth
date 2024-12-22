package main

import (
	"backend/models"
	"backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Create new Gin instance.
	r := gin.Default()

	// Load .env file and load database config values.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	config := models.Config{
		Host:				os.Getenv("DB_HOST"),
		Port:				os.Getenv("DB_PORT"),
		User:				os.Getenv("DB_USER"),
		Password:		os.Getenv("DB_PASSWORD"),
		DBName:			os.Getenv("DB_NAME"),
		SSLMode:		os.Getenv("DB_SSLMODE"),
	}

	// Initialize database.
	models.InitDB(config)

	// Load routes.
	routes.AuthRoutes(r)

	// Run server
	r.Run(":8080")
}
