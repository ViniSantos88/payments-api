package services

import (
	repo "github.com/ViniSantos88/payments-api/application/repositories"
	"github.com/ViniSantos88/payments-api/application/repositories/postgresql"
	"github.com/ViniSantos88/payments-api/domain"
)

var paymentsServiceImpl *PaymentsServiceImpl

// PaymentsService is the interface for Payments services
type PaymentsService interface {
	CreateAccounts(account domain.Accounts) (accountID int64, err error)
	GetAccount(accountID int64) (*domain.Accounts, error)
	CreateTransactions(transaction domain.Transactions) (transactionID int64, err error)
}

// PaymentsServiceImpl is the implementation for PaymentsService
type PaymentsServiceImpl struct {
	PaymentsRepo repo.PaymentsRepo
}

// GetPaymentsServiceImpl returns an instance of PaymentsServiceImpl
func GetPaymentsServiceImpl() *PaymentsServiceImpl {
	if paymentsServiceImpl == nil {
		paymentsServiceImpl = NewPaymentsServiceImpl(postgresql.GetPostgreSQLRepoImpl())

	}
	return paymentsServiceImpl
}

// NewPaymentsServiceImpl is a dependency injection an instance of PaymentsServiceImpl
func NewPaymentsServiceImpl(paymentsRepo repo.PaymentsRepo) *PaymentsServiceImpl {
	return &PaymentsServiceImpl{
		PaymentsRepo: paymentsRepo,
	}
}
