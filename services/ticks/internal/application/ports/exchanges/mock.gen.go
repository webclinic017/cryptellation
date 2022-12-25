// Code generated by MockGen. DO NOT EDIT.
// Source: adapter.go

// Package exchanges is a generated GoMock package.
package exchanges

import (
	tick "github.com/digital-feather/cryptellation/services/ticks/pkg/models/tick"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdapter is a mock of Adapter interface
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// ListenSymbol mocks base method
func (m *MockAdapter) ListenSymbol(symbol string) (chan tick.Tick, chan struct{}, error) {
	ret := m.ctrl.Call(m, "ListenSymbol", symbol)
	ret0, _ := ret[0].(chan tick.Tick)
	ret1, _ := ret[1].(chan struct{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListenSymbol indicates an expected call of ListenSymbol
func (mr *MockAdapterMockRecorder) ListenSymbol(symbol interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenSymbol", reflect.TypeOf((*MockAdapter)(nil).ListenSymbol), symbol)
}
