// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	types "github.com/skip-mev/slinky/providers/types"
	mock "github.com/stretchr/testify/mock"
)

// APIDataHandler is an autogenerated mock type for the APIDataHandler type
type APIDataHandler[K types.ResponseKey, V types.ResponseValue] struct {
	mock.Mock
}

// CreateURL provides a mock function with given fields: ids
func (_m *APIDataHandler[K, V]) CreateURL(ids []K) (string, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for CreateURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]K) (string, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]K) string); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]K) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseResponse provides a mock function with given fields: ids, response
func (_m *APIDataHandler[K, V]) ParseResponse(ids []K, response *http.Response) types.GetResponse[K, V] {
	ret := _m.Called(ids, response)

	if len(ret) == 0 {
		panic("no return value specified for ParseResponse")
	}

	var r0 types.GetResponse[K, V]
	if rf, ok := ret.Get(0).(func([]K, *http.Response) types.GetResponse[K, V]); ok {
		r0 = rf(ids, response)
	} else {
		r0 = ret.Get(0).(types.GetResponse[K, V])
	}

	return r0
}

// NewAPIDataHandler creates a new instance of APIDataHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIDataHandler[K types.ResponseKey, V types.ResponseValue](t interface {
	mock.TestingT
	Cleanup(func())
}) *APIDataHandler[K, V] {
	mock := &APIDataHandler[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
