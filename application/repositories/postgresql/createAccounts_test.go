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
/*                          CreateAccounts                            */
/**********************************************************************/
func TestStart_CreateAccounts(t *testing.T) {
	groupTest := "CreateAccounts"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestCreateAccountsFunction(t *testing.T) {
	//init sqlMock
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	account := domain.Accounts{
		DocumentNumber: "12345678900",
	}

	pr := postgresql.NewPostgreSQLRepoImpl(sqlxDB)

	t.Run("Test create account ok", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"account_id"}).AddRow(123456)

		mock.ExpectQuery("INSERT INTO accounts").WillReturnRows(rows)

		accountID, _ := pr.CreateAccounts(account)

		require.NotNil(t, accountID, "Record recording error")
	})

	t.Run("Test error database", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		mock.ExpectQuery("INSERT INTO accounts").WillReturnError(fmt.Errorf("error"))
		_, err := pr.CreateAccounts(account)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})

	t.Run("Test error mapping data", func(t *testing.T) {
		errMsg := domain.ErrInternalServerError
		rows := sqlmock.NewRows([]string{"account_id"}).AddRow("8d9fs")

		mock.ExpectQuery("INSERT INTO accounts").WillReturnRows(rows)

		_, err := pr.CreateAccounts(account)

		require.Equal(t, errMsg.Error(), err.Error(), "Error message is wrong")
	})
}
func TestEnd_CreateAccounts(t *testing.T) {
	groupTest := "CreateAccounts"
	testutil.PrintTestGroupEnd(groupTest)
}
