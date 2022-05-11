// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Requester2 is an autogenerated mock type for the Requester2 type
type Requester2 struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester2) Get(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewRequester2T interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequester2 creates a new instance of Requester2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequester2(t NewRequester2T) *Requester2 {
	mock := &Requester2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
