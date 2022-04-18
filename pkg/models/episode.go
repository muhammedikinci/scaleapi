package models

import "time"

type Episode struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    string    `json:"duration"`
	Year        int       `json:"year"`
	SeasonID    uint      `json:"season_id"`
	Audio       string    `json:"audio"`
	Subtitles   string    `json:"subtitles"`
}
