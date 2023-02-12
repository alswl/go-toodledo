// Code generated by mockery v2.19.0. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// AccountExtService is an autogenerated mock type for the AccountExtService type
type AccountExtService struct {
	mock.Mock
}

// CachedMe provides a mock function with given fields:
func (_m *AccountExtService) CachedMe() (*models.Account, bool, error) {
	ret := _m.Called()

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func() *models.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindLastFetchInfo provides a mock function with given fields:
func (_m *AccountExtService) FindLastFetchInfo() (*models.Account, error) {
	ret := _m.Called()

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func() *models.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
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

// Me provides a mock function with given fields:
func (_m *AccountExtService) Me() (*models.Account, error) {
	ret := _m.Called()

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func() *models.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
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

// ModifyLastFetchInfo provides a mock function with given fields: account
func (_m *AccountExtService) ModifyLastFetchInfo(account *models.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAccountExtService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountExtService creates a new instance of AccountExtService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountExtService(t mockConstructorTestingTNewAccountExtService) *AccountExtService {
	mock := &AccountExtService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}