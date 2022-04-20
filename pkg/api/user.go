package api

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type userRepository interface {
	FindByUserName(username string) (models.User, error)
	AddUser(username, password string) (models.User, error)
	AddMovieToFavorite(username string, movie models.Movie) error
	AddSerieToFavorite(username string, serie models.Serie) error
	GetFavorites(username string) (models.Favorite, error)
}

type UserAPI struct {
	Repository      userRepository
	MovieRepository movieRepository
	SerieRepository serieRepository
}

var hmacSampleSecret []byte = []byte("very-secret")

func (ua UserAPI) Login(user dtos.LoginRegisterRequest) (dtos.LoginResponse, error) {
	if v, ok := user.Validate(); !ok {
		return dtos.LoginResponse{
			Status:  false,
			Message: v,
		}, nil
	}

	result, err := ua.Repository.FindByUserName(user.Username)

	if err != nil {
		return dtos.LoginResponse{
			Status:  false,
			Message: "User not found",
		}, nil
	}

	if !user.CheckPasswordHash(result.Password) {
		return dtos.LoginResponse{
			Status:  false,
			Message: "Credentials does not match",
		}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return dtos.LoginResponse{}, err
	}

	return dtos.LoginResponse{
		Status:   true,
		Message:  "Login successful",
		Username: result.Username,
		Token:    tokenString,
	}, nil
}

func (ua UserAPI) Register(user dtos.LoginRegisterRequest) (dtos.RegisterResponse, error) {
	if v, ok := user.Validate(); !ok {
		return dtos.RegisterResponse{
			Status:  false,
			Message: v,
		}, nil
	}

	user.HashPassword()

	_, err := ua.Repository.AddUser(user.Username, user.Password)

	if err != nil {
		return dtos.RegisterResponse{}, err
	}

	return dtos.RegisterResponse{
		Status:  true,
		Message: "Registration successful",
	}, nil
}

func (ua UserAPI) CheckTokenAndGetUser(tokenString string) (models.User, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signin")
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return models.User{}, false
	}

	var username string
	result := false

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = claims["username"].(string)
	} else {
		return models.User{}, result
	}

	user, err := ua.Repository.FindByUserName(username)

	if err == nil && user.Username != "" {
		result = true
	}

	return user, result
}

func (ua UserAPI) AddMovieToFavorite(username string, movieId int) bool {
	movie, err := ua.MovieRepository.FindById(movieId)

	if err != nil {
		return false
	}

	err = ua.Repository.AddMovieToFavorite(username, movie)

	return err == nil
}

func (ua UserAPI) AddSerieToFavorite(username string, serieId int) bool {
	serie, err := ua.SerieRepository.FindById(serieId)

	if err != nil {
		return false
	}

	err = ua.Repository.AddSerieToFavorite(username, serie)

	return err == nil
}

func (ua UserAPI) GetFavorites(username string) (models.Favorite, error) {
	return ua.Repository.GetFavorites(username)
}
