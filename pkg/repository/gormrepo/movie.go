package gormrepo

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (mr movieRepository) RemoveMovie(id int) error {
	movie, err := mr.FindById(id)

	if err != nil {
		return err
	}

	result := mr.db.Select(clause.Associations).Delete(&movie)

	if result.Error != nil {
		mr.errorLog.Println(result.Error)
		return result.Error
	}

	return nil
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

func (mr movieRepository) Filter(title string, genre string) ([]models.Movie, error) {
	var movies []models.Movie

	result := mr.db.Where("LOWER(title) LIKE ? AND LOWER(genre) LIKE ?", wrapLike(title), wrapLike(genre)).Find(&movies)

	if result.Error != nil {
		mr.errorLog.Println(result.Error)
		return nil, result.Error
	}

	return movies, nil
}
