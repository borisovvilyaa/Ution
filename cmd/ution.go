package main

import (
	"log"
	"ution/internal/storage"

	"ution/internal/api" // Adjust the import path to match your project structure
)

func main() {
	// Initialize the database connection
	if err := storage.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Start the API server
	api.StartServer()
}
