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
/*                         CreateAccounts                             */
/**********************************************************************/
func TestStart_CreateAccounts(t *testing.T) {
	groupTest := "CreateAccounts"
	testutil.PrintTestGroupStart(groupTest)
	//testutil.SetUp()
}

func TestCreateAccountsFunction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	paymentsRepoMock := repomocks.NewMockPaymentsRepo(mockCtrl)
	ps := services.NewPaymentsServiceImpl(paymentsRepoMock)

	account := domain.Accounts{
		DocumentNumber: "12345678900",
	}

	t.Run("Test create account OK", func(t *testing.T) {
		mockID := int64(123456)

		paymentsRepoMock.EXPECT().CreateAccounts(gomock.Any()).Return(mockID, nil)

		accountID, _ := ps.CreateAccounts(account)

		require.Equal(t, mockID, accountID, "Id is wrong")
	})

	t.Run("Test get account error", func(t *testing.T) {
		expectedError := domain.ErrInternalServerError
		paymentsRepoMock.EXPECT().CreateAccounts(gomock.Any()).Return(int64(0), expectedError)

		_, err := ps.CreateAccounts(account)

		require.Equal(t, expectedError.Error(), err.Error(), "Error message is wrong")
	})

}
func TestEnd_CreateAccounts(t *testing.T) {
	groupTest := "CreateAccounts"
	testutil.PrintTestGroupEnd(groupTest)
}
