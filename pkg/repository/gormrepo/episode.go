package gormrepo

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
)

type episodeRepository struct {
	db       *gorm.DB
	errorLog *log.Logger
}

func (er episodeRepository) AddEpisode(e models.Episode) (models.Episode, error) {
	result := er.db.Create(&e)

	if result.Error != nil {
		er.errorLog.Println(result.Error)
		return models.Episode{}, result.Error
	}

	return e, nil
}

func (er episodeRepository) FindById(id int) (models.Episode, error) {
	var season models.Episode

	result := er.db.First(&season, "id = ?", id)

	if result.Error != nil {
		er.errorLog.Println(result.Error)
		return models.Episode{}, result.Error
	}

	return season, nil
}

func (er episodeRepository) FindAllEpisodesInSeason(seasonId int) ([]models.Episode, error) {
	var seasons []models.Episode

	result := er.db.Find(&seasons, "season_id = ?", seasonId)

	if result.Error != nil {
		er.errorLog.Println(result.Error)
		return []models.Episode{}, result.Error
	}

	return seasons, nil
}
