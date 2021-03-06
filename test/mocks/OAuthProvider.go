// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// OAuthProvider is an autogenerated mock type for the OAuthProvider type
type OAuthProvider struct {
	mock.Mock
}

// AuthCodeURL provides a mock function with given fields: state, opts
func (_m *OAuthProvider) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, state)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...oauth2.AuthCodeOption) string); ok {
		r0 = rf(state, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Exchange provides a mock function with given fields: ctx, code, opts
func (_m *OAuthProvider) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, code)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oauth2.AuthCodeOption) *oauth2.Token); ok {
		r0 = rf(ctx, code, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...oauth2.AuthCodeOption) error); ok {
		r1 = rf(ctx, code, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
