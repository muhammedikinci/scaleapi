package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type movieApi interface {
	GetAllMovies() ([]models.Movie, error)
}

type movieHandler struct {
	api movieApi
}

func NewMovieHandler(movieApi movieApi) *movieHandler {
	return &movieHandler{
		api: movieApi,
	}
}

func (mh movieHandler) GetAllMovies(c echo.Context) error {
	movies, err := mh.api.GetAllMovies()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movies)
}
