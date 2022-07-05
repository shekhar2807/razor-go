// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PathUtils is an autogenerated mock type for the PathUtils type
type PathUtils struct {
	mock.Mock
}

// GetDefaultPath provides a mock function with given fields:
func (_m *PathUtils) GetDefaultPath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobFilePath provides a mock function with given fields:
func (_m *PathUtils) GetJobFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPathUtils interface {
	mock.TestingT
	Cleanup(func())
}

// NewPathUtils creates a new instance of PathUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPathUtils(t mockConstructorTestingTNewPathUtils) *PathUtils {
	mock := &PathUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
