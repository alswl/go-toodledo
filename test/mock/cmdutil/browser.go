// Code generated by mockery v2.14.1. DO NOT EDIT.

package mockcmdutil

import mock "github.com/stretchr/testify/mock"

// Browser is an autogenerated mock type for the Browser type
type Browser struct {
	mock.Mock
}

// Browse provides a mock function with given fields: _a0
func (_m *Browser) Browse(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBrowser interface {
	mock.TestingT
	Cleanup(func())
}

// NewBrowser creates a new instance of Browser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBrowser(t mockConstructorTestingTNewBrowser) *Browser {
	mock := &Browser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
