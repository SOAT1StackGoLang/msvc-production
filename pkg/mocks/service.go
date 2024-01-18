// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SOAT1StackGoLang/msvc-production/pkg/service (interfaces: ProductionService)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/service.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/service ProductionService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/SOAT1StackGoLang/msvc-production/pkg/domain"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockProductionService is a mock of ProductionService interface.
type MockProductionService struct {
	ctrl     *gomock.Controller
	recorder *MockProductionServiceMockRecorder
}

// MockProductionServiceMockRecorder is the mock recorder for MockProductionService.
type MockProductionServiceMockRecorder struct {
	mock *MockProductionService
}

// NewMockProductionService creates a new mock instance.
func NewMockProductionService(ctrl *gomock.Controller) *MockProductionService {
	mock := &MockProductionService{ctrl: ctrl}
	mock.recorder = &MockProductionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductionService) EXPECT() *MockProductionServiceMockRecorder {
	return m.recorder
}

// UpdateOrderStatus mocks base method.
func (m *MockProductionService) UpdateOrderStatus(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 domain.OrderStatus) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockProductionServiceMockRecorder) UpdateOrderStatus(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockProductionService)(nil).UpdateOrderStatus), arg0, arg1, arg2, arg3)
}