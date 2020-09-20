// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// CacheClient is an autogenerated mock type for the CacheClient type
type CacheClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: key, src
func (_m *CacheClient) Get(key string, src interface{}) error {
	ret := _m.Called(key, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: key, value, expiration
func (_m *CacheClient) Set(key string, value interface{}, expiration time.Duration) error {
	ret := _m.Called(key, value, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
