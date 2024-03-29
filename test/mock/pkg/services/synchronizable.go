// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockservices

import mock "github.com/stretchr/testify/mock"

// Synchronizable is an autogenerated mock type for the Synchronizable type
type Synchronizable struct {
	mock.Mock
}

// Clean provides a mock function with given fields:
func (_m *Synchronizable) Clean() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clean")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PartialSync provides a mock function with given fields: lastEditTime
func (_m *Synchronizable) PartialSync(lastEditTime *int32) error {
	ret := _m.Called(lastEditTime)

	if len(ret) == 0 {
		panic("no return value specified for PartialSync")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*int32) error); ok {
		r0 = rf(lastEditTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sync provides a mock function with given fields:
func (_m *Synchronizable) Sync() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Sync")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSynchronizable creates a new instance of Synchronizable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynchronizable(t interface {
	mock.TestingT
	Cleanup(func())
}) *Synchronizable {
	mock := &Synchronizable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
