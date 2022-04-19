package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
)

type userApi interface {
	Login(user dtos.LoginRegisterRequest) (dtos.LoginResponse, error)
	Register(user dtos.LoginRegisterRequest) (dtos.RegisterResponse, error)
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
