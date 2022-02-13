// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocksyncer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ToodledoSyncer is an autogenerated mock type for the ToodledoSyncer type
type ToodledoSyncer struct {
	mock.Mock
}

// Start provides a mock function with given fields: _a0
func (_m *ToodledoSyncer) Start(_a0 context.Context) {
	_m.Called(_a0)
}

// Stop provides a mock function with given fields:
func (_m *ToodledoSyncer) Stop() {
	_m.Called()
}

// SyncOnce provides a mock function with given fields:
func (_m *ToodledoSyncer) SyncOnce() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// sync provides a mock function with given fields:
func (_m *ToodledoSyncer) sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
