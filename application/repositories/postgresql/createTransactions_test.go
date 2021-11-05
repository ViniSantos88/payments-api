package postgresql_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ViniSantos88/payments-api/__test__/testutil"
	"github.com/ViniSantos88/payments-api/application/repositories/postgresql"
	"github.com/ViniSantos88/payments-api/domain"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

/**********************************************************************/
/*                        CreateTransactions                          */
/**********************************************************************/
func TestStart_CreateTransactions(t *testing.T) {
	groupTest := "CreateTransactions"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestCreateTransactionsFunction(t *testing.T) {
	//init sqlMock
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	transaction := domain.Transactions{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          122.90,
	}

	pr := postgresql.NewPostgreSQLRepoImpl(sqlxDB)

	t.Run("Test create transactions ok", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"transaction_id"}).AddRow(123456)

		mock.ExpectQuery("INSERT INTO transactions").WillReturnRows(rows)

		transactionID, _ := pr.CreateTransactions(transaction)

		require.NotNil(t, transactionID, "Record recording error")
	})

	t.Run("Test error database", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		mock.ExpectQuery("INSERT INTO transactions").WillReturnError(fmt.Errorf("error"))
		_, err := pr.CreateTransactions(transaction)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})

	t.Run("Test error mapping data", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		rows := sqlmock.NewRows([]string{"transaction_id"}).AddRow("8d9fs")

		mock.ExpectQuery("INSERT INTO transactions").WillReturnRows(rows)

		_, err := pr.CreateTransactions(transaction)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})
}
func TestEnd_CreateTransactions(t *testing.T) {
	groupTest := "CreateTransactions"
	testutil.PrintTestGroupEnd(groupTest)
}
