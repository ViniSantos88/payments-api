package services

import (
	"github.com/ViniSantos88/payments-api/domain"
)

//CreateAccounts is the accounts creation service
func (ps *PaymentsServiceImpl) CreateAccounts(accounts domain.Accounts) (ID int64, err error) {

	ID, err = ps.PaymentsRepo.CreateAccounts(accounts)
	if err != nil {
		return 0, err
	}

	return ID, nil
}
