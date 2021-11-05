package postgresql_test

import (
	"database/sql"
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
/*                            GetAccount                             */
/**********************************************************************/
func TestStart_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestGetAccountFunction(t *testing.T) {
	//init sqlMock
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	pr := postgresql.NewPostgreSQLRepoImpl(sqlxDB)

	t.Run("Test get account ok", func(t *testing.T) {
		accountID := int64(12345)
		rows := sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow(12345, "12345678900")

		mock.ExpectQuery("SELECT  account_id, document_number").WillReturnRows(rows)

		account, _ := pr.GetAccount(accountID)

		require.NotNil(t, account, "Record recording error")
		require.Equal(t, accountID, account.AccountID, "Account id is wrong")
	})

	t.Run("Test get account not found", func(t *testing.T) {
		errMsg := domain.ErrAccountNotFound
		accountID := int64(12345)

		rows := sqlmock.NewRows([]string{"account_id", "document_number"}).RowError(0, sql.ErrNoRows)
		mock.ExpectQuery("SELECT  account_id, document_number").WillReturnRows(rows)

		_, err := pr.GetAccount(accountID)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")

	})

	t.Run("Test error database", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		accountID := int64(12345)
		mock.ExpectQuery("SELECT  account_id, document_number").WillReturnError(fmt.Errorf("error"))
		_, err := pr.GetAccount(accountID)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})

	t.Run("Test error mapping data", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		accountID := int64(12345)
		rows := sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("234sdf", "12345678900")

		mock.ExpectQuery("SELECT  account_id, document_number").WillReturnRows(rows)

		_, err := pr.GetAccount(accountID)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})
}
func TestEnd_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupEnd(groupTest)
}
