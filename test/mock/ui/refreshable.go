// Code generated by mockery v2.19.0. DO NOT EDIT.

package mockui

import (
	tea "github.com/charmbracelet/bubbletea"
	mock "github.com/stretchr/testify/mock"
)

// Refreshable is an autogenerated mock type for the Refreshable type
type Refreshable struct {
	mock.Mock
}

// FetchTasks provides a mock function with given fields: isHardRefresh
func (_m *Refreshable) FetchTasks(isHardRefresh bool) tea.Cmd {
	ret := _m.Called(isHardRefresh)

	var r0 tea.Cmd
	if rf, ok := ret.Get(0).(func(bool) tea.Cmd); ok {
		r0 = rf(isHardRefresh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Cmd)
		}
	}

	return r0
}

type mockConstructorTestingTNewRefreshable interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefreshable creates a new instance of Refreshable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefreshable(t mockConstructorTestingTNewRefreshable) *Refreshable {
	mock := &Refreshable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
