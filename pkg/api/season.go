package api

import (
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type seasonRepository interface {
	AddSeason(s models.Season) (models.Season, error)
	FindById(id int) (models.Season, error)
	FindAllSeasonsInSerie(serieId int) ([]models.Season, error)
}

type SeasonAPI struct {
	Repository seasonRepository
}

func (sa SeasonAPI) AddSeason(sr dtos.SeasonRequest) (models.Season, string, error) {
	if v, ok := sr.Validate(); !ok {
		return models.Season{}, v, nil
	}

	season, err := sa.Repository.AddSeason(sr.GetSeasonModel())

	if err != nil {
		return models.Season{}, "", err
	}

	return season, "", nil
}

func (sa SeasonAPI) FindById(id int) (models.Season, error) {
	return sa.Repository.FindById(id)
}

func (sa SeasonAPI) FindAllSeasonsInSerie(serieId int) ([]models.Season, error) {
	return sa.Repository.FindAllSeasonsInSerie(serieId)
}
