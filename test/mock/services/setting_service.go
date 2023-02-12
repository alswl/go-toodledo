// Code generated by mockery v2.19.0. DO NOT EDIT.

package mockservices

import mock "github.com/stretchr/testify/mock"

// SettingService is an autogenerated mock type for the SettingService type
type SettingService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: name
func (_m *SettingService) Delete(name string) error {
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
func (_m *SettingService) Find(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: name, body
func (_m *SettingService) Put(name string, body string) error {
	ret := _m.Called(name, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSettingService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSettingService creates a new instance of SettingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSettingService(t mockConstructorTestingTNewSettingService) *SettingService {
	mock := &SettingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
