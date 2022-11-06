// Code generated by mockery v2.14.1. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskRichService is an autogenerated mock type for the TaskRichService type
type TaskRichService struct {
	mock.Mock
}

// Find provides a mock function with given fields: id
func (_m *TaskRichService) Find(id int64) (*models.RichTask, error) {
	ret := _m.Called(id)

	var r0 *models.RichTask
	if rf, ok := ret.Get(0).(func(int64) *models.RichTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RichTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rich provides a mock function with given fields: tasks
func (_m *TaskRichService) Rich(tasks *models.Task) (*models.RichTask, error) {
	ret := _m.Called(tasks)

	var r0 *models.RichTask
	if rf, ok := ret.Get(0).(func(*models.Task) *models.RichTask); ok {
		r0 = rf(tasks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RichTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Task) error); ok {
		r1 = rf(tasks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RichThem provides a mock function with given fields: tasks
func (_m *TaskRichService) RichThem(tasks []*models.Task) ([]*models.RichTask, error) {
	ret := _m.Called(tasks)

	var r0 []*models.RichTask
	if rf, ok := ret.Get(0).(func([]*models.Task) []*models.RichTask); ok {
		r0 = rf(tasks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.RichTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*models.Task) error); ok {
		r1 = rf(tasks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskRichService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskRichService creates a new instance of TaskRichService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskRichService(t mockConstructorTestingTNewTaskRichService) *TaskRichService {
	mock := &TaskRichService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
