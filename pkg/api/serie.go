package api

import (
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type serieRepository interface {
	GetAllSeries() ([]models.Serie, error)
	AddSerie(models.Serie) (models.Serie, error)
	FindById(id int) (models.Serie, error)
	Filter(title string, genre string) ([]models.Serie, error)
}

type SerieAPI struct {
	Repository serieRepository
}

func (sa SerieAPI) GetAllSeries() ([]models.Serie, error) {
	return sa.Repository.GetAllSeries()
}

func (sa SerieAPI) AddSerie(s dtos.SerieRequest) (models.Serie, string, error) {
	if v, ok := s.Validate(); !ok {
		return models.Serie{}, v, nil
	}

	serie, err := sa.Repository.AddSerie(s.GetSerieModel())

	if err != nil {
		return models.Serie{}, "", err
	}

	return serie, "", nil
}

func (sa SerieAPI) FindById(id int) (models.Serie, error) {
	return sa.Repository.FindById(id)
}

func (sa SerieAPI) Filter(title, genre string) ([]models.Serie, error) {
	return sa.Repository.Filter(title, genre)
}
