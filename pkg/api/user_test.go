package api

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammedikinci/scaleapi/pkg/api/mocks"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestWhenCallingLoginWithNonValidDTOFunctionReturnLoginResponseModel(t *testing.T) {
	userApi := UserAPI{}

	response, err := userApi.Login(dtos.LoginRegisterRequest{})

	assert.Equal(t, response.Status, false)
	assert.Equal(t, response.Message, dtos.ErrUsernameAndPasswordLengthError)
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingRegisterWithNonValidDTOFunctionReturnLoginResponseModel(t *testing.T) {
	userApi := UserAPI{}

	response, err := userApi.Register(dtos.LoginRegisterRequest{})

	assert.Equal(t, response.Status, false)
	assert.Equal(t, response.Message, dtos.ErrUsernameAndPasswordLengthError)
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingLoginWithNonUserDataDTOReturnUserNotFoundLoginResponseModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockUserRepository(ctrl)

	userRepository.EXPECT().FindByUserName(gomock.Any()).Return(models.User{}, errors.New("user not found"))

	userApi := UserAPI{Repository: userRepository}

	response, err := userApi.Login(dtos.LoginRegisterRequest{Username: "test", Password: "test"})

	assert.Equal(t, response.Status, false)
	assert.Equal(t, response.Message, ErrUserNotFound)
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingLoginWithWrongUserDataDTOReturnCredentialsDoesNotMatchLoginResponseModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockUserRepository(ctrl)

	userRepository.EXPECT().FindByUserName(gomock.Any()).Return(models.User{}, nil)

	userApi := UserAPI{Repository: userRepository}

	response, err := userApi.Login(dtos.LoginRegisterRequest{Username: "test", Password: "test"})

	assert.Equal(t, response.Status, false)
	assert.Equal(t, response.Message, ErrCredentialDoesNotMatch)
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingLoginWithTrueUserDataDTOReturnSuccessfulLoginResponseModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockUserRepository(ctrl)

	request := dtos.LoginRegisterRequest{Username: "test", Password: "test"}
	request.HashPassword()

	userRepository.
		EXPECT().
		FindByUserName(gomock.Any()).
		Return(models.User{Username: "test", Password: request.Password}, nil)

	userApi := UserAPI{Repository: userRepository}

	response, err := userApi.Login(dtos.LoginRegisterRequest{Username: "test", Password: "test"})

	assert.Equal(t, response.Status, true)
	assert.Equal(t, response.Message, "Login successful")
	assert.True(t, response.Token != "")
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingRegisterWithAlreadyTakenUsernameDataDTOReturnUsernameAlreadyTakenLoginResponseModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockUserRepository(ctrl)

	userRepository.EXPECT().FindByUserName(gomock.Any()).Return(models.User{Username: "test"}, nil)

	userApi := UserAPI{Repository: userRepository}

	response, err := userApi.Register(dtos.LoginRegisterRequest{Username: "test", Password: "test"})

	assert.Equal(t, response.Status, false)
	assert.Equal(t, response.Message, ErrUsernameAlreadyTaken)
	assert.ErrorIs(t, err, nil)
}

func TestWhenCallingRegisterReturnRegistrationSuccessfulResponseModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockUserRepository(ctrl)

	userRepository.EXPECT().FindByUserName(gomock.Any()).Return(models.User{}, nil)
	userRepository.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(models.User{}, nil)

	userApi := UserAPI{Repository: userRepository}

	response, err := userApi.Register(dtos.LoginRegisterRequest{Username: "test", Password: "test"})

	assert.Equal(t, response.Status, true)
	assert.Equal(t, response.Message, "Registration successful")
	assert.ErrorIs(t, err, nil)
}
