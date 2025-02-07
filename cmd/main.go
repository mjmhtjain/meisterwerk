package main

import (
	"log"
	"os"

	"github.com/mjmhtjain/meisterwerk/internal/router"
)

func main() {
	// Initialize routes
	router := router.NewRouter().Setup()

	// Start the server
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
