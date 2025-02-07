package main

import (
	"log"

	"github.com/mjmhtjain/meisterwerk/internal/router"
)

func main() {
	// Initialize routes
	router := router.NewRouter().Setup()

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
