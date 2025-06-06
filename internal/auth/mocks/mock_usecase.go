// Code generated by MockGen. DO NOT EDIT.
// Source: internal/auth/domain/model.go
//
// Generated by this command:
//
//	mockgen -source=internal/auth/domain/model.go -destination=internal/auth/mocks/mock_usecase.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/asruldev/cab/internal/auth/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
	isgomock struct{}
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthRepository) CreateUser(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthRepositoryMockRecorder) CreateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthRepository)(nil).CreateUser), user)
}

// FindByEmail mocks base method.
func (m *MockAuthRepository) FindByEmail(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockAuthRepositoryMockRecorder) FindByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockAuthRepository)(nil).FindByEmail), email)
}

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
	isgomock struct{}
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthUsecase) Login(email, password string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockAuthUsecaseMockRecorder) Login(email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUsecase)(nil).Login), email, password)
}

// Register mocks base method.
func (m *MockAuthUsecase) Register(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthUsecaseMockRecorder) Register(email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthUsecase)(nil).Register), email, password)
}
