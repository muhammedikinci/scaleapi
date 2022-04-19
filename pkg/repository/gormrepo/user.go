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
