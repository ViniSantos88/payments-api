package services

import (
	"github.com/ViniSantos88/payments-api/domain"
)

//CreateAccounts is the transaction creation service
func (ps *PaymentsServiceImpl) CreateTransactions(transaction domain.Transactions) (ID int64, err error) {

	ID, err = ps.PaymentsRepo.CreateTransactions(transaction)
	if err != nil {
		return 0, err
	}

	return ID, nil
}
