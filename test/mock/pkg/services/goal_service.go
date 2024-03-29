// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GoalService is an autogenerated mock type for the GoalService type
type GoalService struct {
	mock.Mock
}

// Archive provides a mock function with given fields: id, isArchived
func (_m *GoalService) Archive(id int, isArchived bool) (*models.Goal, error) {
	ret := _m.Called(id, isArchived)

	if len(ret) == 0 {
		panic("no return value specified for Archive")
	}

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(int, bool) (*models.Goal, error)); ok {
		return rf(id, isArchived)
	}
	if rf, ok := ret.Get(0).(func(int, bool) *models.Goal); ok {
		r0 = rf(id, isArchived)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(int, bool) error); ok {
		r1 = rf(id, isArchived)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: name
func (_m *GoalService) Create(name string) (*models.Goal, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Goal, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Goal); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *GoalService) Delete(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *GoalService) Find(name string) (*models.Goal, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Goal, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Goal); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *GoalService) FindByID(id int64) (*models.Goal, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Goal, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Goal); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAll provides a mock function with given fields:
func (_m *GoalService) ListAll() ([]*models.Goal, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []*models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Goal, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Goal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllWithArchived provides a mock function with given fields:
func (_m *GoalService) ListAllWithArchived() ([]*models.Goal, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAllWithArchived")
	}

	var r0 []*models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Goal, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Goal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rename provides a mock function with given fields: name, newName
func (_m *GoalService) Rename(name string, newName string) (*models.Goal, error) {
	ret := _m.Called(name, newName)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.Goal, error)); ok {
		return rf(name, newName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.Goal); ok {
		r0 = rf(name, newName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, newName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGoalService creates a new instance of GoalService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGoalService(t interface {
	mock.TestingT
	Cleanup(func())
}) *GoalService {
	mock := &GoalService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
