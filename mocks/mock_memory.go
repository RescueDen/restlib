// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reaction-eng/restlib/cache (interfaces: RawMemoryCache)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRawMemoryCache is a mock of RawMemoryCache interface
type MockRawMemoryCache struct {
	ctrl     *gomock.Controller
	recorder *MockRawMemoryCacheMockRecorder
}

// MockRawMemoryCacheMockRecorder is the mock recorder for MockRawMemoryCache
type MockRawMemoryCacheMockRecorder struct {
	mock *MockRawMemoryCache
}

// NewMockRawMemoryCache creates a new mock instance
func NewMockRawMemoryCache(ctrl *gomock.Controller) *MockRawMemoryCache {
	mock := &MockRawMemoryCache{ctrl: ctrl}
	mock.recorder = &MockRawMemoryCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRawMemoryCache) EXPECT() *MockRawMemoryCacheMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRawMemoryCache) Get(arg0 string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRawMemoryCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRawMemoryCache)(nil).Get), arg0)
}

// SetDefault mocks base method
func (m *MockRawMemoryCache) SetDefault(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDefault", arg0, arg1)
}

// SetDefault indicates an expected call of SetDefault
func (mr *MockRawMemoryCacheMockRecorder) SetDefault(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockRawMemoryCache)(nil).SetDefault), arg0, arg1)
}
