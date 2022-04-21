package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
)

type userApi interface {
	Login(user dtos.LoginRegisterRequest) (dtos.LoginResponse, error)
	Register(user dtos.LoginRegisterRequest) (dtos.RegisterResponse, error)
	AddMovieToFavorite(username string, movieId int) bool
	AddSerieToFavorite(username string, serieId int) bool
	GetFavorites(username string) (models.Favorite, error)
	GetFilteredFavorites(title, genre, username string) (models.Favorite, error)
}

type userHandler struct {
	api userApi
}

func NewUserHandler(api userApi) *userHandler {
	return &userHandler{
		api: api,
	}
}

func (uh userHandler) Login(c echo.Context) error {
	request := new(dtos.LoginRegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := uh.api.Login(*request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uh userHandler) Register(c echo.Context) error {
	request := new(dtos.LoginRegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := uh.api.Register(*request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uh userHandler) AddMovieToFavorite(c echo.Context) error {
	username := c.Get("username").(string)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "id parameter is not valid",
		})
	}

	if username == "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "user not found",
		})
	}

	if uh.api.AddMovieToFavorite(username, id) {
		return c.JSON(http.StatusOK, dtos.ResponseMessage{
			Status:  true,
			Message: "movie added to favorite",
		})
	}

	return c.JSON(http.StatusInternalServerError, dtos.ResponseMessage{
		Status:  false,
		Message: "movie cannot add to favorite",
	})
}

func (uh userHandler) AddSerieToFavorite(c echo.Context) error {
	username := c.Get("username").(string)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "id parameter is not valid",
		})
	}

	if username == "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "user not found",
		})
	}

	if uh.api.AddSerieToFavorite(username, id) {
		return c.JSON(http.StatusOK, dtos.ResponseMessage{
			Status:  true,
			Message: "serie added to favorite",
		})
	}

	return c.JSON(http.StatusInternalServerError, dtos.ResponseMessage{
		Status:  false,
		Message: "serie cannot add to favorite",
	})
}

func (uh userHandler) GetFavorites(c echo.Context) error {
	username := c.Get("username").(string)

	if username == "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "user not found",
		})
	}

	favorites, err := uh.api.GetFavorites(username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ResponseMessage{
			Status:  false,
			Message: "cannot getting customer favorites",
		})
	}

	return c.JSON(http.StatusOK, favorites)
}

func (uh userHandler) GetFilteredFavorites(c echo.Context) error {
	username := c.Get("username").(string)

	if username == "" {
		return c.JSON(http.StatusBadRequest, dtos.ResponseMessage{
			Status:  false,
			Message: "user not found",
		})
	}

	favorites, err := uh.api.GetFilteredFavorites(c.QueryParam("title"), c.QueryParam("genre"), username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ResponseMessage{
			Status:  false,
			Message: "cannot getting customer favorites",
		})
	}

	return c.JSON(http.StatusOK, favorites)
}
