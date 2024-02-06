// Code generated by MockGen. DO NOT EDIT.
// Source: nodemetrics.go
//
// Generated by this command:
//
//	mockgen -source=nodemetrics.go -package=nodemetrics -destination=mock_nodemetrics.go NodeMetrics
//
// Package nodemetrics is a generated GoMock package.
package nodemetrics

import (
	reflect "reflect"

	v1alpha1 "github.com/yevgeny-shnaidman/amd-gpu-operator/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/apps/v1"
)

// MockNodeMetrics is a mock of NodeMetrics interface.
type MockNodeMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMetricsMockRecorder
}

// MockNodeMetricsMockRecorder is the mock recorder for MockNodeMetrics.
type MockNodeMetricsMockRecorder struct {
	mock *MockNodeMetrics
}

// NewMockNodeMetrics creates a new mock instance.
func NewMockNodeMetrics(ctrl *gomock.Controller) *MockNodeMetrics {
	mock := &MockNodeMetrics{ctrl: ctrl}
	mock.recorder = &MockNodeMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeMetrics) EXPECT() *MockNodeMetricsMockRecorder {
	return m.recorder
}

// SetNodeMetricsAsDesired mocks base method.
func (m *MockNodeMetrics) SetNodeMetricsAsDesired(ds *v1.DaemonSet, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNodeMetricsAsDesired", ds, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNodeMetricsAsDesired indicates an expected call of SetNodeMetricsAsDesired.
func (mr *MockNodeMetricsMockRecorder) SetNodeMetricsAsDesired(ds, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeMetricsAsDesired", reflect.TypeOf((*MockNodeMetrics)(nil).SetNodeMetricsAsDesired), ds, devConfig)
}