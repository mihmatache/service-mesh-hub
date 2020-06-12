// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/external-apis/pkg/api/k8s/apps/v1 (interfaces: DeploymentClient,ReplicaSetClient)

// Package mock_k8s_apps_clients is a generated GoMock package.
package mock_k8s_apps_clients

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	v10 "k8s.io/api/apps/v1"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockDeploymentClient is a mock of DeploymentClient interface.
type MockDeploymentClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentClientMockRecorder
}

// MockDeploymentClientMockRecorder is the mock recorder for MockDeploymentClient.
type MockDeploymentClientMockRecorder struct {
	mock *MockDeploymentClient
}

// NewMockDeploymentClient creates a new mock instance.
func NewMockDeploymentClient(ctrl *gomock.Controller) *MockDeploymentClient {
	mock := &MockDeploymentClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentClient) EXPECT() *MockDeploymentClientMockRecorder {
	return m.recorder
}

// CreateDeployment mocks base method.
func (m *MockDeploymentClient) CreateDeployment(arg0 context.Context, arg1 *v10.Deployment, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockDeploymentClientMockRecorder) CreateDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).CreateDeployment), varargs...)
}

// DeleteAllOfDeployment mocks base method.
func (m *MockDeploymentClient) DeleteAllOfDeployment(arg0 context.Context, arg1 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfDeployment indicates an expected call of DeleteAllOfDeployment.
func (mr *MockDeploymentClientMockRecorder) DeleteAllOfDeployment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).DeleteAllOfDeployment), varargs...)
}

// DeleteDeployment mocks base method.
func (m *MockDeploymentClient) DeleteDeployment(arg0 context.Context, arg1 types.NamespacedName, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment.
func (mr *MockDeploymentClientMockRecorder) DeleteDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).DeleteDeployment), varargs...)
}

// GetDeployment mocks base method.
func (m *MockDeploymentClient) GetDeployment(arg0 context.Context, arg1 types.NamespacedName) (*v10.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", arg0, arg1)
	ret0, _ := ret[0].(*v10.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockDeploymentClientMockRecorder) GetDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).GetDeployment), arg0, arg1)
}

// ListDeployment mocks base method.
func (m *MockDeploymentClient) ListDeployment(arg0 context.Context, arg1 ...client.ListOption) (*v10.DeploymentList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeployment", varargs...)
	ret0, _ := ret[0].(*v10.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployment indicates an expected call of ListDeployment.
func (mr *MockDeploymentClientMockRecorder) ListDeployment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).ListDeployment), varargs...)
}

// PatchDeployment mocks base method.
func (m *MockDeploymentClient) PatchDeployment(arg0 context.Context, arg1 *v10.Deployment, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchDeployment indicates an expected call of PatchDeployment.
func (mr *MockDeploymentClientMockRecorder) PatchDeployment(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).PatchDeployment), varargs...)
}

// PatchDeploymentStatus mocks base method.
func (m *MockDeploymentClient) PatchDeploymentStatus(arg0 context.Context, arg1 *v10.Deployment, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchDeploymentStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchDeploymentStatus indicates an expected call of PatchDeploymentStatus.
func (mr *MockDeploymentClientMockRecorder) PatchDeploymentStatus(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDeploymentStatus", reflect.TypeOf((*MockDeploymentClient)(nil).PatchDeploymentStatus), varargs...)
}

// UpdateDeployment mocks base method.
func (m *MockDeploymentClient) UpdateDeployment(arg0 context.Context, arg1 *v10.Deployment, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockDeploymentClientMockRecorder) UpdateDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).UpdateDeployment), varargs...)
}

// UpdateDeploymentStatus mocks base method.
func (m *MockDeploymentClient) UpdateDeploymentStatus(arg0 context.Context, arg1 *v10.Deployment, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDeploymentStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeploymentStatus indicates an expected call of UpdateDeploymentStatus.
func (mr *MockDeploymentClientMockRecorder) UpdateDeploymentStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeploymentStatus", reflect.TypeOf((*MockDeploymentClient)(nil).UpdateDeploymentStatus), varargs...)
}

// UpsertDeployment mocks base method.
func (m *MockDeploymentClient) UpsertDeployment(arg0 context.Context, arg1 *v10.Deployment, arg2 ...v1.DeploymentTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertDeployment indicates an expected call of UpsertDeployment.
func (mr *MockDeploymentClientMockRecorder) UpsertDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertDeployment", reflect.TypeOf((*MockDeploymentClient)(nil).UpsertDeployment), varargs...)
}

