package repository

import (
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(data model.User) error {
	ret := m.Called(data)
	return ret.Error(0) 
}

func (m *mockUserRepository) FetchAll() ([]model.User, error) {
	ret := m.Called(0)
	return ret.Get(0).([]model.User), ret.Error(1)
}

func (m *mockUserRepository) FetchByEmail(email string) (model.User, error) {
	ret := m.Called(0)
	return ret.Get(0).(model.User), ret.Error(1)
}