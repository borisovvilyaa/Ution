package api

import (
	"log"
	"net/http"

	"ution/internal/api/handlers" // Adjust the import path to match your project structure
)

// CORS middleware to handle CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins or specify your frontend origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Specify allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Specify allowed headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r) // Call the next handler
	})
}

// StartServer initializes the HTTP server with the defined routes.
func StartServer() {
	// Define the routes for the API
	http.Handle("/api/register", corsMiddleware(http.HandlerFunc(handlers.RegisterUser)))
	http.Handle("/api/movies/id", corsMiddleware(http.HandlerFunc(handlers.FetchMovieByID))) // Get movie by ID
	http.Handle("/api/movies/n", corsMiddleware(http.HandlerFunc(handlers.FetchNMovies)))    // Get N movies

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
