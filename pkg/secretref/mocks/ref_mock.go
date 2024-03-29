// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RefClient is an autogenerated mock type for the RefClient type
type RefClient struct {
	mock.Mock
}

// GetSecretFromRef provides a mock function with given fields: ctx, refVal, secretNamespace
func (_m *RefClient) GetSecretFromRef(ctx context.Context, refVal string, secretNamespace string) (string, error) {
	ret := _m.Called(ctx, refVal, secretNamespace)

	if len(ret) == 0 {
		panic("no return value specified for GetSecretFromRef")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, refVal, secretNamespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, refVal, secretNamespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, refVal, secretNamespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MapComponentConfigSecretsRefs provides a mock function with given fields: ctx, config, namespace
func (_m *RefClient) MapComponentConfigSecretsRefs(ctx context.Context, config map[string][]string, namespace string) error {
	ret := _m.Called(ctx, config, namespace)

	if len(ret) == 0 {
		panic("no return value specified for MapComponentConfigSecretsRefs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string][]string, string) error); ok {
		r0 = rf(ctx, config, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MapConfigSecretsRefs provides a mock function with given fields: ctx, config, namespace
func (_m *RefClient) MapConfigSecretsRefs(ctx context.Context, config map[string]string, namespace string) error {
	ret := _m.Called(ctx, config, namespace)

	if len(ret) == 0 {
		panic("no return value specified for MapConfigSecretsRefs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string) error); ok {
		r0 = rf(ctx, config, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRefClient creates a new instance of RefClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRefClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *RefClient {
	mock := &RefClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
