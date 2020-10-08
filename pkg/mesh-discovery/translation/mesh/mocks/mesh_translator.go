// Code generated by MockGen. DO NOT EDIT.
// Source: ./mesh_translator.go

// Package mock_mesh is a generated GoMock package.
package mock_mesh

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	input "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/input"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// TranslateMeshes mocks base method
func (m *MockTranslator) TranslateMeshes(in input.Snapshot) v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateMeshes", in)
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// TranslateMeshes indicates an expected call of TranslateMeshes
func (mr *MockTranslatorMockRecorder) TranslateMeshes(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateMeshes", reflect.TypeOf((*MockTranslator)(nil).TranslateMeshes), in)
}
