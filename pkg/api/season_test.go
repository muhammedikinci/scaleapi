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

func TestWhenCallingAddSeasonWithNonValidDTOFunctionReturnErrorMessageWithEmptyModel(t *testing.T) {
	seasonApi := SeasonAPI{}

	season, message, err := seasonApi.AddSeason(dtos.SeasonRequest{})

	assert.Equal(t, message, dtos.ErrEmptyTitle)
	assert.Equal(t, season.Title, "")
	assert.ErrorIs(t, err, nil)
}

func TestWhenSeasonRepositoryReturnErrorFunctionReturnNotNilErrorWithEmptyModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockSeasonRepository(ctrl)

	mockErr := errors.New("season cannot added to database")
	mockDto := dtos.SeasonRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddSeason(gomock.Any()).Return(models.Season{}, mockErr)

	seasonApi := SeasonAPI{Repository: m}

	season, message, err := seasonApi.AddSeason(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, season.Title, "")
	assert.ErrorIs(t, err, mockErr)
}

func TestWhenSeasonRepositoryReturnSuccessfulSeasonModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockSeasonRepository(ctrl)

	mockDto := dtos.SeasonRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddSeason(gomock.Any()).Return(mockDto.GetSeasonModel(), nil)

	seasonApi := SeasonAPI{Repository: m}

	season, message, err := seasonApi.AddSeason(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, season.Title, "test")
	assert.Equal(t, season.Description, "Test")
	assert.ErrorIs(t, err, nil)
}
