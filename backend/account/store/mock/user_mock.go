// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source user.go -destination ./mock/user_mock.go -package mock UserDatabase
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	models "github.com/dark-vinci/wapp/backend/sdk/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserDatabase is a mock of UserDatabase interface.
type MockUserDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockUserDatabaseMockRecorder
}

// MockUserDatabaseMockRecorder is the mock recorder for MockUserDatabase.
type MockUserDatabaseMockRecorder struct {
	mock *MockUserDatabase
}

// NewMockUserDatabase creates a new mock instance.
func NewMockUserDatabase(ctrl *gomock.Controller) *MockUserDatabase {
	mock := &MockUserDatabase{ctrl: ctrl}
	mock.recorder = &MockUserDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDatabase) EXPECT() *MockUserDatabaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDatabase) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDatabaseMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDatabase)(nil).CreateUser), ctx, user)
}

// Delete mocks base method.
func (m *MockUserDatabase) Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, deletedAt, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserDatabaseMockRecorder) Delete(ctx, id, deletedAt, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDatabase)(nil).Delete), ctx, id, deletedAt, tx)
}

// GetUserByID mocks base method.
func (m *MockUserDatabase) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserDatabaseMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserDatabase)(nil).GetUserByID), ctx, id)
}
