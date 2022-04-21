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

func TestWhenCallingAddEpisodeWithNonValidDTOFunctionReturnErrorMessageWithEmptyModel(t *testing.T) {
	episodeApi := EpisodeAPI{}

	episode, message, err := episodeApi.AddEpisode(dtos.EpisodeRequest{})

	assert.Equal(t, message, dtos.ErrEmptyTitleAndDescription)
	assert.Equal(t, episode.Title, "")
	assert.ErrorIs(t, err, nil)
}

func TestWhenEpisodeRepositoryReturnErrorFunctionReturnNotNilErrorWithEmptyModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockEpisodeRepository(ctrl)

	mockErr := errors.New("episode cannot added to database")
	mockDto := dtos.EpisodeRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddEpisode(mockDto.GetEpisodeModel()).Return(models.Episode{}, mockErr)

	episodeApi := EpisodeAPI{Repository: m}

	episode, message, err := episodeApi.AddEpisode(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, episode.Title, "")
	assert.ErrorIs(t, err, mockErr)
}

func TestWhenEpisodeRepositoryReturnEpisodeModel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockEpisodeRepository(ctrl)

	mockDto := dtos.EpisodeRequest{Title: "test", Description: "Test"}

	m.EXPECT().AddEpisode(mockDto.GetEpisodeModel()).Return(mockDto.GetEpisodeModel(), nil)

	episodeApi := EpisodeAPI{Repository: m}

	episode, message, err := episodeApi.AddEpisode(mockDto)

	assert.Equal(t, message, "")
	assert.Equal(t, episode.Title, "test")
	assert.ErrorIs(t, err, nil)
}
