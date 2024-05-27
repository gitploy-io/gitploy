// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	interactor "github.com/gitploy-io/gitploy/internal/interactor"
	ent "github.com/gitploy-io/gitploy/model/ent"
	extent "github.com/gitploy-io/gitploy/model/extent"
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

// ActivateRepo mocks base method.
func (m *MockInteractor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateRepo", ctx, u, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateRepo indicates an expected call of ActivateRepo.
func (mr *MockInteractorMockRecorder) ActivateRepo(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateRepo", reflect.TypeOf((*MockInteractor)(nil).ActivateRepo), ctx, u, r)
}

// CompareCommits mocks base method.
func (m *MockInteractor) CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, opt *interactor.ListOptions) ([]*extent.Commit, []*extent.CommitFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareCommits", ctx, u, r, base, head, opt)
	ret0, _ := ret[0].([]*extent.Commit)
	ret1, _ := ret[1].([]*extent.CommitFile)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CompareCommits indicates an expected call of CompareCommits.
func (mr *MockInteractorMockRecorder) CompareCommits(ctx, u, r, base, head, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareCommits", reflect.TypeOf((*MockInteractor)(nil).CompareCommits), ctx, u, r, base, head, opt)
}

// CompareCommitsFromLastestDeployment mocks base method.
func (m *MockInteractor) CompareCommitsFromLastestDeployment(ctx context.Context, r *ent.Repo, d *ent.Deployment, options *interactor.ListOptions) ([]*extent.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareCommitsFromLastestDeployment", ctx, r, d, options)
	ret0, _ := ret[0].([]*extent.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareCommitsFromLastestDeployment indicates an expected call of CompareCommitsFromLastestDeployment.
func (mr *MockInteractorMockRecorder) CompareCommitsFromLastestDeployment(ctx, r, d, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareCommitsFromLastestDeployment", reflect.TypeOf((*MockInteractor)(nil).CompareCommitsFromLastestDeployment), ctx, r, d, options)
}

// CreateLock mocks base method.
func (m *MockInteractor) CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLock", ctx, l)
	ret0, _ := ret[0].(*ent.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLock indicates an expected call of CreateLock.
func (mr *MockInteractorMockRecorder) CreateLock(ctx, l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLock", reflect.TypeOf((*MockInteractor)(nil).CreateLock), ctx, l)
}

// CreateRemoteDeploymentStatus mocks base method.
func (m *MockInteractor) CreateRemoteDeploymentStatus(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, ds *extent.RemoteDeploymentStatus) (*extent.RemoteDeploymentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRemoteDeploymentStatus", ctx, u, r, d, ds)
	ret0, _ := ret[0].(*extent.RemoteDeploymentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRemoteDeploymentStatus indicates an expected call of CreateRemoteDeploymentStatus.
func (mr *MockInteractorMockRecorder) CreateRemoteDeploymentStatus(ctx, u, r, d, ds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRemoteDeploymentStatus", reflect.TypeOf((*MockInteractor)(nil).CreateRemoteDeploymentStatus), ctx, u, r, d, ds)
}

// DeactivateRepo mocks base method.
func (m *MockInteractor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateRepo", ctx, u, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateRepo indicates an expected call of DeactivateRepo.
func (mr *MockInteractorMockRecorder) DeactivateRepo(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateRepo", reflect.TypeOf((*MockInteractor)(nil).DeactivateRepo), ctx, u, r)
}

// DeleteLock mocks base method.
func (m *MockInteractor) DeleteLock(ctx context.Context, l *ent.Lock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLock", ctx, l)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLock indicates an expected call of DeleteLock.
func (mr *MockInteractorMockRecorder) DeleteLock(ctx, l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLock", reflect.TypeOf((*MockInteractor)(nil).DeleteLock), ctx, l)
}

// Deploy mocks base method.
func (m *MockInteractor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error) {
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

// DeployToRemote mocks base method.
func (m *MockInteractor) DeployToRemote(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployToRemote", ctx, u, r, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployToRemote indicates an expected call of DeployToRemote.
func (mr *MockInteractorMockRecorder) DeployToRemote(ctx, u, r, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployToRemote", reflect.TypeOf((*MockInteractor)(nil).DeployToRemote), ctx, u, r, d, env)
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

// FindDeploymentOfRepoByNumber mocks base method.
func (m *MockInteractor) FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentOfRepoByNumber", ctx, r, number)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentOfRepoByNumber indicates an expected call of FindDeploymentOfRepoByNumber.
func (mr *MockInteractorMockRecorder) FindDeploymentOfRepoByNumber(ctx, r, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentOfRepoByNumber", reflect.TypeOf((*MockInteractor)(nil).FindDeploymentOfRepoByNumber), ctx, r, number)
}

// FindLockByID mocks base method.
func (m *MockInteractor) FindLockByID(ctx context.Context, id int) (*ent.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLockByID", ctx, id)
	ret0, _ := ret[0].(*ent.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLockByID indicates an expected call of FindLockByID.
func (mr *MockInteractorMockRecorder) FindLockByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLockByID", reflect.TypeOf((*MockInteractor)(nil).FindLockByID), ctx, id)
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

// FindPrevSuccessDeployment mocks base method.
func (m *MockInteractor) FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrevSuccessDeployment", ctx, d)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrevSuccessDeployment indicates an expected call of FindPrevSuccessDeployment.
func (mr *MockInteractorMockRecorder) FindPrevSuccessDeployment(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrevSuccessDeployment", reflect.TypeOf((*MockInteractor)(nil).FindPrevSuccessDeployment), ctx, d)
}

// FindRepoOfUserByID mocks base method.
func (m *MockInteractor) FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByID", ctx, u, id)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByID indicates an expected call of FindRepoOfUserByID.
func (mr *MockInteractorMockRecorder) FindRepoOfUserByID(ctx, u, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByID", reflect.TypeOf((*MockInteractor)(nil).FindRepoOfUserByID), ctx, u, id)
}

// FindRepoOfUserByNamespaceName mocks base method.
func (m *MockInteractor) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, opt *interactor.FindRepoOfUserByNamespaceNameOptions) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByNamespaceName", ctx, u, opt)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByNamespaceName indicates an expected call of FindRepoOfUserByNamespaceName.
func (mr *MockInteractorMockRecorder) FindRepoOfUserByNamespaceName(ctx, u, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByNamespaceName", reflect.TypeOf((*MockInteractor)(nil).FindRepoOfUserByNamespaceName), ctx, u, opt)
}

// FindReviewOfUser mocks base method.
func (m *MockInteractor) FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindReviewOfUser", ctx, u, d)
	ret0, _ := ret[0].(*ent.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReviewOfUser indicates an expected call of FindReviewOfUser.
func (mr *MockInteractorMockRecorder) FindReviewOfUser(ctx, u, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReviewOfUser", reflect.TypeOf((*MockInteractor)(nil).FindReviewOfUser), ctx, u, d)
}

// FindUserByID mocks base method.
func (m *MockInteractor) FindUserByID(ctx context.Context, id int64) (*ent.User, error) {
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
func (m *MockInteractor) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*extent.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, u, r, branch)
	ret0, _ := ret[0].(*extent.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockInteractorMockRecorder) GetBranch(ctx, u, r, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockInteractor)(nil).GetBranch), ctx, u, r, branch)
}

// GetCommit mocks base method.
func (m *MockInteractor) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*extent.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, u, r, sha)
	ret0, _ := ret[0].(*extent.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockInteractorMockRecorder) GetCommit(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockInteractor)(nil).GetCommit), ctx, u, r, sha)
}

// GetConfig mocks base method.
func (m *MockInteractor) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, u, r)
	ret0, _ := ret[0].(*extent.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockInteractorMockRecorder) GetConfig(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockInteractor)(nil).GetConfig), ctx, u, r)
}

// GetDefaultBranch mocks base method.
func (m *MockInteractor) GetDefaultBranch(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultBranch", ctx, u, r)
	ret0, _ := ret[0].(*extent.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultBranch indicates an expected call of GetDefaultBranch.
func (mr *MockInteractorMockRecorder) GetDefaultBranch(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultBranch", reflect.TypeOf((*MockInteractor)(nil).GetDefaultBranch), ctx, u, r)
}

// GetEvaluatedConfig mocks base method.
func (m *MockInteractor) GetEvaluatedConfig(ctx context.Context, u *ent.User, r *ent.Repo, v *extent.EvalValues) (*extent.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvaluatedConfig", ctx, u, r, v)
	ret0, _ := ret[0].(*extent.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvaluatedConfig indicates an expected call of GetEvaluatedConfig.
func (mr *MockInteractorMockRecorder) GetEvaluatedConfig(ctx, u, r, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvaluatedConfig", reflect.TypeOf((*MockInteractor)(nil).GetEvaluatedConfig), ctx, u, r, v)
}

// GetTag mocks base method.
func (m *MockInteractor) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*extent.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, u, r, tag)
	ret0, _ := ret[0].(*extent.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockInteractorMockRecorder) GetTag(ctx, u, r, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockInteractor)(nil).GetTag), ctx, u, r, tag)
}

// HasLockOfRepoForEnv mocks base method.
func (m *MockInteractor) HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasLockOfRepoForEnv", ctx, r, env)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasLockOfRepoForEnv indicates an expected call of HasLockOfRepoForEnv.
func (mr *MockInteractorMockRecorder) HasLockOfRepoForEnv(ctx, r, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLockOfRepoForEnv", reflect.TypeOf((*MockInteractor)(nil).HasLockOfRepoForEnv), ctx, r, env)
}

// ListBranches mocks base method.
func (m *MockInteractor) ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, opt *interactor.ListOptions) ([]*extent.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", ctx, u, r, opt)
	ret0, _ := ret[0].([]*extent.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockInteractorMockRecorder) ListBranches(ctx, u, r, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockInteractor)(nil).ListBranches), ctx, u, r, opt)
}

