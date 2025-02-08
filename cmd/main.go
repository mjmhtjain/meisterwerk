package main

import (
	"log"
	"os"

	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/database"
	"github.com/mjmhtjain/meisterwerk/internal/router"
)

func main() {
	// Initialize database
	dbConfig := config.NewDatabaseConfig()
	db, err := database.NewDBClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database or run migrations: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize routes
	router := router.NewRouter().Setup()

	// Start the server
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
