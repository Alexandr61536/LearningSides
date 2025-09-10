package main

import (
	"log"

	"github.com/joho/godotenv"

	"backend/internal/database"
	"backend/internal/transport"
)

// @title LearningSides
// @version 1.0
// @description API for LearningSides

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("Starting...")
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	database.InitDB()
	transport.Listener()
}
