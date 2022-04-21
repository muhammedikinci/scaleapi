package models

type Serie struct {
	ID             int      `json:"id"`
	Title          string   `json:"title"`
	Image          string   `json:"image"`
	Description    string   `json:"description"`
	Rating         float32  `json:"rating"`
	ReleaseDate    string   `json:"release_date"`
	Director       string   `json:"director"`
	Writer         string   `json:"writer"`
	Stars          string   `json:"stars"`
	IMDBID         string   `json:"imdb_id"`
	Year           int      `json:"year"`
	Genre          string   `json:"genre"`
	Seasons        []Season `json:"seasons"`
	SerieFavorites []Serie  `json:"-" gorm:"many2many:user_series"`
}
