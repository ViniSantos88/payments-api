package services

import (
	"github.com/ViniSantos88/payments-api/domain"
)

//GetAccount is the get account detail
func (ps *PaymentsServiceImpl) GetAccount(accountID int64) (*domain.Accounts, error) {

	response, err := ps.PaymentsRepo.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
