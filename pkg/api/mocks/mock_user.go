// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/muhammedikinci/scaleapi/pkg/models"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddMovieToFavorite mocks base method.
func (m *MockUserRepository) AddMovieToFavorite(username string, movie models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMovieToFavorite", username, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMovieToFavorite indicates an expected call of AddMovieToFavorite.
func (mr *MockUserRepositoryMockRecorder) AddMovieToFavorite(username, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMovieToFavorite", reflect.TypeOf((*MockUserRepository)(nil).AddMovieToFavorite), username, movie)
}

// AddSerieToFavorite mocks base method.
func (m *MockUserRepository) AddSerieToFavorite(username string, serie models.Serie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSerieToFavorite", username, serie)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSerieToFavorite indicates an expected call of AddSerieToFavorite.
func (mr *MockUserRepositoryMockRecorder) AddSerieToFavorite(username, serie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSerieToFavorite", reflect.TypeOf((*MockUserRepository)(nil).AddSerieToFavorite), username, serie)
}

// AddUser mocks base method.
func (m *MockUserRepository) AddUser(username, password string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", username, password)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepositoryMockRecorder) AddUser(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepository)(nil).AddUser), username, password)
}

// FindByUserName mocks base method.
func (m *MockUserRepository) FindByUserName(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserName", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserName indicates an expected call of FindByUserName.
func (mr *MockUserRepositoryMockRecorder) FindByUserName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserName", reflect.TypeOf((*MockUserRepository)(nil).FindByUserName), username)
}

// GetFavorites mocks base method.
func (m *MockUserRepository) GetFavorites(username string) (models.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", username)
	ret0, _ := ret[0].(models.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockUserRepositoryMockRecorder) GetFavorites(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockUserRepository)(nil).GetFavorites), username)
}
