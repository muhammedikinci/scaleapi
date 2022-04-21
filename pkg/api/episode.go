package api

import (
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

//go:generate mockgen -source $GOFILE -destination ./mocks/mock_$GOFILE -package mocks
type EpisodeRepository interface {
	AddEpisode(s models.Episode) (models.Episode, error)
	FindById(id int) (models.Episode, error)
	FindAllEpisodesInSeason(serieId int) ([]models.Episode, error)
}

type EpisodeAPI struct {
	Repository EpisodeRepository
}

func (ea EpisodeAPI) AddEpisode(sr dtos.EpisodeRequest) (models.Episode, string, error) {
	if v, ok := sr.Validate(); !ok {
		return models.Episode{}, v, nil
	}

	episode, err := ea.Repository.AddEpisode(sr.GetEpisodeModel())

	if err != nil {
		return models.Episode{}, "", err
	}

	return episode, "", nil
}

func (ea EpisodeAPI) FindById(id int) (models.Episode, error) {
	return ea.Repository.FindById(id)
}

func (ea EpisodeAPI) FindAllEpisodesInSeason(seasonId int) ([]models.Episode, error) {
	return ea.Repository.FindAllEpisodesInSeason(seasonId)
}
