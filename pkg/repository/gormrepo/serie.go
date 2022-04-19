package gormrepo

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
)

type serieRepository struct {
	db       *gorm.DB
	errorLog *log.Logger
}

func (sr serieRepository) GetAllSeries() ([]models.Serie, error) {
	var series []models.Serie

	result := sr.db.Preload("Seasons").Find(&series)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return nil, result.Error
	}

	return series, nil
}

func (sr serieRepository) AddSerie(s models.Serie) (models.Serie, error) {
	result := sr.db.Create(&s)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return models.Serie{}, result.Error
	}

	return s, nil
}

func (sr serieRepository) FindById(id int) (models.Serie, error) {
	var serie models.Serie

	result := sr.db.Preload("Seasons").First(&serie, "id = ?", id)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return models.Serie{}, result.Error
	}

	return serie, nil
}

func (sr serieRepository) Filter(title string, genre string) ([]models.Serie, error) {
	var series []models.Serie

	result := sr.db.Where("LOWER(title) LIKE ? AND LOWER(genre) LIKE ?", wrapLike(title), wrapLike(genre)).Preload("Seasons").Find(&series)

	if result.Error != nil {
		sr.errorLog.Println(result.Error)
		return nil, result.Error
	}

	return series, nil
}

func wrapLike(field string) string {
	return "%" + field + "%"
}
