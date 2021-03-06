package dtos

import "github.com/muhammedikinci/scaleapi/pkg/models"

const ErrEmptyTitle = "Title cannot be empty"

type SeasonRequest struct {
	ID          int              `json:"-"`
	Title       string           `json:"title"`
	Image       string           `json:"image"`
	Description string           `json:"description"`
	Rating      float32          `json:"rating"`
	ReleaseDate string           `json:"release_date"`
	Year        int              `json:"year"`
	SerieID     uint             `json:"serie_id"`
	Episodes    []models.Episode `json:"-"`
}

func (sr SeasonRequest) Validate() (string, bool) {
	if sr.Title == "" {
		return ErrEmptyTitle, false
	}

	return "", true
}

func (sr SeasonRequest) GetSeasonModel() models.Season {
	return models.Season(sr)
}
