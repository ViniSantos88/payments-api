package postgresql

import (
	"github.com/ViniSantos88/payments-api/domain"
)

// CreateAccounts is the service to create bank account
func (repo *PostgreSQLRepoImpl) CreateAccounts(accounts domain.Accounts) (accountID int64, err error) {

	sql :=
		`INSERT INTO accounts
	        (document_number)
	   VALUES 
        	(:document_number)
	   returning account_id`

	rows, err := repo.Db.NamedQuery(sql, accounts)
	if err != nil {
		return 0, domain.ErrInternalServerError
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&accountID)
		if err != nil {
			return 0, domain.ErrInternalServerError
		}
	}

	return accountID, nil
}
