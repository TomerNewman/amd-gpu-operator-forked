// Code generated by MockGen. DO NOT EDIT.
// Source: device_config_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=device_config_reconciler.go -package=controllers -destination=mock_device_config_reconciler.go deviceConfigReconcilerHelperAPI
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/yevgeny-shnaidman/amd-gpu-operator/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
)

// MockdeviceConfigReconcilerHelperAPI is a mock of deviceConfigReconcilerHelperAPI interface.
type MockdeviceConfigReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockdeviceConfigReconcilerHelperAPIMockRecorder
}

// MockdeviceConfigReconcilerHelperAPIMockRecorder is the mock recorder for MockdeviceConfigReconcilerHelperAPI.
type MockdeviceConfigReconcilerHelperAPIMockRecorder struct {
	mock *MockdeviceConfigReconcilerHelperAPI
}

// NewMockdeviceConfigReconcilerHelperAPI creates a new mock instance.
func NewMockdeviceConfigReconcilerHelperAPI(ctrl *gomock.Controller) *MockdeviceConfigReconcilerHelperAPI {
	mock := &MockdeviceConfigReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockdeviceConfigReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdeviceConfigReconcilerHelperAPI) EXPECT() *MockdeviceConfigReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// finalizeDeviceConfig mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) finalizeDeviceConfig(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "finalizeDeviceConfig", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// finalizeDeviceConfig indicates an expected call of finalizeDeviceConfig.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) finalizeDeviceConfig(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "finalizeDeviceConfig", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).finalizeDeviceConfig), ctx, devConfig)
}

// getRequestedDeviceConfig mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) getRequestedDeviceConfig(ctx context.Context, namespacedName types.NamespacedName) (*v1alpha1.DeviceConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedDeviceConfig", ctx, namespacedName)
	ret0, _ := ret[0].(*v1alpha1.DeviceConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedDeviceConfig indicates an expected call of getRequestedDeviceConfig.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) getRequestedDeviceConfig(ctx, namespacedName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedDeviceConfig", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).getRequestedDeviceConfig), ctx, namespacedName)
}

// handleBuildConfigMap mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleBuildConfigMap(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBuildConfigMap", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleBuildConfigMap indicates an expected call of handleBuildConfigMap.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleBuildConfigMap(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBuildConfigMap", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleBuildConfigMap), ctx, devConfig)
}

// handleKMMModule mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleKMMModule(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleKMMModule", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleKMMModule indicates an expected call of handleKMMModule.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleKMMModule(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleKMMModule", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleKMMModule), ctx, devConfig)
}

// handleNodeLabeller mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleNodeLabeller(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleNodeLabeller", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleNodeLabeller indicates an expected call of handleNodeLabeller.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleNodeLabeller(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleNodeLabeller", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleNodeLabeller), ctx, devConfig)
}

// setFinalizer mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) setFinalizer(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setFinalizer", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// setFinalizer indicates an expected call of setFinalizer.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) setFinalizer(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setFinalizer", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).setFinalizer), ctx, devConfig)
}
