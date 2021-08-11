// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/hanjunlee/gitploy/ent"
	vo "github.com/hanjunlee/gitploy/vo"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Activate mocks base method.
func (m *MockStore) Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activate", ctx, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Activate indicates an expected call of Activate.
func (mr *MockStoreMockRecorder) Activate(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockStore)(nil).Activate), ctx, r)
}

// CheckNotificationRecordOfEvent mocks base method.
func (m *MockStore) CheckNotificationRecordOfEvent(ctx context.Context, e *ent.Event) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckNotificationRecordOfEvent", ctx, e)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckNotificationRecordOfEvent indicates an expected call of CheckNotificationRecordOfEvent.
func (mr *MockStoreMockRecorder) CheckNotificationRecordOfEvent(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckNotificationRecordOfEvent", reflect.TypeOf((*MockStore)(nil).CheckNotificationRecordOfEvent), ctx, e)
}

// CloseChatCallback mocks base method.
func (m *MockStore) CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseChatCallback", ctx, cb)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseChatCallback indicates an expected call of CloseChatCallback.
func (mr *MockStoreMockRecorder) CloseChatCallback(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseChatCallback", reflect.TypeOf((*MockStore)(nil).CloseChatCallback), ctx, cb)
}

// CreateApproval mocks base method.
func (m *MockStore) CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApproval", ctx, a)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApproval indicates an expected call of CreateApproval.
func (mr *MockStoreMockRecorder) CreateApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApproval", reflect.TypeOf((*MockStore)(nil).CreateApproval), ctx, a)
}

// CreateChatCallback mocks base method.
func (m *MockStore) CreateChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatCallback", ctx, cu, repo, cb)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatCallback indicates an expected call of CreateChatCallback.
func (mr *MockStoreMockRecorder) CreateChatCallback(ctx, cu, repo, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatCallback", reflect.TypeOf((*MockStore)(nil).CreateChatCallback), ctx, cu, repo, cb)
}

// CreateChatUser mocks base method.
func (m *MockStore) CreateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatUser", ctx, u, cu)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatUser indicates an expected call of CreateChatUser.
func (mr *MockStoreMockRecorder) CreateChatUser(ctx, u, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatUser", reflect.TypeOf((*MockStore)(nil).CreateChatUser), ctx, u, cu)
}

// CreateDeployment mocks base method.
func (m *MockStore) CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", ctx, d)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockStoreMockRecorder) CreateDeployment(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockStore)(nil).CreateDeployment), ctx, d)
}

// CreateDeploymentStatus mocks base method.
func (m *MockStore) CreateDeploymentStatus(ctx context.Context, s *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeploymentStatus", ctx, s)
	ret0, _ := ret[0].(*ent.DeploymentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeploymentStatus indicates an expected call of CreateDeploymentStatus.
func (mr *MockStoreMockRecorder) CreateDeploymentStatus(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeploymentStatus", reflect.TypeOf((*MockStore)(nil).CreateDeploymentStatus), ctx, s)
}

// CreateEvent mocks base method.
func (m *MockStore) CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, e)
	ret0, _ := ret[0].(*ent.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockStoreMockRecorder) CreateEvent(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockStore)(nil).CreateEvent), ctx, e)
}

// CreateNotification mocks base method.
func (m *MockStore) CreateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotification", ctx, n)
	ret0, _ := ret[0].(*ent.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNotification indicates an expected call of CreateNotification.
func (mr *MockStoreMockRecorder) CreateNotification(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotification", reflect.TypeOf((*MockStore)(nil).CreateNotification), ctx, n)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, u)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, u)
}

// Deactivate mocks base method.
func (m *MockStore) Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", ctx, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockStoreMockRecorder) Deactivate(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockStore)(nil).Deactivate), ctx, r)
}

// DeleteApproval mocks base method.
func (m *MockStore) DeleteApproval(ctx context.Context, a *ent.Approval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApproval", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApproval indicates an expected call of DeleteApproval.
func (mr *MockStoreMockRecorder) DeleteApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApproval", reflect.TypeOf((*MockStore)(nil).DeleteApproval), ctx, a)
}

// FindApprovalByID mocks base method.
func (m *MockStore) FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApprovalByID", ctx, id)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApprovalByID indicates an expected call of FindApprovalByID.
func (mr *MockStoreMockRecorder) FindApprovalByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApprovalByID", reflect.TypeOf((*MockStore)(nil).FindApprovalByID), ctx, id)
}

