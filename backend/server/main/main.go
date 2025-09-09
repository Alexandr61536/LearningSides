package main

import (
	"fmt"

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
	fmt.Println("Starting...")
	transport.Listener()
}
