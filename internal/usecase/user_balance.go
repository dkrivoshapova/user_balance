package usecase

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/vladjong/user_balance/internal/adapters/db"
	"github.com/vladjong/user_balance/internal/entities"
	"github.com/vladjong/user_balance/pkg/fileworker"
)

type userBalanseUseCase struct {
	storage    db.UserBalanse
	fileworker fileworker.FileWorker
}

func New(storage db.UserBalanse, fileworker fileworker.FileWorker) *userBalanseUseCase {
	return &userBalanseUseCase{
		storage:    storage,
		fileworker: fileworker,
	}
}

func (u *userBalanseUseCase) GetCustomerBalance(id int) (user entities.Customer, err error) {
	return u.storage.GetCustomerBalance(id)
}

func (u *userBalanseUseCase) PostCustomerBalance(id int, value decimal.Decimal) error {
	if ok := checkValue(value); !ok {
		return errors.New("error: value is negative")
	}
	customer := entities.Customer{
		Id:      id,
		Balance: value,
	}
	transaction := entities.Transaction{
		CustomeId:           id,
		ServiceID:           4,
		OrderID:             4,
		Cost:                value,
		TransactionDatiTime: time.Now(),
	}
	return u.storage.PostCustomerBalance(customer, transaction)
}

func (u *userBalanseUseCase) PostReserveBalance(customerId, serviceId, orderId int, value decimal.Decimal) error {
	if ok := checkValue(value); !ok {
		return errors.New("error: value is negative")
	}
	transaction := entities.Transaction{
		CustomeId:           customerId,
		ServiceID:           serviceId,
		OrderID:             orderId,
		Cost:                value,
		TransactionDatiTime: time.Now(),
	}
	return u.storage.PostReserveBalance(transaction)
}

func (u *userBalanseUseCase) PostDeReservingBalance(customerId, serviceId, orderId int, value decimal.Decimal, status bool) error {
	if ok := checkValue(value); !ok {
		return errors.New("error: value is negative")
	}
	transaction := entities.Transaction{
		CustomeId: customerId,
		ServiceID: serviceId,
		OrderID:   orderId,
		Cost:      value,
	}
	history := entities.History{
		TransactionId:      customerId,
		StatusTransaction:  status,
		AccountingDatetime: time.Now(),
	}
	return u.storage.PostDeReservingBalance(transaction, history)
}

func (u *userBalanseUseCase) GetHistoryReport(date time.Time) (string, error) {
	report, err := u.storage.GetHistoryReport(date)
	if err != nil {
		return "", nil
	}
	headers := []string{"id", "name", "all_sum"}
	dateStr := date.Format("2006-01")
	return u.fileworker.Record(report, headers, dateStr)
}

func (u *userBalanseUseCase) GetCustomerReport(id int, date time.Time) (report []entities.CustomerReport, err error) {
	return u.storage.GetCustomerReport(id, date)
}

func checkValue(value decimal.Decimal) bool {
	return !value.IsNegative()
}
