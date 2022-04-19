package gormrepo

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
)

type seasonRepository struct {
	db       *gorm.DB
	errorLog *log.Logger
}

func (sr seasonRepository) AddSeason(s models.Season) (models.Season, error) {
	result := sr.db.Create(&s)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return models.Season{}, result.Error
	}

	return s, nil
}

func (sr seasonRepository) FindById(id int) (models.Season, error) {
	var season models.Season

	result := sr.db.Preload("Episodes").First(&season, "id = ?", id)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return models.Season{}, result.Error
	}

	return season, nil
}

func (sr seasonRepository) FindAllSeasonsInSerie(serieId int) ([]models.Season, error) {
	var seasons []models.Season

	result := sr.db.Preload("Episodes").Find(&seasons, "serie_id = ?", serieId)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return []models.Season{}, result.Error
	}

	return seasons, nil
}
