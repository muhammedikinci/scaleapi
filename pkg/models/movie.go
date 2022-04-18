package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	Director    string    `json:"director"`
	Writer      string    `json:"writer"`
	Stars       string    `json:"stars"`
	Duration    string    `json:"duration"`
	IMDBID      string    `json:"imdb_id"`
	Year        int       `json:"year"`
	Genre       []Genre   `gorm:"many2many:movie_genres"`
}
