package api

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
)

type movieRepository struct {
	db       *gorm.DB
	errorLog *log.Logger
}

func (mr movieRepository) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie

	result := mr.db.Find(&movies)

	if result.Error != nil {
		mr.errorLog.Println(result.Error)
		return nil, result.Error
	}

	return movies, nil
}

func (mr movieRepository) AddMovie(m models.Movie) (models.Movie, error) {
	result := mr.db.Create(&m)

	if result.Error != nil {
		mr.errorLog.Println(result.Error)
		return models.Movie{}, result.Error
	}

	return m, nil
}

func (mr movieRepository) FindById(id int) (models.Movie, error) {
	var movie models.Movie

	result := mr.db.First(&movie, "id = ?", id)

	if result.Error != nil {
		mr.errorLog.Println(result.Error)
		return models.Movie{}, result.Error
	}

	return movie, nil
}
