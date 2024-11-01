package handlers

import (
	"encoding/json"
	"net/http"
	"ution/internal/storage" // Adjust the import path as needed
)

// HandleFetchPopularMovies handles the request for popular movies.
func HandleFetchPopularMovies(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Call the storage function to fetch popular movies
	movies, err := storage.FetchPopularMovies() // Fetching movies here
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the movies as a JSON response
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SearchMovies handles the request for searching movies by title.
func SearchMovies(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the search query from the URL parameters
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	// Call the storage function to search for movies
	movies, err := storage.SearchMovies(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the movies as a JSON response
	json.NewEncoder(w).Encode(movies)
}

// FetchMovieByID handles the request for fetching a specific movie by ID.
func FetchMovieByID(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the movie ID from the URL parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	// Fetch the movie by ID from the storage
	movie, err := storage.FetchMovieByID(id) // You need to implement this in the storage package
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the movie as a JSON response
	json.NewEncoder(w).Encode(movie)
}

// FetchNMovies handles the request for fetching N movies (could be similar to popular).
func FetchNMovies(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// You could implement this based on your requirements, e.g., a limit on the number of popular movies
	// For simplicity, let's just call the FetchPopularMovies function here
	movies, err := storage.FetchPopularMovies() // Adjust this if you want to implement a limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the movies as a JSON response
	json.NewEncoder(w).Encode(movies)
}
