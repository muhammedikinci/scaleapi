package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type movieApi interface {
	GetAllMovies() ([]models.Movie, error)
	AddMovie(movieDto dtos.MovieRequest) (models.Movie, string, error)
	FindById(id int) (models.Movie, error)
	Filter(title string, genre string) ([]models.Movie, error)
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
	m := new(dtos.MovieRequest)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	movie, message, err := mh.api.AddMovie(*m)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if message != "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: message,
		})
	}

	return c.JSON(http.StatusOK, movie)
}

func (mh movieHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	movie, err := mh.api.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, movie)
}

func (mh movieHandler) Filter(c echo.Context) error {
	title := c.QueryParam("title")
	genre := c.QueryParam("genre")

	movies, err := mh.api.Filter(title, genre)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movies)
}
