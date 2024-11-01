package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"ution/internal/models/Errors"

	"github.com/joho/godotenv"
	_ "ution/internal/models/Errors" // Adjust the import path to your actual project structure
	"ution/internal/models/Movie"
	// Adjust the import path to your actual project structure
	errors "ution/internal/models/Errors/database" // Adjust the import path to your actual project structure
)

const baseURL = "https://api.themoviedb.org/3"

// FetchMovieByID retrieves a specific movie by its ID from TMDB API.
func FetchMovieByID(id string) (models.Movie, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return models.Movie{}, Errors.ErrEnvLoad
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("API_TMDB")
	if apiKey == "" {
		return models.Movie{}, errors.ErrApiKeyMissing
	}

	// Form the URL request for the movie by ID
	url := fmt.Sprintf("%s/movie/%s?api_key=%s", baseURL, id, apiKey)

	// Perform GET request
	response, err := http.Get(url)
	if err != nil {
		return models.Movie{}, errors.ErrRequestFailed(err)
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return models.Movie{}, errors.ErrInvalidStatusCode(response.StatusCode)
	}

	// Decode JSON response
	var movie models.Movie
	if err := json.NewDecoder(response.Body).Decode(&movie); err != nil {
		return models.Movie{}, errors.ErrResponseDecode(err)
	}

	return movie, nil
}

// FetchPopularMovies retrieves popular movies from the TMDB API.
func FetchPopularMovies() ([]models.Movie, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, Errors.ErrEnvLoad
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("API_TMDB")
	if apiKey == "" {
		return nil, errors.ErrApiKeyMissing
	}

	// Form the URL request
	url := fmt.Sprintf("%s/movie/popular?api_key=%s", baseURL, apiKey)

	// Perform GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrRequestFailed(err)
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return nil, errors.ErrInvalidStatusCode(response.StatusCode)
	}

	// Decode JSON response
	var moviesResponse models.MoviesResponse
	if err := json.NewDecoder(response.Body).Decode(&moviesResponse); err != nil {
		return nil, errors.ErrResponseDecode(err)
	}

	return moviesResponse.Results, nil
}

// SearchMovies retrieves movies based on the search query.
func SearchMovies(query string) ([]models.Movie, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, Errors.ErrEnvLoad
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("API_TMDB")
	if apiKey == "" {
		return nil, errors.ErrApiKeyMissing
	}

	// Form the URL request for searching movies
	url := fmt.Sprintf("%s/search/movie?api_key=%s&query=%s", baseURL, apiKey, query)

	// Perform GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.ErrRequestFailed(err)
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return nil, errors.ErrInvalidStatusCode(response.StatusCode)
	}

	// Decode JSON response
	var moviesResponse models.MoviesResponse
	if err := json.NewDecoder(response.Body).Decode(&moviesResponse); err != nil {
		return nil, errors.ErrResponseDecode(err)
	}
	return moviesResponse.Results, nil
}
