package models

type Season struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ReleaseDate string    `json:"release_date"`
	Year        int       `json:"year"`
	SerieID     uint      `json:"serie_id"`
	Episodes    []Episode `json:"episodes"`
}
