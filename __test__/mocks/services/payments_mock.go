// Code generated by MockGen. DO NOT EDIT.
// Source: application/services/payments.go

// Package servicesmocks is a generated GoMock package.
package servicesmocks

import (
	reflect "reflect"

	domain "github.com/ViniSantos88/payments-api/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockPaymentsService is a mock of PaymentsService interface.
type MockPaymentsService struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentsServiceMockRecorder
}

// MockPaymentsServiceMockRecorder is the mock recorder for MockPaymentsService.
type MockPaymentsServiceMockRecorder struct {
	mock *MockPaymentsService
}

// NewMockPaymentsService creates a new mock instance.
func NewMockPaymentsService(ctrl *gomock.Controller) *MockPaymentsService {
	mock := &MockPaymentsService{ctrl: ctrl}
	mock.recorder = &MockPaymentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentsService) EXPECT() *MockPaymentsServiceMockRecorder {
	return m.recorder
}

// CreateAccounts mocks base method.
func (m *MockPaymentsService) CreateAccounts(account domain.Accounts) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccounts", account)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccounts indicates an expected call of CreateAccounts.
func (mr *MockPaymentsServiceMockRecorder) CreateAccounts(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccounts", reflect.TypeOf((*MockPaymentsService)(nil).CreateAccounts), account)
}

// CreateTransactions mocks base method.
func (m *MockPaymentsService) CreateTransactions(transaction domain.Transactions) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransactions", transaction)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransactions indicates an expected call of CreateTransactions.
func (mr *MockPaymentsServiceMockRecorder) CreateTransactions(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransactions", reflect.TypeOf((*MockPaymentsService)(nil).CreateTransactions), transaction)
}

// GetAccount mocks base method.
func (m *MockPaymentsService) GetAccount(accountID int64) (*domain.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", accountID)
	ret0, _ := ret[0].(*domain.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockPaymentsServiceMockRecorder) GetAccount(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockPaymentsService)(nil).GetAccount), accountID)
}
