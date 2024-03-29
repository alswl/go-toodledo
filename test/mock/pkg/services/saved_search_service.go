// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// SavedSearchService is an autogenerated mock type for the SavedSearchService type
type SavedSearchService struct {
	mock.Mock
}

// ListAll provides a mock function with given fields:
func (_m *SavedSearchService) ListAll() ([]*models.SavedSearch, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []*models.SavedSearch
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.SavedSearch, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.SavedSearch); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.SavedSearch)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSavedSearchService creates a new instance of SavedSearchService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSavedSearchService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SavedSearchService {
	mock := &SavedSearchService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
