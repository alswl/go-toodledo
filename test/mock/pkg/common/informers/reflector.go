// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockinformers

import (
	informers "github.com/alswl/go-toodledo/pkg/common/informers"
	mock "github.com/stretchr/testify/mock"
)

// Reflector is an autogenerated mock type for the Reflector type
type Reflector struct {
	mock.Mock
}

// Chan provides a mock function with given fields:
func (_m *Reflector) Chan() <-chan informers.T {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chan")
	}

	var r0 <-chan informers.T
	if rf, ok := ret.Get(0).(func() <-chan informers.T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan informers.T)
		}
	}

	return r0
}

// HasSynced provides a mock function with given fields:
func (_m *Reflector) HasSynced() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HasSynced")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LastSynced provides a mock function with given fields:
func (_m *Reflector) LastSynced() informers.U {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LastSynced")
	}

	var r0 informers.U
	if rf, ok := ret.Get(0).(func() informers.U); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(informers.U)
		}
	}

	return r0
}

// ListNewer provides a mock function with given fields:
func (_m *Reflector) ListNewer() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListNewer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifyModified provides a mock function with given fields: _a0
func (_m *Reflector) NotifyModified(_a0 interface{}) {
	_m.Called(_a0)
}

// Run provides a mock function with given fields: stop
func (_m *Reflector) Run(stop <-chan struct{}) {
	_m.Called(stop)
}

// NewReflector creates a new instance of Reflector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReflector(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reflector {
	mock := &Reflector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
