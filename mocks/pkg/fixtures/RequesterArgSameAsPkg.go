// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequesterArgSameAsPkg is an autogenerated mock type for the RequesterArgSameAsPkg type
type RequesterArgSameAsPkg struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsPkg) Get(_a0 string) {
	_m.Called(_a0)
}

type NewRequesterArgSameAsPkgT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterArgSameAsPkg creates a new instance of RequesterArgSameAsPkg. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterArgSameAsPkg(t NewRequesterArgSameAsPkgT) *RequesterArgSameAsPkg {
	mock := &RequesterArgSameAsPkg{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
