// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ContextService is an autogenerated mock type for the ContextService type
type ContextService struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (_m *ContextService) Create(name string) (*models.Context, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Context, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Context); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
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
func (_m *ContextService) Delete(name string) error {
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
func (_m *ContextService) Find(name string) (*models.Context, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *models.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Context, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Context); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
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
func (_m *ContextService) FindByID(id int64) (*models.Context, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *models.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Context, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Context); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
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
func (_m *ContextService) ListAll() ([]*models.Context, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []*models.Context
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Context, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Context)
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
func (_m *ContextService) Rename(name string, newName string) (*models.Context, error) {
	ret := _m.Called(name, newName)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 *models.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.Context, error)); ok {
		return rf(name, newName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.Context); ok {
		r0 = rf(name, newName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, newName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContextService creates a new instance of ContextService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContextService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContextService {
	mock := &ContextService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
