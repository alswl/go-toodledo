// Code generated by mockery v2.14.0. DO NOT EDIT.

package mockcomponents

import mock "github.com/stretchr/testify/mock"

// ResizeInterface is an autogenerated mock type for the ResizeInterface type
type ResizeInterface struct {
	mock.Mock
}

// Resize provides a mock function with given fields: width, height
func (_m *ResizeInterface) Resize(width int, height int) {
	_m.Called(width, height)
}

type mockConstructorTestingTNewResizeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewResizeInterface creates a new instance of ResizeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResizeInterface(t mockConstructorTestingTNewResizeInterface) *ResizeInterface {
	mock := &ResizeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
