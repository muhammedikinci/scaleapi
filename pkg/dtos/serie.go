package dtos

import "github.com/muhammedikinci/scaleapi/pkg/models"

type SerieRequest struct {
	ID          int             `json:"-"`
	Title       string          `json:"title"`
	Image       string          `json:"image"`
	Description string          `json:"description"`
	Rating      float32         `json:"rating"`
	ReleaseDate string          `json:"release_date"`
	Director    string          `json:"director"`
	Writer      string          `json:"writer"`
	Stars       string          `json:"stars"`
	IMDBID      string          `json:"imdb_id"`
	Year        int             `json:"year"`
	Genre       string          `json:"genre"`
	Seasons     []models.Season `json:"-"`
}

func (sr SerieRequest) Validate() (string, bool) {
	if sr.Title == "" || sr.Description == "" {
		return ErrEmptyTitleAndDescription, false
	}

	return "", true
}

func (sr SerieRequest) GetSerieModel() models.Serie {
	return models.Serie(sr)
}
