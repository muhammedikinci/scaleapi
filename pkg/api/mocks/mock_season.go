// Code generated by MockGen. DO NOT EDIT.
// Source: season.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/muhammedikinci/scaleapi/pkg/models"
)

// MockSeasonRepository is a mock of SeasonRepository interface.
type MockSeasonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSeasonRepositoryMockRecorder
}

// MockSeasonRepositoryMockRecorder is the mock recorder for MockSeasonRepository.
type MockSeasonRepositoryMockRecorder struct {
	mock *MockSeasonRepository
}

// NewMockSeasonRepository creates a new mock instance.
func NewMockSeasonRepository(ctrl *gomock.Controller) *MockSeasonRepository {
	mock := &MockSeasonRepository{ctrl: ctrl}
	mock.recorder = &MockSeasonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeasonRepository) EXPECT() *MockSeasonRepositoryMockRecorder {
	return m.recorder
}

// AddSeason mocks base method.
func (m *MockSeasonRepository) AddSeason(s models.Season) (models.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSeason", s)
	ret0, _ := ret[0].(models.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSeason indicates an expected call of AddSeason.
func (mr *MockSeasonRepositoryMockRecorder) AddSeason(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSeason", reflect.TypeOf((*MockSeasonRepository)(nil).AddSeason), s)
}

// FindAllSeasonsInSerie mocks base method.
func (m *MockSeasonRepository) FindAllSeasonsInSerie(serieId int) ([]models.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllSeasonsInSerie", serieId)
	ret0, _ := ret[0].([]models.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllSeasonsInSerie indicates an expected call of FindAllSeasonsInSerie.
func (mr *MockSeasonRepositoryMockRecorder) FindAllSeasonsInSerie(serieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllSeasonsInSerie", reflect.TypeOf((*MockSeasonRepository)(nil).FindAllSeasonsInSerie), serieId)
}

// FindById mocks base method.
func (m *MockSeasonRepository) FindById(id int) (models.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(models.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockSeasonRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockSeasonRepository)(nil).FindById), id)
}