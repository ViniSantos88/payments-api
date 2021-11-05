package repositories

import (
	"github.com/ViniSantos88/payments-api/domain"
)

// PaymentsRepo is the data interface for Payments repositories
type PaymentsRepo interface {
	CreateAccounts(accounts domain.Accounts) (accountID int64, err error)
	GetAccount(accountID int64) (accountResponse domain.Accounts, err error)
	CreateTransactions(transaction domain.Transactions) (transactionID int64, err error)
}
