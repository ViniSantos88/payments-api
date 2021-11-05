package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	repomocks "github.com/ViniSantos88/payments-api/__test__/mocks/repositories"
	"github.com/ViniSantos88/payments-api/__test__/testutil"
	"github.com/ViniSantos88/payments-api/application/services"
	"github.com/ViniSantos88/payments-api/domain"
)

/**********************************************************************/
/*                         CreateTransactions                         */
/**********************************************************************/
func TestStart_CreateTransactions(t *testing.T) {
	groupTest := "CreateTransactions"
	testutil.PrintTestGroupStart(groupTest)
	//testutil.SetUp()
}

func TestCreateTransactionsFunction(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	paymentsRepoMock := repomocks.NewMockPaymentsRepo(mockCtrl)
	ps := services.NewPaymentsServiceImpl(paymentsRepoMock)

	transaction := domain.Transactions{
		AccountID:       12345,
		OperationTypeID: 1,
		Amount:          120.00,
	}

	t.Run("Test create transaction OK", func(t *testing.T) {
		mockID := int64(123456)

		paymentsRepoMock.EXPECT().CreateTransactions(gomock.Any()).Return(mockID, nil)

		transactionID, _ := ps.CreateTransactions(transaction)

		require.Equal(t, mockID, transactionID, "Id is wrong")
	})

	t.Run("Test get transaction error", func(t *testing.T) {
		expectedError := domain.ErrInternalServerError
		paymentsRepoMock.EXPECT().CreateTransactions(gomock.Any()).Return(int64(0), expectedError)

		_, err := ps.CreateTransactions(transaction)

		require.Equal(t, expectedError.Error(), err.Error(), "Error message is wrong")
	})

}
func TestEnd_CreateTransactions(t *testing.T) {
	groupTest := "CreateTransactions"
	testutil.PrintTestGroupEnd(groupTest)
}
