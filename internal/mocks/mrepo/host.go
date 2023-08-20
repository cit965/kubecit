// Code generated by MockGen. DO NOT EDIT.
// Source: kubecit/internal/biz (interfaces: CloudHostRepo)

// Package mrepo is a generated GoMock package.
package mrepo

import (
	context "context"
	biz "kubecit/internal/biz"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCloudHostRepo is a mock of CloudHostRepo interface.
type MockCloudHostRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCloudHostRepoMockRecorder
}

// MockCloudHostRepoMockRecorder is the mock recorder for MockCloudHostRepo.
type MockCloudHostRepoMockRecorder struct {
	mock *MockCloudHostRepo
}

// NewMockCloudHostRepo creates a new mock instance.
func NewMockCloudHostRepo(ctrl *gomock.Controller) *MockCloudHostRepo {
	mock := &MockCloudHostRepo{ctrl: ctrl}
	mock.recorder = &MockCloudHostRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudHostRepo) EXPECT() *MockCloudHostRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCloudHostRepo) Create(arg0 context.Context, arg1 *biz.CloudHost) (*biz.CloudHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*biz.CloudHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCloudHostRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCloudHostRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockCloudHostRepo) Delete(arg0 context.Context, arg1 string) (*biz.CloudHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*biz.CloudHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCloudHostRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCloudHostRepo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockCloudHostRepo) Get(arg0 context.Context, arg1 string) (*biz.CloudHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*biz.CloudHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCloudHostRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCloudHostRepo)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockCloudHostRepo) List(arg0 context.Context) ([]*biz.CloudHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*biz.CloudHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCloudHostRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCloudHostRepo)(nil).List), arg0)
}

// Sync mocks base method.
func (m *MockCloudHostRepo) Sync(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockCloudHostRepoMockRecorder) Sync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockCloudHostRepo)(nil).Sync), arg0)
}

// Update mocks base method.
func (m *MockCloudHostRepo) Update(arg0 context.Context, arg1 string, arg2 *biz.CloudHost) (*biz.CloudHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*biz.CloudHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCloudHostRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCloudHostRepo)(nil).Update), arg0, arg1, arg2)
}
