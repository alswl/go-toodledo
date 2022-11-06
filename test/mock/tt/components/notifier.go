// Code generated by mockery v2.14.1. DO NOT EDIT.

package mockcomponents

import mock "github.com/stretchr/testify/mock"

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

// Error provides a mock function with given fields: msg
func (_m *Notifier) Error(msg string) {
	_m.Called(msg)
}

// Info provides a mock function with given fields: msg
func (_m *Notifier) Info(msg string) {
	_m.Called(msg)
}

// Warn provides a mock function with given fields: msg
func (_m *Notifier) Warn(msg string) {
	_m.Called(msg)
}

type mockConstructorTestingTNewNotifier interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotifier creates a new instance of Notifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotifier(t mockConstructorTestingTNewNotifier) *Notifier {
	mock := &Notifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}