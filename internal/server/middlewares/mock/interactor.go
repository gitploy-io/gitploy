// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/hanjunlee/gitploy/ent"
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

// FindUserByHash mocks base method.
func (m *MockInteractor) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByHash", ctx, hash)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByHash indicates an expected call of FindUserByHash.
func (mr *MockInteractorMockRecorder) FindUserByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByHash", reflect.TypeOf((*MockInteractor)(nil).FindUserByHash), ctx, hash)
}

// GetLicense mocks base method.
func (m *MockInteractor) GetLicense(ctx context.Context) (*vo.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLicense", ctx)
	ret0, _ := ret[0].(*vo.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLicense indicates an expected call of GetLicense.
func (mr *MockInteractorMockRecorder) GetLicense(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLicense", reflect.TypeOf((*MockInteractor)(nil).GetLicense), ctx)
}
