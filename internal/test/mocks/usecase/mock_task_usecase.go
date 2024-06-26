// Code generated by MockGen. DO NOT EDIT.
// Source: task_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sakaguchi-0725/go-todo/internal/domain"
	input "github.com/sakaguchi-0725/go-todo/internal/usecase/input"
)

// MockTaskUsecase is a mock of TaskUsecase interface.
type MockTaskUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUsecaseMockRecorder
}

// MockTaskUsecaseMockRecorder is the mock recorder for MockTaskUsecase.
type MockTaskUsecaseMockRecorder struct {
	mock *MockTaskUsecase
}

// NewMockTaskUsecase creates a new mock instance.
func NewMockTaskUsecase(ctrl *gomock.Controller) *MockTaskUsecase {
	mock := &MockTaskUsecase{ctrl: ctrl}
	mock.recorder = &MockTaskUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUsecase) EXPECT() *MockTaskUsecaseMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskUsecase) CreateTask(input input.TaskInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskUsecaseMockRecorder) CreateTask(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskUsecase)(nil).CreateTask), input)
}

// DeleteTask mocks base method.
func (m *MockTaskUsecase) DeleteTask(taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskUsecaseMockRecorder) DeleteTask(taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskUsecase)(nil).DeleteTask), taskId)
}

// GetAllTasks mocks base method.
func (m *MockTaskUsecase) GetAllTasks() ([]domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks")
	ret0, _ := ret[0].([]domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockTaskUsecaseMockRecorder) GetAllTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockTaskUsecase)(nil).GetAllTasks))
}

// GetTaskById mocks base method.
func (m *MockTaskUsecase) GetTaskById(taskId uint) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskById", taskId)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskById indicates an expected call of GetTaskById.
func (mr *MockTaskUsecaseMockRecorder) GetTaskById(taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskById", reflect.TypeOf((*MockTaskUsecase)(nil).GetTaskById), taskId)
}

// UpdateTask mocks base method.
func (m *MockTaskUsecase) UpdateTask(input input.TaskInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskUsecaseMockRecorder) UpdateTask(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskUsecase)(nil).UpdateTask), input)
}
