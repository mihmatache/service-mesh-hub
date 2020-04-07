// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/api/rbac/v1"
	rest "k8s.io/client-go/rest"
)

// MockRbacClient is a mock of RbacClient interface.
type MockRbacClient struct {
	ctrl     *gomock.Controller
	recorder *MockRbacClientMockRecorder
}

// MockRbacClientMockRecorder is the mock recorder for MockRbacClient.
type MockRbacClientMockRecorder struct {
	mock *MockRbacClient
}

// NewMockRbacClient creates a new mock instance.
func NewMockRbacClient(ctrl *gomock.Controller) *MockRbacClient {
	mock := &MockRbacClient{ctrl: ctrl}
	mock.recorder = &MockRbacClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRbacClient) EXPECT() *MockRbacClientMockRecorder {
	return m.recorder
}

// BindClusterRolesToServiceAccount mocks base method.
func (m *MockRbacClient) BindClusterRolesToServiceAccount(targetServiceAccount *v1.ServiceAccount, roles []*v10.ClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindClusterRolesToServiceAccount", targetServiceAccount, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindClusterRolesToServiceAccount indicates an expected call of BindClusterRolesToServiceAccount.
func (mr *MockRbacClientMockRecorder) BindClusterRolesToServiceAccount(targetServiceAccount, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindClusterRolesToServiceAccount", reflect.TypeOf((*MockRbacClient)(nil).BindClusterRolesToServiceAccount), targetServiceAccount, roles)
}

// MockRemoteAuthorityConfigCreator is a mock of RemoteAuthorityConfigCreator interface.
type MockRemoteAuthorityConfigCreator struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteAuthorityConfigCreatorMockRecorder
}

// MockRemoteAuthorityConfigCreatorMockRecorder is the mock recorder for MockRemoteAuthorityConfigCreator.
type MockRemoteAuthorityConfigCreatorMockRecorder struct {
	mock *MockRemoteAuthorityConfigCreator
}

// NewMockRemoteAuthorityConfigCreator creates a new mock instance.
func NewMockRemoteAuthorityConfigCreator(ctrl *gomock.Controller) *MockRemoteAuthorityConfigCreator {
	mock := &MockRemoteAuthorityConfigCreator{ctrl: ctrl}
	mock.recorder = &MockRemoteAuthorityConfigCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteAuthorityConfigCreator) EXPECT() *MockRemoteAuthorityConfigCreatorMockRecorder {
	return m.recorder
}

// ConfigFromRemoteServiceAccount mocks base method.
func (m *MockRemoteAuthorityConfigCreator) ConfigFromRemoteServiceAccount(ctx context.Context, targetClusterCfg *rest.Config, serviceAccountRef *types.ResourceRef) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFromRemoteServiceAccount", ctx, targetClusterCfg, serviceAccountRef)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigFromRemoteServiceAccount indicates an expected call of ConfigFromRemoteServiceAccount.
func (mr *MockRemoteAuthorityConfigCreatorMockRecorder) ConfigFromRemoteServiceAccount(ctx, targetClusterCfg, serviceAccountRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFromRemoteServiceAccount", reflect.TypeOf((*MockRemoteAuthorityConfigCreator)(nil).ConfigFromRemoteServiceAccount), ctx, targetClusterCfg, serviceAccountRef)
}

// MockClusterAuthorization is a mock of ClusterAuthorization interface.
type MockClusterAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAuthorizationMockRecorder
}

// MockClusterAuthorizationMockRecorder is the mock recorder for MockClusterAuthorization.
type MockClusterAuthorizationMockRecorder struct {
	mock *MockClusterAuthorization
}

// NewMockClusterAuthorization creates a new mock instance.
func NewMockClusterAuthorization(ctrl *gomock.Controller) *MockClusterAuthorization {
	mock := &MockClusterAuthorization{ctrl: ctrl}
	mock.recorder = &MockClusterAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterAuthorization) EXPECT() *MockClusterAuthorizationMockRecorder {
	return m.recorder
}

// CreateAuthConfigForCluster mocks base method.
func (m *MockClusterAuthorization) CreateAuthConfigForCluster(ctx context.Context, targetClusterCfg *rest.Config, serviceAccountRef *types.ResourceRef) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthConfigForCluster", ctx, targetClusterCfg, serviceAccountRef)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthConfigForCluster indicates an expected call of CreateAuthConfigForCluster.
func (mr *MockClusterAuthorizationMockRecorder) CreateAuthConfigForCluster(ctx, targetClusterCfg, serviceAccountRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthConfigForCluster", reflect.TypeOf((*MockClusterAuthorization)(nil).CreateAuthConfigForCluster), ctx, targetClusterCfg, serviceAccountRef)
}

// MockRemoteAuthorityManager is a mock of RemoteAuthorityManager interface.
type MockRemoteAuthorityManager struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteAuthorityManagerMockRecorder
}

// MockRemoteAuthorityManagerMockRecorder is the mock recorder for MockRemoteAuthorityManager.
type MockRemoteAuthorityManagerMockRecorder struct {
	mock *MockRemoteAuthorityManager
}

// NewMockRemoteAuthorityManager creates a new mock instance.
func NewMockRemoteAuthorityManager(ctrl *gomock.Controller) *MockRemoteAuthorityManager {
	mock := &MockRemoteAuthorityManager{ctrl: ctrl}
	mock.recorder = &MockRemoteAuthorityManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteAuthorityManager) EXPECT() *MockRemoteAuthorityManagerMockRecorder {
	return m.recorder
}

// ApplyRemoteServiceAccount mocks base method.
func (m *MockRemoteAuthorityManager) ApplyRemoteServiceAccount(ctx context.Context, newServiceAccountRef *types.ResourceRef, roles []*v10.ClusterRole) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyRemoteServiceAccount", ctx, newServiceAccountRef, roles)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyRemoteServiceAccount indicates an expected call of ApplyRemoteServiceAccount.
func (mr *MockRemoteAuthorityManagerMockRecorder) ApplyRemoteServiceAccount(ctx, newServiceAccountRef, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRemoteServiceAccount", reflect.TypeOf((*MockRemoteAuthorityManager)(nil).ApplyRemoteServiceAccount), ctx, newServiceAccountRef, roles)
}
