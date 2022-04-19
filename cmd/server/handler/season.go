package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type seasonApi interface {
	AddSeason(sr dtos.SeasonRequest) (models.Season, string, error)
	FindById(id int) (models.Season, error)
	FindAllSeasonsInSerie(serieId int) ([]models.Season, error)
}

type seasonHandler struct {
	api seasonApi
}

func NewSeasonHandler(seasonApi seasonApi) *seasonHandler {
	return &seasonHandler{
		api: seasonApi,
	}
}

func (sh seasonHandler) AddSeason(c echo.Context) error {
	s := new(dtos.SeasonRequest)

	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	season, message, err := sh.api.AddSeason(*s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if message != "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: message,
		})
	}

	return c.JSON(http.StatusOK, season)
}

func (sh seasonHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "id parameter is not valid",
		})
	}

	season, err := sh.api.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, season)
}

func (sh seasonHandler) FindAllSeasonsInSerie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("serieId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "serieId parameter is not valid",
		})
	}

	season, err := sh.api.FindAllSeasonsInSerie(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, season)
}
