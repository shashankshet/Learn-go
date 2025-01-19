package main

import (
	"log"

	"github.com/shashankshet/go_api/db"
	"github.com/shashankshet/go_api/router"
)

func main() {
	// Initialize database connection
	db.InitDB()
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	// Set up the router
	r := router.SetupRouter()

	// Start the server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
