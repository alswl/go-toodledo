// Code generated by mockery v2.14.0. DO NOT EDIT.

package mockfetcher

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

// Start provides a mock function with given fields: _a0
func (_m *Fetcher) Start(_a0 context.Context) {
	_m.Called(_a0)
}

// Stop provides a mock function with given fields:
func (_m *Fetcher) Stop() {
	_m.Called()
}

// fetch provides a mock function with given fields:
func (_m *Fetcher) fetch() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFetcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewFetcher creates a new instance of Fetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFetcher(t mockConstructorTestingTNewFetcher) *Fetcher {
	mock := &Fetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
