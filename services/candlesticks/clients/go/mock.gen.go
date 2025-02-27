// Code generated by MockGen. DO NOT EDIT.
// Source: interfacer.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	candlestick "github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/candlestick"
	gomock "github.com/golang/mock/gomock"
)

// MockInterfacer is a mock of Interfacer interface.
type MockInterfacer struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacerMockRecorder
}

// MockInterfacerMockRecorder is the mock recorder for MockInterfacer.
type MockInterfacerMockRecorder struct {
	mock *MockInterfacer
}

// NewMockInterfacer creates a new mock instance.
func NewMockInterfacer(ctrl *gomock.Controller) *MockInterfacer {
	mock := &MockInterfacer{ctrl: ctrl}
	mock.recorder = &MockInterfacerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfacer) EXPECT() *MockInterfacerMockRecorder {
	return m.recorder
}

// ReadCandlesticks mocks base method.
func (m *MockInterfacer) ReadCandlesticks(ctx context.Context, payload ReadCandlestickPayload) (*candlestick.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCandlesticks", ctx, payload)
	ret0, _ := ret[0].(*candlestick.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCandlesticks indicates an expected call of ReadCandlesticks.
func (mr *MockInterfacerMockRecorder) ReadCandlesticks(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCandlesticks", reflect.TypeOf((*MockInterfacer)(nil).ReadCandlesticks), ctx, payload)
}
