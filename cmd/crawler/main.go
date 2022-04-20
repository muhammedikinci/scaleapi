package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/muhammedikinci/scaleapi/pkg/api"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/repository/gormrepo"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=postgres dbname=scaleflix", "Database Connection String")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	repository, err := gormrepo.NewRepository(errorLog, infoLog, dsn)

	if err != nil {
		panic(err)
	}

	movieApi := api.MovieAPI{
		Repository: repository.Movie,
	}

	serieApi := api.SerieAPI{
		Repository: repository.Serie,
	}

	seasonApi := api.SeasonAPI{
		Repository: repository.Season,
	}

	episodeApi := api.EpisodeAPI{
		Repository: repository.Episode,
	}

	getSeries(serieApi, seasonApi, episodeApi)
	getMovies(movieApi)
}

func getSeries(serieApi api.SerieAPI, seasonApi api.SeasonAPI, episodeApi api.EpisodeAPI) {
	for i := 1; i < 10; i++ {
		resp, err := http.Get(fmt.Sprintf("https://api.tvmaze.com/shows/%d", i))

		if err != nil {
			panic(err)
		}

		serie := &Serie{}

		json.NewDecoder(resp.Body).Decode(serie)
		resp.Body.Close()

		savedSerie, _, err := serieApi.AddSerie(dtos.SerieRequest{
			Title:       serie.Title,
			Image:       serie.Image.Original,
			Description: serie.Description,
			Rating:      serie.Rating.Average,
			ReleaseDate: serie.ReleaseDate,
			Director:    serie.Director,
			Writer:      serie.Writer,
			Stars:       serie.Stars,
			IMDBID:      serie.IMDBID.IMDB,
			Year:        serie.Year,
			Genre:       strings.Join(serie.Genre, ","),
		})

		if err != nil {
			panic(err)
		}

		seasons := new([]Season)

		resp, err = http.Get(fmt.Sprintf("https://api.tvmaze.com/shows/%d/seasons", i))

		if err != nil {
			panic(err)
		}

		json.NewDecoder(resp.Body).Decode(seasons)

		resp.Body.Close()

		for si, v := range *seasons {
			seasonTitle := v.Title

			if seasonTitle == "" {
				seasonTitle = fmt.Sprintf("Season %d", si)
			}

			savedSeason, m, err := seasonApi.AddSeason(dtos.SeasonRequest{
				Title:       seasonTitle,
				Image:       v.Image.Original,
				Description: v.Description,
				Rating:      0,
				ReleaseDate: v.ReleaseDate,
				Year:        0,
				SerieID:     uint(savedSerie.ID),
			})

			if err != nil || m != "" {
				panic(err)
			}

			episodes := new([]Episode)

			resp, err = http.Get(fmt.Sprintf("https://api.tvmaze.com/seasons/%d/episodes", v.Id))

			if err != nil {
				panic(err)
			}

			json.NewDecoder(resp.Body).Decode(episodes)

			resp.Body.Close()

			for _, e := range *episodes {
				_, _, err := episodeApi.AddEpisode(dtos.EpisodeRequest{
					Title:       e.Title,
					Image:       e.Image.Original,
					Description: e.Description,
					Rating:      e.Rating.Average,
					ReleaseDate: e.ReleaseDate,
					Duration:    e.Duration,
					Year:        e.Year,
					SeasonID:    uint(savedSeason.ID),
					Audio:       e.Audio,
					Subtitles:   e.Subtitles,
				})

				if err != nil {
					panic(err)
				}
			}
		}
	}

}

func getMovies(movieApi api.MovieAPI) {
	m := []string{
		"https://www.omdbapi.com/?t=harry%20potter&apikey=dd1debeb",
		"https://www.omdbapi.com/?t=Interstellar&apikey=dd1debeb",
		"https://www.omdbapi.com/?t=edge%20of%20tomorrow&apikey=dd1debeb",
		"https://www.omdbapi.com/?t=blade%20runner&apikey=dd1debebb",
		"https://www.omdbapi.com/?t=amazing%20spider&apikey=dd1debeb",
	}

	for _, v := range m {
		resp, err := http.Get(v)

		if err != nil {
			panic(err)
		}

		movie := &Movie{}

		json.NewDecoder(resp.Body).Decode(movie)

		year, _ := strconv.Atoi(movie.Year)
		rating, _ := strconv.ParseFloat(movie.Rating, 64)

		_, _, err = movieApi.AddMovie(dtos.MovieRequest{
			Title:       movie.Title,
			Image:       movie.Image,
			Description: movie.Description,
			Rating:      float32(rating),
			ReleaseDate: movie.ReleaseDate,
			Director:    movie.Director,
			Writer:      movie.Writer,
			Stars:       movie.Stars,
			Duration:    movie.Duration,
			IMDBID:      movie.IMDBID,
			Year:        year,
			Genre:       movie.Genre,
		})

		if err != nil {
			panic(err)
		}

		resp.Body.Close()
	}
}
