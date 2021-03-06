// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIndex is a mock of Index interface
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockIndex) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockIndexMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIndex)(nil).Get), key)
}

// GetTwo mocks base method
func (m *MockIndex) GetTwo(key1, key2 string) (interface{}, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwo", key1, key2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// GetTwo indicates an expected call of GetTwo
func (mr *MockIndexMockRecorder) GetTwo(key1, key2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwo", reflect.TypeOf((*MockIndex)(nil).GetTwo), key1, key2)
}

// Put mocks base method
func (m *MockIndex) Put(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", key, value)
}

// Put indicates an expected call of Put
func (mr *MockIndexMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIndex)(nil).Put), key, value)
}
