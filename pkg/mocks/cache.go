// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SOAT1StackGoLang/msvc-production/pkg/service (interfaces: CacheService)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/cache.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/service CacheService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "github.com/SOAT1StackGoLang/msvc-production/pkg/service"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCacheService is a mock of CacheService interface.
type MockCacheService struct {
	ctrl     *gomock.Controller
	recorder *MockCacheServiceMockRecorder
}

// MockCacheServiceMockRecorder is the mock recorder for MockCacheService.
type MockCacheServiceMockRecorder struct {
	mock *MockCacheService
}

// NewMockCacheService creates a new mock instance.
func NewMockCacheService(ctrl *gomock.Controller) *MockCacheService {
	mock := &MockCacheService{ctrl: ctrl}
	mock.recorder = &MockCacheServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheService) EXPECT() *MockCacheServiceMockRecorder {
	return m.recorder
}

// OrderStatusChanged mocks base method.
func (m *MockCacheService) OrderStatusChanged(arg0 context.Context, arg1 any, arg2 domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderStatusChanged", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrderStatusChanged indicates an expected call of OrderStatusChanged.
func (mr *MockCacheServiceMockRecorder) OrderStatusChanged(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderStatusChanged", reflect.TypeOf((*MockCacheService)(nil).OrderStatusChanged), arg0, arg1, arg2)
}
