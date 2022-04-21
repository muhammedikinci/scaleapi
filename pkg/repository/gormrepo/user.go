package gormrepo

import (
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db       *gorm.DB
	errorLog *log.Logger
}

func (ur userRepository) Find(username, password string) (models.User, error) {
	var user models.User

	result := ur.db.Where("username = ? AND password = ?", username, password).First(&user)

	if result.Error != nil {
		ur.errorLog.Println(result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}

func (ur userRepository) AddUser(username, password string) (models.User, error) {
	u := models.User{
		Username: username,
		Password: password,
		Role:     models.BasicUser,
	}

	result := ur.db.Create(&u)

	if result.Error != nil {
		ur.errorLog.Println(result.Error)
		return models.User{}, result.Error
	}

	return u, nil
}

func (ur userRepository) FindByUserName(username string) (models.User, error) {
	var user models.User

	result := ur.db.First(&user, "username = ?", username)

	if result.Error != nil {
		ur.errorLog.Println(result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}

func (ur userRepository) AddMovieToFavorite(username string, movie models.Movie) error {
	user, err := ur.FindByUserName(username)

	if err != nil {
		ur.errorLog.Println(err)
		return err
	}

	association := ur.db.Model(&user).Association("MovieFavorites")

	if association.Error != nil {
		ur.errorLog.Println(association.Error)
		return association.Error
	}

	result := association.Append(&movie)

	if result != nil {
		ur.errorLog.Println(result)
		return result
	}

	return nil
}

func (ur userRepository) AddSerieToFavorite(username string, serie models.Serie) error {
	user, err := ur.FindByUserName(username)

	if err != nil {
		ur.errorLog.Println(err)
		return err
	}

	association := ur.db.Model(&user).Association("SerieFavorites")

	if association.Error != nil {
		ur.errorLog.Println(association.Error)
		return association.Error
	}

	result := association.Append(&serie)

	if result != nil {
		ur.errorLog.Println(result)
		return result
	}

	return result
}

func (ur userRepository) GetFavorites(username string) (models.Favorite, error) {
	user, err := ur.FindByUserName(username)

	if err != nil {
		ur.errorLog.Println(err)
		return models.Favorite{}, err
	}

	series := []models.Serie{}
	movies := []models.Movie{}

	result := ur.db.Model(&user).Association("SerieFavorites").Find(&series)

	if result != nil {
		ur.errorLog.Println(result.Error())
		return models.Favorite{}, result
	}

	result = ur.db.Model(&user).Association("MovieFavorites").Find(&movies)

	if result != nil {
		ur.errorLog.Println(result.Error())
		return models.Favorite{}, result
	}

	return models.Favorite{
		Movies: movies,
		Series: series,
	}, nil
}

func (ur userRepository) AddAdmin(username, password string) (models.User, error) {
	u := models.User{
		Username: username,
		Password: password,
		Role:     models.ContentManager,
	}

	result := ur.db.Create(&u)

	if result.Error != nil {
		ur.errorLog.Println(result.Error)
		return models.User{}, result.Error
	}

	return u, nil
}
