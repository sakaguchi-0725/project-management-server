// Code generated by MockGen. DO NOT EDIT.
// Source: task_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sakaguchi-0725/go-todo/internal/domain"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskRepository) CreateTask(task domain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskRepositoryMockRecorder) CreateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskRepository)(nil).CreateTask), task)
}

// DeleteTask mocks base method.
func (m *MockTaskRepository) DeleteTask(taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskRepositoryMockRecorder) DeleteTask(taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskRepository)(nil).DeleteTask), taskId)
}

// GetAllTasks mocks base method.
func (m *MockTaskRepository) GetAllTasks() ([]domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks")
	ret0, _ := ret[0].([]domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockTaskRepositoryMockRecorder) GetAllTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockTaskRepository)(nil).GetAllTasks))
}

// GetTaskById mocks base method.
func (m *MockTaskRepository) GetTaskById(taksId uint) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskById", taksId)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskById indicates an expected call of GetTaskById.
func (mr *MockTaskRepositoryMockRecorder) GetTaskById(taksId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskById", reflect.TypeOf((*MockTaskRepository)(nil).GetTaskById), taksId)
}

// UpdateTask mocks base method.
func (m *MockTaskRepository) UpdateTask(task domain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskRepositoryMockRecorder) UpdateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTask), task)
}
