// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockui

import mock "github.com/stretchr/testify/mock"

// FocusableInterface is an autogenerated mock type for the FocusableInterface type
type FocusableInterface struct {
	mock.Mock
}

// Blur provides a mock function with given fields:
func (_m *FocusableInterface) Blur() {
	_m.Called()
}

// Focus provides a mock function with given fields:
func (_m *FocusableInterface) Focus() {
	_m.Called()
}

// NewFocusableInterface creates a new instance of FocusableInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFocusableInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FocusableInterface {
	mock := &FocusableInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
