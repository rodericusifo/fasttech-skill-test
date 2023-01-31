// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	mock "github.com/stretchr/testify/mock"

	types "github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
)

// ICartResource is an autogenerated mock type for the ICartResource type
type ICartResource struct {
	mock.Mock
}

// CreateCart provides a mock function with given fields: payload
func (_m *ICartResource) CreateCart(payload *model.Cart) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Cart) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCart provides a mock function with given fields: payload
func (_m *ICartResource) DeleteCart(payload *model.Cart) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Cart) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCart provides a mock function with given fields: query, payload
func (_m *ICartResource) GetCart(query *types.Query, payload *model.Cart) (*model.Cart, error) {
	ret := _m.Called(query, payload)

	var r0 *model.Cart
	if rf, ok := ret.Get(0).(func(*types.Query, *model.Cart) *model.Cart); ok {
		r0 = rf(query, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Query, *model.Cart) error); ok {
		r1 = rf(query, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListCart provides a mock function with given fields: query, payload
func (_m *ICartResource) GetListCart(query *types.Query, payload *model.Cart) ([]*model.Cart, error) {
	ret := _m.Called(query, payload)

	var r0 []*model.Cart
	if rf, ok := ret.Get(0).(func(*types.Query, *model.Cart) []*model.Cart); ok {
		r0 = rf(query, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Query, *model.Cart) error); ok {
		r1 = rf(query, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCart provides a mock function with given fields: payload
func (_m *ICartResource) UpdateCart(payload *model.Cart) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Cart) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewICartResource interface {
	mock.TestingT
	Cleanup(func())
}

// NewICartResource creates a new instance of ICartResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICartResource(t mockConstructorTestingTNewICartResource) *ICartResource {
	mock := &ICartResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
