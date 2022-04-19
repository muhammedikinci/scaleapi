package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type serieApi interface {
	GetAllSeries() ([]models.Serie, error)
	AddSerie(models.Serie) (models.Serie, error)
	FindById(id int) (models.Serie, error)
	Filter(title string, genre string) ([]models.Serie, error)
}

type serieHandler struct {
	api serieApi
}

func NewSerieHandler(serieApi serieApi) *serieHandler {
	return &serieHandler{
		api: serieApi,
	}
}

func (sh serieHandler) GetAllSeries(c echo.Context) error {
	series, err := sh.api.GetAllSeries()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, series)
}

func (sh serieHandler) AddSerie(c echo.Context) error {
	s := new(models.Serie)

	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	serie, err := sh.api.AddSerie(*s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, serie)
}

func (sh serieHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	serie, err := sh.api.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, serie)
}

func (sh serieHandler) Filter(c echo.Context) error {
	title := c.QueryParam("title")
	genre := c.QueryParam("genre")

	series, err := sh.api.Filter(title, genre)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, series)
}