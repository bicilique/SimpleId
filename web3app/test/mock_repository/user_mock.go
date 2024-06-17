// Code generated by MockGen. DO NOT EDIT.
// Source: repository/user_repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/user_repository.go -destination=mock/user_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

// import (
// 	model "SimpleId/model"
// 	reflect "reflect"

// 	gomock "go.uber.org/mock/gomock"
// )

// // MockUserRepository is a mock of UserRepository interface.
// type MockUserRepository struct {
// 	ctrl     *gomock.Controller
// 	recorder *MockUserRepositoryMockRecorder
// }

// // MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
// type MockUserRepositoryMockRecorder struct {
// 	mock *MockUserRepository
// }

// // NewMockUserRepository creates a new mock instance.
// func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
// 	mock := &MockUserRepository{ctrl: ctrl}
// 	mock.recorder = &MockUserRepositoryMockRecorder{mock}
// 	return mock
// }

// // EXPECT returns an object that allows the caller to indicate expected use.
// func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
// 	return m.recorder
// }

// // Insert mocks base method.
// func (m *MockUserRepository) Insert(entity model.User) (model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Insert", entity)
// 	ret0, _ := ret[0].(model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // Insert indicates an expected call of Insert.
// func (mr *MockUserRepositoryMockRecorder) Insert(entity any) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRepository)(nil).Insert), entity)
// }

// // SelectAll mocks base method.
// func (m *MockUserRepository) SelectAll() ([]model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "SelectAll")
// 	ret0, _ := ret[0].([]model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // SelectAll indicates an expected call of SelectAll.
// func (mr *MockUserRepositoryMockRecorder) SelectAll() *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockUserRepository)(nil).SelectAll))
// }

// // SelectById mocks base method.
// func (m *MockUserRepository) SelectById(id string) (model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "SelectById", id)
// 	ret0, _ := ret[0].(model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // SelectById indicates an expected call of SelectById.
// func (mr *MockUserRepositoryMockRecorder) SelectById(id any) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockUserRepository)(nil).SelectById), id)
// }

// // SelectByName mocks base method.
// func (m *MockUserRepository) SelectByName(field string, value any) (model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "SelectByName", field, value)
// 	ret0, _ := ret[0].(model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // SelectByName indicates an expected call of SelectByName.
// func (mr *MockUserRepositoryMockRecorder) SelectByName(field, value any) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByName", reflect.TypeOf((*MockUserRepository)(nil).SelectByName), field, value)
// }

// // Update mocks base method.
// func (m *MockUserRepository) Update(entity model.User) (model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "Update", entity)
// 	ret0, _ := ret[0].(model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // Update indicates an expected call of Update.
// func (mr *MockUserRepositoryMockRecorder) Update(entity any) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), entity)
// }

// // UpdateById mocks base method.
// func (m *MockUserRepository) UpdateById(id string, updates map[string]any) (model.User, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateById", id, updates)
// 	ret0, _ := ret[0].(model.User)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // UpdateById indicates an expected call of UpdateById.
// func (mr *MockUserRepositoryMockRecorder) UpdateById(id, updates any) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockUserRepository)(nil).UpdateById), id, updates)
// }
