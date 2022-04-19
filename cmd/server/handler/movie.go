package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type movieApi interface {
	GetAllMovies() ([]models.Movie, error)
	AddMovie(models.Movie) (models.Movie, error)
	FindById(id int) (models.Movie, error)
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movies)
}

func (mh movieHandler) AddMovie(c echo.Context) error {
	m := new(models.Movie)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	movie, err := mh.api.AddMovie(*m)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movie)
}

func (mh movieHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	movie, err := mh.api.FindById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movie)
}
