package postgresql

import (
	"github.com/ViniSantos88/payments-api/domain"
)

// CreateAccounts is the service to create bank account
func (repo *PostgreSQLRepoImpl) CreateTransactions(transactions domain.Transactions) (transactionID int64, err error) {

	sql :=
		`INSERT INTO transactions
	        (
						account_id,
						operationType_id,
						amount,
						eventDate
					)
	   VALUES 
        	(
						:account_id,
						:operation_type_id,
						:amount,
						NOW()
					)
	   returning transaction_id`

	rows, err := repo.Db.NamedQuery(sql, transactions)
	if err != nil {
		return 0, domain.ErrInternalServerError
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&transactionID)
		if err != nil {
			return 0, domain.ErrInternalServerError
		}
	}

	return transactionID, nil
}