// FindApprovalOfUser mocks base method.
func (m *MockStore) FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApprovalOfUser", ctx, d, u)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApprovalOfUser indicates an expected call of FindApprovalOfUser.
func (mr *MockStoreMockRecorder) FindApprovalOfUser(ctx, d, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApprovalOfUser", reflect.TypeOf((*MockStore)(nil).FindApprovalOfUser), ctx, d, u)
}

// FindChatCallbackByHash mocks base method.
func (m *MockStore) FindChatCallbackByHash(ctx context.Context, state string) (*ent.ChatCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatCallbackByHash", ctx, state)
	ret0, _ := ret[0].(*ent.ChatCallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatCallbackByHash indicates an expected call of FindChatCallbackByHash.
func (mr *MockStoreMockRecorder) FindChatCallbackByHash(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatCallbackByHash", reflect.TypeOf((*MockStore)(nil).FindChatCallbackByHash), ctx, state)
}

// FindChatUserByID mocks base method.
func (m *MockStore) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindChatUserByID", ctx, id)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChatUserByID indicates an expected call of FindChatUserByID.
func (mr *MockStoreMockRecorder) FindChatUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChatUserByID", reflect.TypeOf((*MockStore)(nil).FindChatUserByID), ctx, id)
}

// FindDeploymentByID mocks base method.
func (m *MockStore) FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentByID", ctx, id)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentByID indicates an expected call of FindDeploymentByID.
func (mr *MockStoreMockRecorder) FindDeploymentByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentByID", reflect.TypeOf((*MockStore)(nil).FindDeploymentByID), ctx, id)
}

// FindDeploymentByUID mocks base method.
func (m *MockStore) FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentByUID", ctx, uid)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentByUID indicates an expected call of FindDeploymentByUID.
func (mr *MockStoreMockRecorder) FindDeploymentByUID(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentByUID", reflect.TypeOf((*MockStore)(nil).FindDeploymentByUID), ctx, uid)
}

// FindDeploymentOfRepoByNumber mocks base method.
func (m *MockStore) FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentOfRepoByNumber", ctx, r, number)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentOfRepoByNumber indicates an expected call of FindDeploymentOfRepoByNumber.
func (mr *MockStoreMockRecorder) FindDeploymentOfRepoByNumber(ctx, r, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentOfRepoByNumber", reflect.TypeOf((*MockStore)(nil).FindDeploymentOfRepoByNumber), ctx, r, number)
}

// FindLatestSuccessfulDeployment mocks base method.
func (m *MockStore) FindLatestSuccessfulDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLatestSuccessfulDeployment", ctx, d)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLatestSuccessfulDeployment indicates an expected call of FindLatestSuccessfulDeployment.
func (mr *MockStoreMockRecorder) FindLatestSuccessfulDeployment(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLatestSuccessfulDeployment", reflect.TypeOf((*MockStore)(nil).FindLatestSuccessfulDeployment), ctx, d)
}

// FindNotificationByID mocks base method.
func (m *MockStore) FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNotificationByID", ctx, id)
	ret0, _ := ret[0].(*ent.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNotificationByID indicates an expected call of FindNotificationByID.
func (mr *MockStoreMockRecorder) FindNotificationByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNotificationByID", reflect.TypeOf((*MockStore)(nil).FindNotificationByID), ctx, id)
}

// FindPermOfRepo mocks base method.
func (m *MockStore) FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPermOfRepo", ctx, r, u)
	ret0, _ := ret[0].(*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPermOfRepo indicates an expected call of FindPermOfRepo.
func (mr *MockStoreMockRecorder) FindPermOfRepo(ctx, r, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPermOfRepo", reflect.TypeOf((*MockStore)(nil).FindPermOfRepo), ctx, r, u)
}

// FindRepoOfUserByID mocks base method.
func (m *MockStore) FindRepoOfUserByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByID", ctx, u, id)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByID indicates an expected call of FindRepoOfUserByID.
func (mr *MockStoreMockRecorder) FindRepoOfUserByID(ctx, u, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByID", reflect.TypeOf((*MockStore)(nil).FindRepoOfUserByID), ctx, u, id)
}

// FindRepoOfUserByNamespaceName mocks base method.
func (m *MockStore) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByNamespaceName", ctx, u, namespace, name)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByNamespaceName indicates an expected call of FindRepoOfUserByNamespaceName.
func (mr *MockStoreMockRecorder) FindRepoOfUserByNamespaceName(ctx, u, namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByNamespaceName", reflect.TypeOf((*MockStore)(nil).FindRepoOfUserByNamespaceName), ctx, u, namespace, name)
}

// FindUserByHash mocks base method.
func (m *MockStore) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByHash", ctx, hash)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByHash indicates an expected call of FindUserByHash.
func (mr *MockStoreMockRecorder) FindUserByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByHash", reflect.TypeOf((*MockStore)(nil).FindUserByHash), ctx, hash)
}

