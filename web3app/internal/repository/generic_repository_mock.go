package repository

import (
	"github.com/stretchr/testify/mock"
)

type GenericRepositoryMock[T any] struct {
	Mock mock.Mock
}

func (m *GenericRepositoryMock[T]) Insert(entity T) (T, error) {
	// args := m.Mock.Called()
	args := m.Mock.Called(entity)
	return args.Get(0).(T), args.Error(1)
}

func (m *GenericRepositoryMock[T]) SelectAll() ([]T, error) {
	args := m.Mock.Called()
	return args.Get(0).([]T), args.Error(1)
}

func (m *GenericRepositoryMock[T]) SelectById(id string) (T, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(T), args.Error(1)
}

func (m *GenericRepositoryMock[T]) SelectByName(field string, value interface{}) (T, error) {
	args := m.Mock.Called(field, value)
	return args.Get(0).(T), args.Error(1)
}

func (m *GenericRepositoryMock[T]) Update(id int, updates map[string]interface{}) (T, error) {
	args := m.Mock.Called(id, updates)
	return args.Get(0).(T), args.Error(1)
}
