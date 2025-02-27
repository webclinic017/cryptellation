// Code generated by MockGen. DO NOT EDIT.
// Source: adapter.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	backtest "github.com/digital-feather/cryptellation/services/backtests/internal/application/domains/backtest"
	gomock "github.com/golang/mock/gomock"
)

// MockAdapter is a mock of Adapter interface.
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter.
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance.
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// CreateBacktest mocks base method.
func (m *MockAdapter) CreateBacktest(ctx context.Context, bt *backtest.Backtest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBacktest", ctx, bt)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBacktest indicates an expected call of CreateBacktest.
func (mr *MockAdapterMockRecorder) CreateBacktest(ctx, bt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBacktest", reflect.TypeOf((*MockAdapter)(nil).CreateBacktest), ctx, bt)
}

// DeleteBacktest mocks base method.
func (m *MockAdapter) DeleteBacktest(ctx context.Context, bt backtest.Backtest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBacktest", ctx, bt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBacktest indicates an expected call of DeleteBacktest.
func (mr *MockAdapterMockRecorder) DeleteBacktest(ctx, bt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBacktest", reflect.TypeOf((*MockAdapter)(nil).DeleteBacktest), ctx, bt)
}

// LockedBacktest mocks base method.
func (m *MockAdapter) LockedBacktest(ctx context.Context, id uint, fn LockedBacktestCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedBacktest", ctx, id, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockedBacktest indicates an expected call of LockedBacktest.
func (mr *MockAdapterMockRecorder) LockedBacktest(ctx, id, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedBacktest", reflect.TypeOf((*MockAdapter)(nil).LockedBacktest), ctx, id, fn)
}

// ReadBacktest mocks base method.
func (m *MockAdapter) ReadBacktest(ctx context.Context, id uint) (backtest.Backtest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBacktest", ctx, id)
	ret0, _ := ret[0].(backtest.Backtest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBacktest indicates an expected call of ReadBacktest.
func (mr *MockAdapterMockRecorder) ReadBacktest(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBacktest", reflect.TypeOf((*MockAdapter)(nil).ReadBacktest), ctx, id)
}

// UpdateBacktest mocks base method.
func (m *MockAdapter) UpdateBacktest(ctx context.Context, bt backtest.Backtest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBacktest", ctx, bt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBacktest indicates an expected call of UpdateBacktest.
func (mr *MockAdapterMockRecorder) UpdateBacktest(ctx, bt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBacktest", reflect.TypeOf((*MockAdapter)(nil).UpdateBacktest), ctx, bt)
}
