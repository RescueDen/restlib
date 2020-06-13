// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reaction-eng/restlib/images (interfaces: Host)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// Location mocks base method
func (m *MockHost) Location() http.Dir {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(http.Dir)
	return ret0
}

// Location indicates an expected call of Location
func (mr *MockHostMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockHost)(nil).Location))
}

// PrepareImage mocks base method
func (m *MockHost) PrepareImage(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareImage indicates an expected call of PrepareImage
func (mr *MockHostMockRecorder) PrepareImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareImage", reflect.TypeOf((*MockHost)(nil).PrepareImage), arg0)
}