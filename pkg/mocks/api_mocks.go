// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SOAT1StackGoLang/msvc-production/pkg/api (interfaces: ProductionAPI)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/api_mocks.go -package=mocks github.com/SOAT1StackGoLang/msvc-production/pkg/api ProductionAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	api "github.com/SOAT1StackGoLang/msvc-production/pkg/api"
	gomock "go.uber.org/mock/gomock"
)

// MockProductionAPI is a mock of ProductionAPI interface.
type MockProductionAPI struct {
	ctrl     *gomock.Controller
	recorder *MockProductionAPIMockRecorder
}

// MockProductionAPIMockRecorder is the mock recorder for MockProductionAPI.
type MockProductionAPIMockRecorder struct {
	mock *MockProductionAPI
}

// NewMockProductionAPI creates a new mock instance.
func NewMockProductionAPI(ctrl *gomock.Controller) *MockProductionAPI {
	mock := &MockProductionAPI{ctrl: ctrl}
	mock.recorder = &MockProductionAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductionAPI) EXPECT() *MockProductionAPIMockRecorder {
	return m.recorder
}

// UpdateOrder mocks base method.
func (m *MockProductionAPI) UpdateOrder(arg0 api.UpdateOrderRequest) (api.UpdateOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0)
	ret0, _ := ret[0].(api.UpdateOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockProductionAPIMockRecorder) UpdateOrder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockProductionAPI)(nil).UpdateOrder), arg0)
}
