// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/generic_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/generic_repository.go -destination=generic_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder[T]
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder[T any] struct {
	mock *MockRepository[T]
}

// NewMockRepository creates a new mock instance.
func NewMockRepository[T any](ctrl *gomock.Controller) *MockRepository[T] {
	mock := &MockRepository[T]{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository[T]) EXPECT() *MockRepositoryMockRecorder[T] {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRepository[T]) Insert(entity T) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", entity)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder[T]) Insert(entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository[T])(nil).Insert), entity)
}

// SelectAll mocks base method.
func (m *MockRepository[T]) SelectAll() ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockRepositoryMockRecorder[T]) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockRepository[T])(nil).SelectAll))
}

// SelectByField mocks base method.
func (m *MockRepository[T]) SelectByField(field string, value any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByField", field, value)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByField indicates an expected call of SelectByField.
func (mr *MockRepositoryMockRecorder[T]) SelectByField(field, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByField", reflect.TypeOf((*MockRepository[T])(nil).SelectByField), field, value)
}

// SelectById mocks base method.
func (m *MockRepository[T]) SelectById(id string) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockRepositoryMockRecorder[T]) SelectById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockRepository[T])(nil).SelectById), id)
}

// SelectByRequestId mocks base method.
func (m *MockRepository[T]) SelectByRequestId(id string) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByRequestId", id)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByRequestId indicates an expected call of SelectByRequestId.
func (mr *MockRepositoryMockRecorder[T]) SelectByRequestId(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByRequestId", reflect.TypeOf((*MockRepository[T])(nil).SelectByRequestId), id)
}

// SelectByUID mocks base method.
func (m *MockRepository[T]) SelectByUID(uid string) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUID", uid)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUID indicates an expected call of SelectByUID.
func (mr *MockRepositoryMockRecorder[T]) SelectByUID(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUID", reflect.TypeOf((*MockRepository[T])(nil).SelectByUID), uid)
}

// ShowByField mocks base method.
func (m *MockRepository[T]) ShowByField(field string, value any) ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowByField", field, value)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowByField indicates an expected call of ShowByField.
func (mr *MockRepositoryMockRecorder[T]) ShowByField(field, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowByField", reflect.TypeOf((*MockRepository[T])(nil).ShowByField), field, value)
}

// Update mocks base method.
func (m *MockRepository[T]) Update(entity T) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", entity)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder[T]) Update(entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository[T])(nil).Update), entity)
}

// UpdateById mocks base method.
func (m *MockRepository[T]) UpdateById(id string, updates map[string]any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", id, updates)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockRepositoryMockRecorder[T]) UpdateById(id, updates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockRepository[T])(nil).UpdateById), id, updates)
}

// UpdateByRequestId mocks base method.
func (m *MockRepository[T]) UpdateByRequestId(id string, updates map[string]any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByRequestId", id, updates)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByRequestId indicates an expected call of UpdateByRequestId.
func (mr *MockRepositoryMockRecorder[T]) UpdateByRequestId(id, updates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByRequestId", reflect.TypeOf((*MockRepository[T])(nil).UpdateByRequestId), id, updates)
}

// UpdateByUID mocks base method.
func (m *MockRepository[T]) UpdateByUID(id string, updates map[string]any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByUID", id, updates)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByUID indicates an expected call of UpdateByUID.
func (mr *MockRepositoryMockRecorder[T]) UpdateByUID(id, updates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByUID", reflect.TypeOf((*MockRepository[T])(nil).UpdateByUID), id, updates)
}