// MockReplicaSetClient is a mock of ReplicaSetClient interface.
type MockReplicaSetClient struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetClientMockRecorder
}

// MockReplicaSetClientMockRecorder is the mock recorder for MockReplicaSetClient.
type MockReplicaSetClientMockRecorder struct {
	mock *MockReplicaSetClient
}

// NewMockReplicaSetClient creates a new mock instance.
func NewMockReplicaSetClient(ctrl *gomock.Controller) *MockReplicaSetClient {
	mock := &MockReplicaSetClient{ctrl: ctrl}
	mock.recorder = &MockReplicaSetClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetClient) EXPECT() *MockReplicaSetClientMockRecorder {
	return m.recorder
}

// CreateReplicaSet mocks base method.
func (m *MockReplicaSetClient) CreateReplicaSet(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReplicaSet indicates an expected call of CreateReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) CreateReplicaSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).CreateReplicaSet), varargs...)
}

// DeleteAllOfReplicaSet mocks base method.
func (m *MockReplicaSetClient) DeleteAllOfReplicaSet(arg0 context.Context, arg1 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfReplicaSet indicates an expected call of DeleteAllOfReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) DeleteAllOfReplicaSet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).DeleteAllOfReplicaSet), varargs...)
}

// DeleteReplicaSet mocks base method.
func (m *MockReplicaSetClient) DeleteReplicaSet(arg0 context.Context, arg1 types.NamespacedName, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReplicaSet indicates an expected call of DeleteReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) DeleteReplicaSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).DeleteReplicaSet), varargs...)
}

// GetReplicaSet mocks base method.
func (m *MockReplicaSetClient) GetReplicaSet(arg0 context.Context, arg1 types.NamespacedName) (*v10.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicaSet", arg0, arg1)
	ret0, _ := ret[0].(*v10.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicaSet indicates an expected call of GetReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) GetReplicaSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).GetReplicaSet), arg0, arg1)
}

// ListReplicaSet mocks base method.
func (m *MockReplicaSetClient) ListReplicaSet(arg0 context.Context, arg1 ...client.ListOption) (*v10.ReplicaSetList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReplicaSet", varargs...)
	ret0, _ := ret[0].(*v10.ReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReplicaSet indicates an expected call of ListReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) ListReplicaSet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).ListReplicaSet), varargs...)
}

// PatchReplicaSet mocks base method.
func (m *MockReplicaSetClient) PatchReplicaSet(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchReplicaSet indicates an expected call of PatchReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) PatchReplicaSet(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).PatchReplicaSet), varargs...)
}

// PatchReplicaSetStatus mocks base method.
func (m *MockReplicaSetClient) PatchReplicaSetStatus(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchReplicaSetStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchReplicaSetStatus indicates an expected call of PatchReplicaSetStatus.
func (mr *MockReplicaSetClientMockRecorder) PatchReplicaSetStatus(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchReplicaSetStatus", reflect.TypeOf((*MockReplicaSetClient)(nil).PatchReplicaSetStatus), varargs...)
}

// UpdateReplicaSet mocks base method.
func (m *MockReplicaSetClient) UpdateReplicaSet(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicaSet indicates an expected call of UpdateReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) UpdateReplicaSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).UpdateReplicaSet), varargs...)
}

// UpdateReplicaSetStatus mocks base method.
func (m *MockReplicaSetClient) UpdateReplicaSetStatus(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateReplicaSetStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicaSetStatus indicates an expected call of UpdateReplicaSetStatus.
func (mr *MockReplicaSetClientMockRecorder) UpdateReplicaSetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicaSetStatus", reflect.TypeOf((*MockReplicaSetClient)(nil).UpdateReplicaSetStatus), varargs...)
}

// UpsertReplicaSet mocks base method.
func (m *MockReplicaSetClient) UpsertReplicaSet(arg0 context.Context, arg1 *v10.ReplicaSet, arg2 ...v1.ReplicaSetTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertReplicaSet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertReplicaSet indicates an expected call of UpsertReplicaSet.
func (mr *MockReplicaSetClientMockRecorder) UpsertReplicaSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertReplicaSet", reflect.TypeOf((*MockReplicaSetClient)(nil).UpsertReplicaSet), varargs...)
}
