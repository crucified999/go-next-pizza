package main

import (
	"log"

	"github.com/go-next-pizza/internal/app/config"
	"github.com/go-next-pizza/internal/app/container"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Load configuration
	cfg := config.Load()

	// Create container
	app, err := container.New(cfg)
	if err != nil {
		log.Fatal("Failed to create container:", err)
	}
	defer app.DB.Close()

	// Start server
	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := app.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}