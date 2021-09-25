// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	ent "github.com/gitploy-io/gitploy/ent"
	vo "github.com/gitploy-io/gitploy/vo"
	gomock "github.com/golang/mock/gomock"
)

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CheckNotificationRecordOfEvent mocks base method.
func (m *MockInteractor) CheckNotificationRecordOfEvent(ctx context.Context, e *ent.Event) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckNotificationRecordOfEvent", ctx, e)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckNotificationRecordOfEvent indicates an expected call of CheckNotificationRecordOfEvent.
func (mr *MockInteractorMockRecorder) CheckNotificationRecordOfEvent(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckNotificationRecordOfEvent", reflect.TypeOf((*MockInteractor)(nil).CheckNotificationRecordOfEvent), ctx, e)
}

// CreateApproval mocks base method.
func (m *MockInteractor) CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApproval", ctx, a)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApproval indicates an expected call of CreateApproval.
func (mr *MockInteractorMockRecorder) CreateApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApproval", reflect.TypeOf((*MockInteractor)(nil).CreateApproval), ctx, a)
}

// CreateCallback mocks base method.
func (m *MockInteractor) CreateCallback(ctx context.Context, cb *ent.Callback) (*ent.Callback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCallback", ctx, cb)
	ret0, _ := ret[0].(*ent.Callback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCallback indicates an expected call of CreateCallback.
func (mr *MockInteractorMockRecorder) CreateCallback(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCallback", reflect.TypeOf((*MockInteractor)(nil).CreateCallback), ctx, cb)
}

// CreateChatUser mocks base method.
func (m *MockInteractor) CreateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatUser", ctx, cu)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatUser indicates an expected call of CreateChatUser.
func (mr *MockInteractorMockRecorder) CreateChatUser(ctx, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatUser", reflect.TypeOf((*MockInteractor)(nil).CreateChatUser), ctx, cu)
}

// CreateEvent mocks base method.
func (m *MockInteractor) CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, e)
	ret0, _ := ret[0].(*ent.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockInteractorMockRecorder) CreateEvent(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockInteractor)(nil).CreateEvent), ctx, e)
}

// DeleteChatUser mocks base method.
func (m *MockInteractor) DeleteChatUser(ctx context.Context, cu *ent.ChatUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChatUser", ctx, cu)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChatUser indicates an expected call of DeleteChatUser.
func (mr *MockInteractorMockRecorder) DeleteChatUser(ctx, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatUser", reflect.TypeOf((*MockInteractor)(nil).DeleteChatUser), ctx, cu)
}

// Deploy mocks base method.
func (m *MockInteractor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, u, re, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockInteractorMockRecorder) Deploy(ctx, u, re, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInteractor)(nil).Deploy), ctx, u, re, d, env)
}

// FindCallbackByHash mocks base method.
func (m *MockInteractor) FindCallbackByHash(ctx context.Context, hash string) (*ent.Callback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCallbackByHash", ctx, hash)
	ret0, _ := ret[0].(*ent.Callback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCallbackByHash indicates an expected call of FindCallbackByHash.
func (mr *MockInteractorMockRecorder) FindCallbackByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCallbackByHash", reflect.TypeOf((*MockInteractor)(nil).FindCallbackByHash), ctx, hash)
}

// FindChatUserByID mocks base method.
func (m *MockInteractor) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatUserByID", ctx, id)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatUserByID indicates an expected call of FindChatUserByID.
func (mr *MockInteractorMockRecorder) FindChatUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatUserByID", reflect.TypeOf((*MockInteractor)(nil).FindChatUserByID), ctx, id)
}

// FindDeploymentByID mocks base method.
func (m *MockInteractor) FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentByID", ctx, id)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentByID indicates an expected call of FindDeploymentByID.
func (mr *MockInteractorMockRecorder) FindDeploymentByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentByID", reflect.TypeOf((*MockInteractor)(nil).FindDeploymentByID), ctx, id)
}

