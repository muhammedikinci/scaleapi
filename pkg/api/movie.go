package api

import "github.com/muhammedikinci/scaleapi/pkg/models"

func (a api) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie

	result := a.db.Find(&movies)

	if result.Error != nil {
		return nil, result.Error
	}

	return movies, nil
}

func (a api) AddMovie(m models.Movie) (models.Movie, error) {
	result := a.db.Create(&m)

	if result.Error != nil {
		return models.Movie{}, result.Error
	}

	return m, nil
}

func (a api) FindById(id int) (models.Movie, error) {
	var movie models.Movie

	result := a.db.First(&movie, "id = ?", id)

	if result.Error != nil {
		return models.Movie{}, result.Error
	}

	return movie, nil
}