// FindUserByID mocks base method.
func (m *MockStore) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, id)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockStoreMockRecorder) FindUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockStore)(nil).FindUserByID), ctx, id)
}

// FindUserByLogin mocks base method.
func (m *MockStore) FindUserByLogin(ctx context.Context, login string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByLogin", ctx, login)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByLogin indicates an expected call of FindUserByLogin.
func (mr *MockStoreMockRecorder) FindUserByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByLogin", reflect.TypeOf((*MockStore)(nil).FindUserByLogin), ctx, login)
}

// GetNextDeploymentNumberOfRepo mocks base method.
func (m *MockStore) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextDeploymentNumberOfRepo", ctx, r)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextDeploymentNumberOfRepo indicates an expected call of GetNextDeploymentNumberOfRepo.
func (mr *MockStoreMockRecorder) GetNextDeploymentNumberOfRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextDeploymentNumberOfRepo", reflect.TypeOf((*MockStore)(nil).GetNextDeploymentNumberOfRepo), ctx, r)
}

// ListApprovals mocks base method.
func (m *MockStore) ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApprovals", ctx, d)
	ret0, _ := ret[0].([]*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApprovals indicates an expected call of ListApprovals.
func (mr *MockStoreMockRecorder) ListApprovals(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApprovals", reflect.TypeOf((*MockStore)(nil).ListApprovals), ctx, d)
}

// ListDeploymentsOfRepo mocks base method.
func (m *MockStore) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env, status string, page, perPage int) ([]*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentsOfRepo", ctx, r, env, status, page, perPage)
	ret0, _ := ret[0].([]*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentsOfRepo indicates an expected call of ListDeploymentsOfRepo.
func (mr *MockStoreMockRecorder) ListDeploymentsOfRepo(ctx, r, env, status, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentsOfRepo", reflect.TypeOf((*MockStore)(nil).ListDeploymentsOfRepo), ctx, r, env, status, page, perPage)
}

// ListEventsGreaterThanTime mocks base method.
func (m *MockStore) ListEventsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventsGreaterThanTime", ctx, t)
	ret0, _ := ret[0].([]*ent.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventsGreaterThanTime indicates an expected call of ListEventsGreaterThanTime.
func (mr *MockStoreMockRecorder) ListEventsGreaterThanTime(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventsGreaterThanTime", reflect.TypeOf((*MockStore)(nil).ListEventsGreaterThanTime), ctx, t)
}

// ListInactiveDeploymentsLessThanTime mocks base method.
func (m *MockStore) ListInactiveDeploymentsLessThanTime(ctx context.Context, t time.Time, page, perPage int) ([]*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInactiveDeploymentsLessThanTime", ctx, t, page, perPage)
	ret0, _ := ret[0].([]*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInactiveDeploymentsLessThanTime indicates an expected call of ListInactiveDeploymentsLessThanTime.
func (mr *MockStoreMockRecorder) ListInactiveDeploymentsLessThanTime(ctx, t, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInactiveDeploymentsLessThanTime", reflect.TypeOf((*MockStore)(nil).ListInactiveDeploymentsLessThanTime), ctx, t, page, perPage)
}

// ListNotifications mocks base method.
func (m *MockStore) ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotifications", ctx, u, page, perPage)
	ret0, _ := ret[0].([]*ent.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNotifications indicates an expected call of ListNotifications.
func (mr *MockStoreMockRecorder) ListNotifications(ctx, u, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotifications", reflect.TypeOf((*MockStore)(nil).ListNotifications), ctx, u, page, perPage)
}

// ListPermsOfRepo mocks base method.
func (m *MockStore) ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPermsOfRepo", ctx, r, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermsOfRepo indicates an expected call of ListPermsOfRepo.
func (mr *MockStoreMockRecorder) ListPermsOfRepo(ctx, r, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermsOfRepo", reflect.TypeOf((*MockStore)(nil).ListPermsOfRepo), ctx, r, q, page, perPage)
}

// ListPublishingNotificaitonsGreaterThanTime mocks base method.
func (m *MockStore) ListPublishingNotificaitonsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublishingNotificaitonsGreaterThanTime", ctx, t)
	ret0, _ := ret[0].([]*ent.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublishingNotificaitonsGreaterThanTime indicates an expected call of ListPublishingNotificaitonsGreaterThanTime.
func (mr *MockStoreMockRecorder) ListPublishingNotificaitonsGreaterThanTime(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishingNotificaitonsGreaterThanTime", reflect.TypeOf((*MockStore)(nil).ListPublishingNotificaitonsGreaterThanTime), ctx, t)
}

// ListReposOfUser mocks base method.
func (m *MockStore) ListReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReposOfUser", ctx, u, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReposOfUser indicates an expected call of ListReposOfUser.
func (mr *MockStoreMockRecorder) ListReposOfUser(ctx, u, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReposOfUser", reflect.TypeOf((*MockStore)(nil).ListReposOfUser), ctx, u, q, page, perPage)
}

// ListSortedReposOfUser mocks base method.
func (m *MockStore) ListSortedReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSortedReposOfUser", ctx, u, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSortedReposOfUser indicates an expected call of ListSortedReposOfUser.
func (mr *MockStoreMockRecorder) ListSortedReposOfUser(ctx, u, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSortedReposOfUser", reflect.TypeOf((*MockStore)(nil).ListSortedReposOfUser), ctx, u, q, page, perPage)
}

// SyncPerm mocks base method.
func (m *MockStore) SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPerm", ctx, user, perm, sync)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncPerm indicates an expected call of SyncPerm.
func (mr *MockStoreMockRecorder) SyncPerm(ctx, user, perm, sync interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPerm", reflect.TypeOf((*MockStore)(nil).SyncPerm), ctx, user, perm, sync)
}

// UpdateApproval mocks base method.
func (m *MockStore) UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApproval", ctx, a)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApproval indicates an expected call of UpdateApproval.
func (mr *MockStoreMockRecorder) UpdateApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApproval", reflect.TypeOf((*MockStore)(nil).UpdateApproval), ctx, a)
}

// UpdateChatUser mocks base method.
func (m *MockStore) UpdateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChatUser", ctx, u, cu)
	ret0, _ := ret[0].(*ent.ChatUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChatUser indicates an expected call of UpdateChatUser.
func (mr *MockStoreMockRecorder) UpdateChatUser(ctx, u, cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChatUser", reflect.TypeOf((*MockStore)(nil).UpdateChatUser), ctx, u, cu)
}

// UpdateDeployment mocks base method.
func (m *MockStore) UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployment", ctx, d)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockStoreMockRecorder) UpdateDeployment(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockStore)(nil).UpdateDeployment), ctx, d)
}

// UpdateNotification mocks base method.
func (m *MockStore) UpdateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotification", ctx, n)
	ret0, _ := ret[0].(*ent.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNotification indicates an expected call of UpdateNotification.
func (mr *MockStoreMockRecorder) UpdateNotification(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotification", reflect.TypeOf((*MockStore)(nil).UpdateNotification), ctx, n)
}

// UpdateRepo mocks base method.
func (m *MockStore) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRepo", ctx, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRepo indicates an expected call of UpdateRepo.
func (mr *MockStoreMockRecorder) UpdateRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRepo", reflect.TypeOf((*MockStore)(nil).UpdateRepo), ctx, r)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, u)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), ctx, u)
}

// MockSCM is a mock of SCM interface.
type MockSCM struct {
	ctrl     *gomock.Controller
	recorder *MockSCMMockRecorder
}

// MockSCMMockRecorder is the mock recorder for MockSCM.
type MockSCMMockRecorder struct {
	mock *MockSCM
}

// NewMockSCM creates a new mock instance.
func NewMockSCM(ctrl *gomock.Controller) *MockSCM {
	mock := &MockSCM{ctrl: ctrl}
	mock.recorder = &MockSCMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSCM) EXPECT() *MockSCMMockRecorder {
	return m.recorder
}

// CancelDeployment mocks base method.
func (m *MockSCM) CancelDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelDeployment", ctx, u, r, d, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelDeployment indicates an expected call of CancelDeployment.
func (mr *MockSCMMockRecorder) CancelDeployment(ctx, u, r, d, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelDeployment", reflect.TypeOf((*MockSCM)(nil).CancelDeployment), ctx, u, r, d, s)
}

// CompareCommits mocks base method.
func (m *MockSCM) CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareCommits", ctx, u, r, base, head, page, perPage)
	ret0, _ := ret[0].([]*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareCommits indicates an expected call of CompareCommits.
func (mr *MockSCMMockRecorder) CompareCommits(ctx, u, r, base, head, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareCommits", reflect.TypeOf((*MockSCM)(nil).CompareCommits), ctx, u, r, base, head, page, perPage)
}

// CreateDeployment mocks base method.
func (m *MockSCM) CreateDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*vo.RemoteDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", ctx, u, r, d, e)
	ret0, _ := ret[0].(*vo.RemoteDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockSCMMockRecorder) CreateDeployment(ctx, u, r, d, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockSCM)(nil).CreateDeployment), ctx, u, r, d, e)
}

// CreateWebhook mocks base method.
func (m *MockSCM) CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWebhook", ctx, u, r, c)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWebhook indicates an expected call of CreateWebhook.
func (mr *MockSCMMockRecorder) CreateWebhook(ctx, u, r, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWebhook", reflect.TypeOf((*MockSCM)(nil).CreateWebhook), ctx, u, r, c)
}

// DeleteWebhook mocks base method.
func (m *MockSCM) DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWebhook", ctx, u, r, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWebhook indicates an expected call of DeleteWebhook.
func (mr *MockSCMMockRecorder) DeleteWebhook(ctx, u, r, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebhook", reflect.TypeOf((*MockSCM)(nil).DeleteWebhook), ctx, u, r, id)
}

// GetAllPermsWithRepo mocks base method.
func (m *MockSCM) GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPermsWithRepo", ctx, token)
	ret0, _ := ret[0].([]*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPermsWithRepo indicates an expected call of GetAllPermsWithRepo.
func (mr *MockSCMMockRecorder) GetAllPermsWithRepo(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPermsWithRepo", reflect.TypeOf((*MockSCM)(nil).GetAllPermsWithRepo), ctx, token)
}

// GetBranch mocks base method.
func (m *MockSCM) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, u, r, branch)
	ret0, _ := ret[0].(*vo.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockSCMMockRecorder) GetBranch(ctx, u, r, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockSCM)(nil).GetBranch), ctx, u, r, branch)
}

// GetCommit mocks base method.
func (m *MockSCM) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, u, r, sha)
	ret0, _ := ret[0].(*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockSCMMockRecorder) GetCommit(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockSCM)(nil).GetCommit), ctx, u, r, sha)
}

