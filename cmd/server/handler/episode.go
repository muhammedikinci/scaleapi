package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type episodeApi interface {
	AddEpisode(sr dtos.EpisodeRequest) (models.Episode, string, error)
	FindById(id int) (models.Episode, error)
	FindAllEpisodesInSeason(seasonId int) ([]models.Episode, error)
}

type episodeHandler struct {
	api episodeApi
}

func NewEpisodeHandler(episodeApi episodeApi) *episodeHandler {
	return &episodeHandler{
		api: episodeApi,
	}
}

func (eh episodeHandler) AddEpisode(c echo.Context) error {
	s := new(dtos.EpisodeRequest)

	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	episode, message, err := eh.api.AddEpisode(*s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if message != "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: message,
		})
	}

	return c.JSON(http.StatusOK, episode)
}

func (eh episodeHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "id parameter is not valid",
		})
	}

	episode, err := eh.api.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, episode)
}

func (eh episodeHandler) FindAllEpisodesInSeason(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("seasonId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "seasonId parameter is not valid",
		})
	}

	episode, err := eh.api.FindAllEpisodesInSeason(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, episode)
}
