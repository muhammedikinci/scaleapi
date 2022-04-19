package api

import (
	"strings"

	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type movieRepository interface {
	GetAllMovies() ([]models.Movie, error)
	AddMovie(models.Movie) (models.Movie, error)
	FindById(id int) (models.Movie, error)
	Filter(title string, genre string) ([]models.Movie, error)
}

type MovieAPI struct {
	Repository movieRepository
}

func (ma MovieAPI) GetAllMovies() ([]models.Movie, error) {
	return ma.Repository.GetAllMovies()
}

func (ma MovieAPI) AddMovie(movieDto dtos.MovieRequest) (models.Movie, string, error) {
	if v, ok := movieDto.Validate(); !ok {
		return models.Movie{}, v, nil
	}

	res, err := ma.Repository.AddMovie(movieDto.GetMovieModel())

	if err != nil {
		return models.Movie{}, "", err
	}

	return res, "", nil
}

func (ma MovieAPI) FindById(id int) (models.Movie, error) {
	return ma.Repository.FindById(id)
}

func (ma MovieAPI) Filter(title, genre string) ([]models.Movie, error) {
	return ma.Repository.Filter(strings.ToLower(title), strings.ToLower(genre))
}
