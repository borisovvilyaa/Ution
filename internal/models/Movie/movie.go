package models

// Movie represents a single movie structure returned from TMDB.
type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
	PosterPath  string  `json:"poster_path"`
}

// MoviesResponse represents the response structure from TMDB for multiple movies.
type MoviesResponse struct {
	Results []Movie `json:"results"`
}
