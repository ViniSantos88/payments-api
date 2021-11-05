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
/*                           GetAccount                               */
/**********************************************************************/
func TestStart_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupStart(groupTest)
	//testutil.SetUp()
}

func TestGetAccountFunction(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	paymentsRepoMock := repomocks.NewMockPaymentsRepo(mockCtrl)
	ps := services.NewPaymentsServiceImpl(paymentsRepoMock)

	accountMock := domain.Accounts{
		AccountID:      12345,
		DocumentNumber: "12345678900",
	}

	t.Run("Test get account OK", func(t *testing.T) {
		accountID := int64(123456)

		paymentsRepoMock.EXPECT().GetAccount(gomock.Any()).Return(accountMock, nil)

		accountResponse, _ := ps.GetAccount(accountID)

		require.Equal(t, accountMock.AccountID, accountResponse.AccountID, "AccountiD is wrong")
		require.Equal(t, accountMock.DocumentNumber, accountResponse.DocumentNumber, "Document number is wrong")
	})

	t.Run("Test get account error", func(t *testing.T) {
		accountID := int64(123456)
		expectedError := domain.ErrInternalServerError
		paymentsRepoMock.EXPECT().GetAccount(gomock.Any()).Return(domain.Accounts{}, expectedError)

		_, err := ps.GetAccount(accountID)

		require.Equal(t, expectedError.Error(), err.Error(), "Error message is wrong")
	})

}
func TestEnd_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupEnd(groupTest)
}