// GetConfig mocks base method.
func (m *MockSCM) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, u, r)
	ret0, _ := ret[0].(*vo.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockSCMMockRecorder) GetConfig(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockSCM)(nil).GetConfig), ctx, u, r)
}

// GetTag mocks base method.
func (m *MockSCM) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, u, r, tag)
	ret0, _ := ret[0].(*vo.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockSCMMockRecorder) GetTag(ctx, u, r, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockSCM)(nil).GetTag), ctx, u, r, tag)
}

// GetUser mocks base method.
func (m *MockSCM) GetUser(ctx context.Context, token string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, token)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockSCMMockRecorder) GetUser(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockSCM)(nil).GetUser), ctx, token)
}

// ListBranches mocks base method.
func (m *MockSCM) ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", ctx, u, r, page, perPage)
	ret0, _ := ret[0].([]*vo.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockSCMMockRecorder) ListBranches(ctx, u, r, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockSCM)(nil).ListBranches), ctx, u, r, page, perPage)
}

// ListCommitStatuses mocks base method.
func (m *MockSCM) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommitStatuses", ctx, u, r, sha)
	ret0, _ := ret[0].([]*vo.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommitStatuses indicates an expected call of ListCommitStatuses.
func (mr *MockSCMMockRecorder) ListCommitStatuses(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommitStatuses", reflect.TypeOf((*MockSCM)(nil).ListCommitStatuses), ctx, u, r, sha)
}

// ListCommits mocks base method.
func (m *MockSCM) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", ctx, u, r, branch, page, perPage)
	ret0, _ := ret[0].([]*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockSCMMockRecorder) ListCommits(ctx, u, r, branch, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockSCM)(nil).ListCommits), ctx, u, r, branch, page, perPage)
}

// ListTags mocks base method.
func (m *MockSCM) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, u, r, page, perPage)
	ret0, _ := ret[0].([]*vo.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockSCMMockRecorder) ListTags(ctx, u, r, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockSCM)(nil).ListTags), ctx, u, r, page, perPage)
}
