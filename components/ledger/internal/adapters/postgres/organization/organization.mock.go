// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/organization (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=organization.mock.go --package=organization . Repository
//

// Package organization is a generated GoMock package.
package organization

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	http "github.com/LerianStudio/midaz/pkg/net/http"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, organization *mmodel.Organization) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, organization any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, organization)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockRepository) Find(ctx context.Context, id uuid.UUID) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), ctx, id)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context, filter http.Pagination) ([]*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, filter)
	ret0, _ := ret[0].([]*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx, filter)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(ctx context.Context, ids []uuid.UUID) ([]*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", ctx, ids)
	ret0, _ := ret[0].([]*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), ctx, ids)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, id uuid.UUID, organization *mmodel.Organization) (*mmodel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, organization)
	ret0, _ := ret[0].(*mmodel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, id, organization any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, id, organization)
}
