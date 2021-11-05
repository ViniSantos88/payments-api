package postgresql

import (
	"github.com/ViniSantos88/payments-api/domain"
)

// CreateAccounts is the service to get bank account
func (repo *PostgreSQLRepoImpl) GetAccount(accountID int64) (account domain.Accounts, err error) {

	request := &domain.Accounts{
		AccountID: accountID,
	}

	sql :=
		`SELECT  account_id,
		         document_number
	   FROM  accounts
     WHERE account_id =:account_id`

	rows, err := repo.Db.NamedQuery(sql, request)
	if err != nil {
		return domain.Accounts{}, domain.ErrInternalServerError
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&account)
		if err != nil {
			return domain.Accounts{}, domain.ErrInternalServerError
		}
	}

	if rows.Err() != nil {
		return domain.Accounts{}, domain.ErrInternalServerError
	}

	if account == (domain.Accounts{}) {
		return domain.Accounts{}, domain.ErrAccountNotFound
	}

	return account, nil
}
