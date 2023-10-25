// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// URLGetter is an autogenerated mock type for the URLGetter type
type URLGetter struct {
	mock.Mock
}

// GetURL provides a mock function with given fields: alias
func (_m *URLGetter) GetURL(alias string) (string, error) {
	ret := _m.Called(alias)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewURLGetter creates a new instance of URLGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewURLGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *URLGetter {
	mock := &URLGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
