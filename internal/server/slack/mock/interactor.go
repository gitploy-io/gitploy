// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/hanjunlee/gitploy/ent"
	notification "github.com/hanjunlee/gitploy/ent/notification"
	vo "github.com/hanjunlee/gitploy/vo"
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

// CloseChatCallback mocks base method.
func (m *MockInteractor) CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseChatCallback", ctx, cb)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseChatCallback indicates an expected call of CloseChatCallback.
func (mr *MockInteractorMockRecorder) CloseChatCallback(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseChatCallback", reflect.TypeOf((*MockInteractor)(nil).CloseChatCallback), ctx, cb)
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

// CreateChatCallback mocks base method.
func (m *MockInteractor) CreateChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatCallback", ctx, cu, repo, cb)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatCallback indicates an expected call of CreateChatCallback.
func (mr *MockInteractorMockRecorder) CreateChatCallback(ctx, cu, repo, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatCallback", reflect.TypeOf((*MockInteractor)(nil).CreateChatCallback), ctx, cu, repo, cb)
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

// FindChatCallbackByHash mocks base method.
func (m *MockInteractor) FindChatCallbackByHash(ctx context.Context, state string) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatCallbackByHash", ctx, state)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatCallbackByHash indicates an expected call of FindChatCallbackByHash.
func (mr *MockInteractorMockRecorder) FindChatCallbackByHash(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatCallbackByHash", reflect.TypeOf((*MockInteractor)(nil).FindChatCallbackByHash), ctx, state)
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

// Publish mocks base method.
func (m *MockInteractor) Publish(ctx context.Context, typ notification.Type, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, typ, r, d, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockInteractorMockRecorder) Publish(ctx, typ, r, d, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockInteractor)(nil).Publish), ctx, typ, r, d, a)
}

// Rollback mocks base method.
func (m *MockInteractor) Rollback(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx, u, re, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rollback indicates an expected call of Rollback.
func (mr *MockInteractorMockRecorder) Rollback(ctx, u, re, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockInteractor)(nil).Rollback), ctx, u, re, d, env)
}

// SaveChatUser mocks base method.
func (m *MockInteractor) SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveChatUser", ctx, u, cu)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveChatUser indicates an expected call of SaveChatUser.
func (mr *MockInteractorMockRecorder) SaveChatUser(ctx, u, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveChatUser", reflect.TypeOf((*MockInteractor)(nil).SaveChatUser), ctx, u, cu)
}

// Subscribe mocks base method.
func (m *MockInteractor) Subscribe(arg0 func(*ent.User, *ent.Notification)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockInteractorMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockInteractor)(nil).Subscribe), arg0)
}