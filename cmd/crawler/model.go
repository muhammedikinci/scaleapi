package main

type Serie struct {
	Title       string    `json:"name"`
	Image       Image     `json:"image"`
	Description string    `json:"summary"`
	Rating      Rating    `json:"rating"`
	ReleaseDate string    `json:"premiered"`
	Director    string    `json:"director"`
	Writer      string    `json:"writer"`
	Stars       string    `json:"stars"`
	IMDBID      Externals `json:"externals"`
	Year        int       `json:"year"`
	Genre       []string  `json:"genres"`
	Seasons     []Season  `json:"-"`
}

type Image struct {
	Original string `json:"original"`
}

type Externals struct {
	IMDB string `json:"imdb"`
}

type Rating struct {
	Average float32 `json:"average"`
}

type Season struct {
	Id          int       `json:"id"`
	Title       string    `json:"name"`
	Image       Image     `json:"image"`
	Description string    `json:"summary"`
	ReleaseDate string    `json:"premiereDate"`
	Episodes    []Episode `json:"-"`
}

type Episode struct {
	Title       string `json:"name"`
	Image       Image  `json:"image"`
	Description string `json:"summary"`
	Rating      Rating `json:"rating"`
	ReleaseDate string `json:"premiereDate"`
	Duration    string `json:"duration"`
	Year        int    `json:"year"`
	SeasonID    uint   `json:"season_id"`
	Audio       string `json:"audio"`
	Subtitles   string `json:"subtitles"`
}

type Movie struct {
	Title       string `json:"Title"`
	Image       string `json:"Poster"`
	Description string `json:"Plot"`
	Rating      string `json:"imdbRating"`
	ReleaseDate string `json:"Released"`
	Director    string `json:"Director"`
	Writer      string `json:"Writer"`
	Stars       string `json:"Actors"`
	Duration    string `json:"Runtime"`
	IMDBID      string `json:"imdbID"`
	Year        string `json:"Year"`
	Genre       string `json:"Genre"`
}
