// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocksidebar

import (
	sidebar "github.com/alswl/go-toodledo/pkg/ui/sidebar"
	mock "github.com/stretchr/testify/mock"
)

// ItemChangeSubscriber is an autogenerated mock type for the ItemChangeSubscriber type
type ItemChangeSubscriber struct {
	mock.Mock
}

// Execute provides a mock function with given fields: tab, item
func (_m *ItemChangeSubscriber) Execute(tab string, item sidebar.Item) error {
	ret := _m.Called(tab, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, sidebar.Item) error); ok {
		r0 = rf(tab, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewItemChangeSubscriber interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemChangeSubscriber creates a new instance of ItemChangeSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemChangeSubscriber(t mockConstructorTestingTNewItemChangeSubscriber) *ItemChangeSubscriber {
	mock := &ItemChangeSubscriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
