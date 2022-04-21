package dtos

import "github.com/muhammedikinci/scaleapi/pkg/models"

const ErrEmptyTitleAndDescription = "Title and Description cannot be empty"

type EpisodeRequest struct {
	ID          int     `json:"-"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	Duration    string  `json:"duration"`
	Year        int     `json:"year"`
	SeasonID    uint    `json:"season_id"`
	Audio       string  `json:"audio"`
	Subtitles   string  `json:"subtitles"`
}

func (er EpisodeRequest) Validate() (string, bool) {
	if er.Title == "" || er.Description == "" {
		return ErrEmptyTitleAndDescription, false
	}

	return "", true
}

func (er EpisodeRequest) GetEpisodeModel() models.Episode {
	return models.Episode(er)
}
