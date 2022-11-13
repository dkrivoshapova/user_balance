// Code generated by MockGen. DO NOT EDIT.
// Source: user_balance_interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
	entities "github.com/vladjong/user_balance/internal/entities"
)

// MockUserBalanse is a mock of UserBalanse interface.
type MockUserBalanse struct {
	ctrl     *gomock.Controller
	recorder *MockUserBalanseMockRecorder
}

// MockUserBalanseMockRecorder is the mock recorder for MockUserBalanse.
type MockUserBalanseMockRecorder struct {
	mock *MockUserBalanse
}

// NewMockUserBalanse creates a new mock instance.
func NewMockUserBalanse(ctrl *gomock.Controller) *MockUserBalanse {
	mock := &MockUserBalanse{ctrl: ctrl}
	mock.recorder = &MockUserBalanseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserBalanse) EXPECT() *MockUserBalanseMockRecorder {
	return m.recorder
}

// GetCustomerBalance mocks base method.
func (m *MockUserBalanse) GetCustomerBalance(id int) (entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerBalance", id)
	ret0, _ := ret[0].(entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerBalance indicates an expected call of GetCustomerBalance.
func (mr *MockUserBalanseMockRecorder) GetCustomerBalance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerBalance", reflect.TypeOf((*MockUserBalanse)(nil).GetCustomerBalance), id)
}

// GetCustomerReport mocks base method.
func (m *MockUserBalanse) GetCustomerReport(id int, date time.Time) ([]entities.CustomerReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerReport", id, date)
	ret0, _ := ret[0].([]entities.CustomerReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerReport indicates an expected call of GetCustomerReport.
func (mr *MockUserBalanseMockRecorder) GetCustomerReport(id, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerReport", reflect.TypeOf((*MockUserBalanse)(nil).GetCustomerReport), id, date)
}

// GetHistoryReport mocks base method.
func (m *MockUserBalanse) GetHistoryReport(date time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryReport", date)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistoryReport indicates an expected call of GetHistoryReport.
func (mr *MockUserBalanseMockRecorder) GetHistoryReport(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryReport", reflect.TypeOf((*MockUserBalanse)(nil).GetHistoryReport), date)
}

// PostCustomerBalance mocks base method.
func (m *MockUserBalanse) PostCustomerBalance(id int, value decimal.Decimal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCustomerBalance", id, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCustomerBalance indicates an expected call of PostCustomerBalance.
func (mr *MockUserBalanseMockRecorder) PostCustomerBalance(id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCustomerBalance", reflect.TypeOf((*MockUserBalanse)(nil).PostCustomerBalance), id, value)
}

// PostDeReservingBalance mocks base method.
func (m *MockUserBalanse) PostDeReservingBalance(customerId, serviceId, orderId int, value decimal.Decimal, status bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDeReservingBalance", customerId, serviceId, orderId, value, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostDeReservingBalance indicates an expected call of PostDeReservingBalance.
func (mr *MockUserBalanseMockRecorder) PostDeReservingBalance(customerId, serviceId, orderId, value, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDeReservingBalance", reflect.TypeOf((*MockUserBalanse)(nil).PostDeReservingBalance), customerId, serviceId, orderId, value, status)
}

// PostReserveBalance mocks base method.
func (m *MockUserBalanse) PostReserveBalance(customerId, serviceId, orderId int, value decimal.Decimal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostReserveBalance", customerId, serviceId, orderId, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostReserveBalance indicates an expected call of PostReserveBalance.
func (mr *MockUserBalanseMockRecorder) PostReserveBalance(customerId, serviceId, orderId, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostReserveBalance", reflect.TypeOf((*MockUserBalanse)(nil).PostReserveBalance), customerId, serviceId, orderId, value)
}
