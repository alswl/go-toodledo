// Code generated by mockery v2.14.0. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ContextLocalService is an autogenerated mock type for the ContextLocalService type
type ContextLocalService struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (_m *ContextLocalService) Create(name string) (*models.Context, error) {
	ret := _m.Called(name)

	var r0 *models.Context
	if rf, ok := ret.Get(0).(func(string) *models.Context); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *ContextLocalService) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *ContextLocalService) Find(name string) (*models.Context, error) {
	ret := _m.Called(name)

	var r0 *models.Context
	if rf, ok := ret.Get(0).(func(string) *models.Context); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *ContextLocalService) FindByID(id int64) (*models.Context, error) {
	ret := _m.Called(id)

	var r0 *models.Context
	if rf, ok := ret.Get(0).(func(int64) *models.Context); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
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

// ListAll provides a mock function with given fields:
func (_m *ContextLocalService) ListAll() ([]*models.Context, error) {
	ret := _m.Called()

	var r0 []*models.Context
	if rf, ok := ret.Get(0).(func() []*models.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalClear provides a mock function with given fields:
func (_m *ContextLocalService) LocalClear() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PartialSync provides a mock function with given fields: lastEditTime
func (_m *ContextLocalService) PartialSync(lastEditTime *int32) error {
	ret := _m.Called(lastEditTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(*int32) error); ok {
		r0 = rf(lastEditTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rename provides a mock function with given fields: name, newName
func (_m *ContextLocalService) Rename(name string, newName string) (*models.Context, error) {
	ret := _m.Called(name, newName)

	var r0 *models.Context
	if rf, ok := ret.Get(0).(func(string, string) *models.Context); ok {
		r0 = rf(name, newName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, newName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sync provides a mock function with given fields:
func (_m *ContextLocalService) Sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewContextLocalService interface {
	mock.TestingT
	Cleanup(func())
}

// NewContextLocalService creates a new instance of ContextLocalService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContextLocalService(t mockConstructorTestingTNewContextLocalService) *ContextLocalService {
	mock := &ContextLocalService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}