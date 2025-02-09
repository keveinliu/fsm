// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/flomesh-io/fsm/pkg/gateway/processor (interfaces: Processor)

// Package processor is a generated GoMock package.
package processor

import (
	reflect "reflect"

	v1alpha2 "github.com/flomesh-io/fsm/pkg/apis/policyattachment/v1alpha2"
	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	v1alpha20 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// MockProcessor is a mock of Processor interface.
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor.
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance.
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// BuildConfigs mocks base method.
func (m *MockProcessor) BuildConfigs() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BuildConfigs")
}

// BuildConfigs indicates an expected call of BuildConfigs.
func (mr *MockProcessorMockRecorder) BuildConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildConfigs", reflect.TypeOf((*MockProcessor)(nil).BuildConfigs))
}

// Delete mocks base method.
func (m *MockProcessor) Delete(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProcessorMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProcessor)(nil).Delete), arg0)
}

// Insert mocks base method.
func (m *MockProcessor) Insert(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockProcessorMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockProcessor)(nil).Insert), arg0)
}

// IsConfigMapReferred mocks base method.
func (m *MockProcessor) IsConfigMapReferred(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConfigMapReferred", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConfigMapReferred indicates an expected call of IsConfigMapReferred.
func (mr *MockProcessorMockRecorder) IsConfigMapReferred(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConfigMapReferred", reflect.TypeOf((*MockProcessor)(nil).IsConfigMapReferred), arg0)
}

// IsEffectiveRoute mocks base method.
func (m *MockProcessor) IsEffectiveRoute(arg0 []v1.ParentReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEffectiveRoute", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEffectiveRoute indicates an expected call of IsEffectiveRoute.
func (mr *MockProcessorMockRecorder) IsEffectiveRoute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEffectiveRoute", reflect.TypeOf((*MockProcessor)(nil).IsEffectiveRoute), arg0)
}

// IsEffectiveTargetRef mocks base method.
func (m *MockProcessor) IsEffectiveTargetRef(arg0 client.Object, arg1 v1alpha20.NamespacedPolicyTargetReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEffectiveTargetRef", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEffectiveTargetRef indicates an expected call of IsEffectiveTargetRef.
func (mr *MockProcessorMockRecorder) IsEffectiveTargetRef(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEffectiveTargetRef", reflect.TypeOf((*MockProcessor)(nil).IsEffectiveTargetRef), arg0, arg1)
}

// IsFilterConfigReferred mocks base method.
func (m *MockProcessor) IsFilterConfigReferred(arg0 string, arg1 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFilterConfigReferred", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFilterConfigReferred indicates an expected call of IsFilterConfigReferred.
func (mr *MockProcessorMockRecorder) IsFilterConfigReferred(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFilterConfigReferred", reflect.TypeOf((*MockProcessor)(nil).IsFilterConfigReferred), arg0, arg1)
}

// IsFilterDefinitionReferred mocks base method.
func (m *MockProcessor) IsFilterDefinitionReferred(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFilterDefinitionReferred", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFilterDefinitionReferred indicates an expected call of IsFilterDefinitionReferred.
func (mr *MockProcessorMockRecorder) IsFilterDefinitionReferred(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFilterDefinitionReferred", reflect.TypeOf((*MockProcessor)(nil).IsFilterDefinitionReferred), arg0)
}

// IsFilterReferred mocks base method.
func (m *MockProcessor) IsFilterReferred(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFilterReferred", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFilterReferred indicates an expected call of IsFilterReferred.
func (mr *MockProcessorMockRecorder) IsFilterReferred(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFilterReferred", reflect.TypeOf((*MockProcessor)(nil).IsFilterReferred), arg0)
}

// IsHeadlessServiceWithoutSelector mocks base method.
func (m *MockProcessor) IsHeadlessServiceWithoutSelector(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHeadlessServiceWithoutSelector", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHeadlessServiceWithoutSelector indicates an expected call of IsHeadlessServiceWithoutSelector.
func (mr *MockProcessorMockRecorder) IsHeadlessServiceWithoutSelector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHeadlessServiceWithoutSelector", reflect.TypeOf((*MockProcessor)(nil).IsHeadlessServiceWithoutSelector), arg0)
}

// IsListenerFilterReferred mocks base method.
func (m *MockProcessor) IsListenerFilterReferred(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsListenerFilterReferred", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsListenerFilterReferred indicates an expected call of IsListenerFilterReferred.
func (mr *MockProcessorMockRecorder) IsListenerFilterReferred(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsListenerFilterReferred", reflect.TypeOf((*MockProcessor)(nil).IsListenerFilterReferred), arg0)
}

// IsRoutableLocalTargetServices mocks base method.
func (m *MockProcessor) IsRoutableLocalTargetServices(arg0 client.Object, arg1 []v1alpha20.LocalPolicyTargetReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRoutableLocalTargetServices", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRoutableLocalTargetServices indicates an expected call of IsRoutableLocalTargetServices.
func (mr *MockProcessorMockRecorder) IsRoutableLocalTargetServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRoutableLocalTargetServices", reflect.TypeOf((*MockProcessor)(nil).IsRoutableLocalTargetServices), arg0, arg1)
}

// IsRoutableNamespacedTargetServices mocks base method.
func (m *MockProcessor) IsRoutableNamespacedTargetServices(arg0 client.Object, arg1 []v1alpha20.NamespacedPolicyTargetReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRoutableNamespacedTargetServices", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRoutableNamespacedTargetServices indicates an expected call of IsRoutableNamespacedTargetServices.
func (mr *MockProcessorMockRecorder) IsRoutableNamespacedTargetServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRoutableNamespacedTargetServices", reflect.TypeOf((*MockProcessor)(nil).IsRoutableNamespacedTargetServices), arg0, arg1)
}

// IsRoutableService mocks base method.
func (m *MockProcessor) IsRoutableService(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRoutableService", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRoutableService indicates an expected call of IsRoutableService.
func (mr *MockProcessorMockRecorder) IsRoutableService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRoutableService", reflect.TypeOf((*MockProcessor)(nil).IsRoutableService), arg0)
}

// IsRoutableTargetService mocks base method.
func (m *MockProcessor) IsRoutableTargetService(arg0 client.Object, arg1 v1alpha20.NamespacedPolicyTargetReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRoutableTargetService", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRoutableTargetService indicates an expected call of IsRoutableTargetService.
func (mr *MockProcessorMockRecorder) IsRoutableTargetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRoutableTargetService", reflect.TypeOf((*MockProcessor)(nil).IsRoutableTargetService), arg0, arg1)
}

// IsSecretReferred mocks base method.
func (m *MockProcessor) IsSecretReferred(arg0 types.NamespacedName) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSecretReferred", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecretReferred indicates an expected call of IsSecretReferred.
func (mr *MockProcessorMockRecorder) IsSecretReferred(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSecretReferred", reflect.TypeOf((*MockProcessor)(nil).IsSecretReferred), arg0)
}

// IsValidLocalTargetRoutes mocks base method.
func (m *MockProcessor) IsValidLocalTargetRoutes(arg0 client.Object, arg1 []v1alpha2.LocalFilterPolicyTargetReference) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidLocalTargetRoutes", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidLocalTargetRoutes indicates an expected call of IsValidLocalTargetRoutes.
func (mr *MockProcessorMockRecorder) IsValidLocalTargetRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidLocalTargetRoutes", reflect.TypeOf((*MockProcessor)(nil).IsValidLocalTargetRoutes), arg0, arg1)
}

// UseEndpointSlices mocks base method.
func (m *MockProcessor) UseEndpointSlices() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseEndpointSlices")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseEndpointSlices indicates an expected call of UseEndpointSlices.
func (mr *MockProcessorMockRecorder) UseEndpointSlices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseEndpointSlices", reflect.TypeOf((*MockProcessor)(nil).UseEndpointSlices))
}