// ListCommitStatuses mocks base method.
func (m *MockInteractor) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*extent.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommitStatuses", ctx, u, r, sha)
	ret0, _ := ret[0].([]*extent.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommitStatuses indicates an expected call of ListCommitStatuses.
func (mr *MockInteractorMockRecorder) ListCommitStatuses(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommitStatuses", reflect.TypeOf((*MockInteractor)(nil).ListCommitStatuses), ctx, u, r, sha)
}

// ListCommits mocks base method.
func (m *MockInteractor) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, opt *interactor.ListOptions) ([]*extent.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", ctx, u, r, branch, opt)
	ret0, _ := ret[0].([]*extent.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockInteractorMockRecorder) ListCommits(ctx, u, r, branch, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockInteractor)(nil).ListCommits), ctx, u, r, branch, opt)
}

// ListDeploymentStatuses mocks base method.
func (m *MockInteractor) ListDeploymentStatuses(ctx context.Context, d *ent.Deployment) ([]*ent.DeploymentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentStatuses", ctx, d)
	ret0, _ := ret[0].([]*ent.DeploymentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentStatuses indicates an expected call of ListDeploymentStatuses.
func (mr *MockInteractorMockRecorder) ListDeploymentStatuses(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentStatuses", reflect.TypeOf((*MockInteractor)(nil).ListDeploymentStatuses), ctx, d)
}

// ListDeploymentsOfRepo mocks base method.
func (m *MockInteractor) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, opt *interactor.ListDeploymentsOfRepoOptions) ([]*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentsOfRepo", ctx, r, opt)
	ret0, _ := ret[0].([]*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentsOfRepo indicates an expected call of ListDeploymentsOfRepo.
func (mr *MockInteractorMockRecorder) ListDeploymentsOfRepo(ctx, r, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListDeploymentsOfRepo), ctx, r, opt)
}

// ListLocksOfRepo mocks base method.
func (m *MockInteractor) ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocksOfRepo", ctx, r)
	ret0, _ := ret[0].([]*ent.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocksOfRepo indicates an expected call of ListLocksOfRepo.
func (mr *MockInteractorMockRecorder) ListLocksOfRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocksOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListLocksOfRepo), ctx, r)
}