// FindPermOfRepo mocks base method.
func (m *MockInteractor) FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPermOfRepo", ctx, r, u)
	ret0, _ := ret[0].(*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPermOfRepo indicates an expected call of FindPermOfRepo.
func (mr *MockInteractorMockRecorder) FindPermOfRepo(ctx, r, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPermOfRepo", reflect.TypeOf((*MockInteractor)(nil).FindPermOfRepo), ctx, r, u)
}

// FindRepoOfUserByNamespaceName mocks base method.
func (m *MockInteractor) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByNamespaceName", ctx, u, namespace, name)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByNamespaceName indicates an expected call of FindRepoOfUserByNamespaceName.
func (mr *MockInteractorMockRecorder) FindRepoOfUserByNamespaceName(ctx, u, namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByNamespaceName", reflect.TypeOf((*MockInteractor)(nil).FindRepoOfUserByNamespaceName), ctx, u, namespace, name)
}

// FindUserByID mocks base method.
func (m *MockInteractor) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, id)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockInteractorMockRecorder) FindUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockInteractor)(nil).FindUserByID), ctx, id)
}

// GetBranch mocks base method.
func (m *MockInteractor) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, u, r, branch)
	ret0, _ := ret[0].(*vo.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockInteractorMockRecorder) GetBranch(ctx, u, r, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockInteractor)(nil).GetBranch), ctx, u, r, branch)
}

// GetCommit mocks base method.
func (m *MockInteractor) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, u, r, sha)
	ret0, _ := ret[0].(*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockInteractorMockRecorder) GetCommit(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockInteractor)(nil).GetCommit), ctx, u, r, sha)
}

// GetConfig mocks base method.
func (m *MockInteractor) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, u, r)
	ret0, _ := ret[0].(*vo.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockInteractorMockRecorder) GetConfig(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockInteractor)(nil).GetConfig), ctx, u, r)
}

// GetNextDeploymentNumberOfRepo mocks base method.
func (m *MockInteractor) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextDeploymentNumberOfRepo", ctx, r)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextDeploymentNumberOfRepo indicates an expected call of GetNextDeploymentNumberOfRepo.
func (mr *MockInteractorMockRecorder) GetNextDeploymentNumberOfRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextDeploymentNumberOfRepo", reflect.TypeOf((*MockInteractor)(nil).GetNextDeploymentNumberOfRepo), ctx, r)
}

// GetTag mocks base method.
func (m *MockInteractor) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, u, r, tag)
	ret0, _ := ret[0].(*vo.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockInteractorMockRecorder) GetTag(ctx, u, r, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockInteractor)(nil).GetTag), ctx, u, r, tag)
}

// ListDeploymentsOfRepo mocks base method.
func (m *MockInteractor) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env, status string, page, perPage int) ([]*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentsOfRepo", ctx, r, env, status, page, perPage)
	ret0, _ := ret[0].([]*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentsOfRepo indicates an expected call of ListDeploymentsOfRepo.
func (mr *MockInteractorMockRecorder) ListDeploymentsOfRepo(ctx, r, env, status, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListDeploymentsOfRepo), ctx, r, env, status, page, perPage)
}

// ListPermsOfRepo mocks base method.
func (m *MockInteractor) ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPermsOfRepo", ctx, r, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermsOfRepo indicates an expected call of ListPermsOfRepo.
func (mr *MockInteractorMockRecorder) ListPermsOfRepo(ctx, r, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListPermsOfRepo), ctx, r, q, page, perPage)
}

// SubscribeEvent mocks base method.
func (m *MockInteractor) SubscribeEvent(fn func(*ent.Event)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeEvent", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeEvent indicates an expected call of SubscribeEvent.
func (mr *MockInteractorMockRecorder) SubscribeEvent(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEvent", reflect.TypeOf((*MockInteractor)(nil).SubscribeEvent), fn)
}

// UnsubscribeEvent mocks base method.
func (m *MockInteractor) UnsubscribeEvent(fn func(*ent.Event)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeEvent", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeEvent indicates an expected call of UnsubscribeEvent.
func (mr *MockInteractorMockRecorder) UnsubscribeEvent(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeEvent", reflect.TypeOf((*MockInteractor)(nil).UnsubscribeEvent), fn)
}

// UpdateChatUser mocks base method.
func (m *MockInteractor) UpdateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChatUser", ctx, cu)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChatUser indicates an expected call of UpdateChatUser.
func (mr *MockInteractorMockRecorder) UpdateChatUser(ctx, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChatUser", reflect.TypeOf((*MockInteractor)(nil).UpdateChatUser), ctx, cu)
}

// UpdateRepo mocks base method.
func (m *MockInteractor) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRepo", ctx, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRepo indicates an expected call of UpdateRepo.
func (mr *MockInteractorMockRecorder) UpdateRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRepo", reflect.TypeOf((*MockInteractor)(nil).UpdateRepo), ctx, r)
}
