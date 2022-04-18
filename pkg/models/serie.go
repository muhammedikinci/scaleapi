package models

import "time"

type Serie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	Director    string    `json:"director"`
	Writer      string    `json:"writer"`
	Stars       string    `json:"stars"`
	IMDBID      string    `json:"imdb_id"`
	Year        int       `json:"year"`
	Genre       []Genre   `gorm:"many2many:serie_genres" json:"genres"`
	Seasons     []Season  `json:"seasons"`
}