// ListPermsOfRepo mocks base method.
func (m *MockInteractor) ListPermsOfRepo(ctx context.Context, r *ent.Repo, opt *interactor.ListPermsOfRepoOptions) ([]*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPermsOfRepo", ctx, r, opt)
	ret0, _ := ret[0].([]*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermsOfRepo indicates an expected call of ListPermsOfRepo.
func (mr *MockInteractorMockRecorder) ListPermsOfRepo(ctx, r, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListPermsOfRepo), ctx, r, opt)
}

// ListReposOfUser mocks base method.
func (m *MockInteractor) ListReposOfUser(ctx context.Context, u *ent.User, opt *interactor.ListReposOfUserOptions) ([]*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReposOfUser", ctx, u, opt)
	ret0, _ := ret[0].([]*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReposOfUser indicates an expected call of ListReposOfUser.
func (mr *MockInteractorMockRecorder) ListReposOfUser(ctx, u, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReposOfUser", reflect.TypeOf((*MockInteractor)(nil).ListReposOfUser), ctx, u, opt)
}

// ListReviews mocks base method.
func (m *MockInteractor) ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReviews", ctx, d)
	ret0, _ := ret[0].([]*ent.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReviews indicates an expected call of ListReviews.
func (mr *MockInteractorMockRecorder) ListReviews(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReviews", reflect.TypeOf((*MockInteractor)(nil).ListReviews), ctx, d)
}

// ListTags mocks base method.
func (m *MockInteractor) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, opt *interactor.ListOptions) ([]*extent.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, u, r, opt)
	ret0, _ := ret[0].([]*extent.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockInteractorMockRecorder) ListTags(ctx, u, r, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockInteractor)(nil).ListTags), ctx, u, r, opt)
}

// RespondReview mocks base method.
func (m *MockInteractor) RespondReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondReview", ctx, rv)
	ret0, _ := ret[0].(*ent.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondReview indicates an expected call of RespondReview.
func (mr *MockInteractorMockRecorder) RespondReview(ctx, rv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondReview", reflect.TypeOf((*MockInteractor)(nil).RespondReview), ctx, rv)
}

// UpdateLock mocks base method.
func (m *MockInteractor) UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLock", ctx, l)
	ret0, _ := ret[0].(*ent.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLock indicates an expected call of UpdateLock.
func (mr *MockInteractorMockRecorder) UpdateLock(ctx, l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLock", reflect.TypeOf((*MockInteractor)(nil).UpdateLock), ctx, l)
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

// UpdateReview mocks base method.
func (m *MockInteractor) UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", ctx, rv)
	ret0, _ := ret[0].(*ent.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockInteractorMockRecorder) UpdateReview(ctx, rv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockInteractor)(nil).UpdateReview), ctx, rv)
}
