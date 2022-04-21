package api

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/muhammedikinci/scaleapi/pkg/api/mocks"
	"github.com/muhammedikinci/scaleapi/pkg/dtos"
	"github.com/muhammedikinci/scaleapi/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestWhenCallingAddMovieWithNonValidDTOFunctionReturnErrorMessageWithEmptyModel(t *testing.T) {
	movieApi := MovieAPI{}

	movie, message, err := movieApi.AddMovie(dtos.MovieRequest{})

	assert.Equal(t, message, dtos.ErrEmptyTitleAndDescription)
	assert.Equal(t, movie.Title, "")
	assert.ErrorIs(t, err, nil)
}

func TestWhenMovieRepositoryReturnErrorFunctionReturnNotNilErrorWithEmptyModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockMovieRepository(ctrl)

	mockErr := errors.New("movie cannot added to database")
	mockDto := dtos.MovieRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddMovie(gomock.Any()).Return(models.Movie{}, mockErr)

	movieApi := MovieAPI{Repository: m}

	movie, message, err := movieApi.AddMovie(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, movie.Title, "")
	assert.ErrorIs(t, err, mockErr)
}

func TestWhenMovieRepositoryReturnSuccessfulMovieModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockMovieRepository(ctrl)

	mockDto := dtos.MovieRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddMovie(gomock.Any()).Return(mockDto.GetMovieModel(), nil)

	movieApi := MovieAPI{Repository: m}

	movie, message, err := movieApi.AddMovie(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, movie.Title, "test")
	assert.Equal(t, movie.Description, "Test")
	assert.ErrorIs(t, err, nil)
}
