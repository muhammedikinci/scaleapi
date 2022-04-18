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
