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
	db, err := database.NewClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize routes
	router := router.NewRouter(db).Setup()

	// Start the server
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
