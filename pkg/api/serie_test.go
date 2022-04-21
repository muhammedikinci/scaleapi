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

func TestWhenCallingAddSerieWithNonValidDTOFunctionReturnErrorMessageWithEmptyModel(t *testing.T) {
	serieApi := SerieAPI{}

	serie, message, err := serieApi.AddSerie(dtos.SerieRequest{})

	assert.Equal(t, message, dtos.ErrEmptyTitleAndDescription)
	assert.Equal(t, serie.Title, "")
	assert.ErrorIs(t, err, nil)
}

func TestWhenSerieRepositoryReturnErrorFunctionReturnNotNilErrorWithEmptyModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockSerieRepository(ctrl)

	mockErr := errors.New("serie cannot added to database")
	mockDto := dtos.SerieRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddSerie(gomock.Any()).Return(models.Serie{}, mockErr)

	serieApi := SerieAPI{Repository: m}

	serie, message, err := serieApi.AddSerie(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, serie.Title, "")
	assert.ErrorIs(t, err, mockErr)
}

func TestWhenSerieRepositoryReturnSuccessfulSerieModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockSerieRepository(ctrl)

	mockDto := dtos.SerieRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddSerie(gomock.Any()).Return(mockDto.GetSerieModel(), nil)

	serieApi := SerieAPI{Repository: m}

	serie, message, err := serieApi.AddSerie(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, serie.Title, "test")
	assert.Equal(t, serie.Description, "Test")
	assert.ErrorIs(t, err, nil)
}
