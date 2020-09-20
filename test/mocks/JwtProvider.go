// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// JwtProvider is an autogenerated mock type for the JwtProvider type
type JwtProvider struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: userId, secret, expires
func (_m *JwtProvider) CreateToken(userId string, secret []byte, expires time.Duration) (string, error) {
	ret := _m.Called(userId, secret, expires)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte, time.Duration) string); ok {
		r0 = rf(userId, secret, expires)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, time.Duration) error); ok {
		r1 = rf(userId, secret, expires)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeToken provides a mock function with given fields: token, secret
func (_m *JwtProvider) DecodeToken(token string, secret []byte) (string, error) {
	ret := _m.Called(token, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte) string); ok {
		r0 = rf(token, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(token, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
