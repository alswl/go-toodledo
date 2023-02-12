// Code generated by mockery v2.19.0. DO NOT EDIT.

package mockui

import mock "github.com/stretchr/testify/mock"

// ContainerizedInterface is an autogenerated mock type for the ContainerizedInterface type
type ContainerizedInterface struct {
	mock.Mock
}

// Children provides a mock function with given fields:
func (_m *ContainerizedInterface) Children() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// FocusChild provides a mock function with given fields: _a0
func (_m *ContainerizedInterface) FocusChild(_a0 string) {
	_m.Called(_a0)
}

// Focused provides a mock function with given fields:
func (_m *ContainerizedInterface) Focused() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *ContainerizedInterface) Next() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewContainerizedInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewContainerizedInterface creates a new instance of ContainerizedInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContainerizedInterface(t mockConstructorTestingTNewContainerizedInterface) *ContainerizedInterface {
	mock := &ContainerizedInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
