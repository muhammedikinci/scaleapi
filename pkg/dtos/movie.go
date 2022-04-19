package dtos

import "github.com/muhammedikinci/scaleapi/pkg/models"

type MovieRequest struct {
	ID          int     `json:"-"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	Director    string  `json:"director"`
	Writer      string  `json:"writer"`
	Stars       string  `json:"stars"`
	Duration    string  `json:"duration"`
	IMDBID      string  `json:"imdb_id"`
	Year        int     `json:"year"`
	Genre       string  `json:"genre"`
}

func (mr MovieRequest) Validate() (string, bool) {
	if mr.Title == "" || mr.Description == "" {
		return "Title and Description cannot be empty", false
	}

	return "", true
}

func (mr MovieRequest) GetMovieModel() models.Movie {
	return models.Movie(mr)
}
