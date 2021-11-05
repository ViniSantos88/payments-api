package handlers_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	servicesmocks "github.com/ViniSantos88/payments-api/__test__/mocks/services"
	"github.com/ViniSantos88/payments-api/__test__/testutil"
	"github.com/ViniSantos88/payments-api/application/handlers"
	"github.com/ViniSantos88/payments-api/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

/**********************************************************************/
/*                           Create                                   */
/**********************************************************************/
func TestStart_CreateTransactions(t *testing.T) {
	groupTest := "CreateTransactions"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestCreateTransactionsFunction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	paymentsServiceMock := servicesmocks.NewMockPaymentsService(mockCtrl)
	pc := handlers.NewPaymentsCtrlImpl(paymentsServiceMock)

	// to set header in the request api
	header := map[string][]string{
		"Content-Type": {"application/json"},
	}

	t.Run("Test Create Transactions OK", func(t *testing.T) {
		//Mocking
		mockID := int64(123456)

		paymentsServiceMock.EXPECT().CreateTransactions(gomock.Any()).Return(mockID, nil)

		//Create a request
		jsonRequestBody :=
			`{
				"account_id": 1, 
				"operation_type_id": 4, 
				"amount": 123.45
			}`

		req, err := http.NewRequest("POST", "/transactions", strings.NewReader(jsonRequestBody))
		req.Header = header

		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.CreateTransactions)

		//Call handler-controller
		pc.CreateTransactions(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusCreated, rr.Code, "Invalid Status code: JSON %s", rr.Body.String())

		var data domain.Transactions

		err = json.NewDecoder(rr.Body).Decode(&data.TransactionID)
		testutil.ValidateError(t, err)

		require.Equal(t, mockID, data.TransactionID, "Transaction ID is wrong")
	})

	t.Run("Test Invalid JSON when create", func(t *testing.T) {
		errMsg := domain.ErrInvalidBody
		//Create a request
		jsonRequestBody := `{
			"account_id": 1, 
			"operation_type_id": 4, 
			"amount": 123.45,
		}`

		req, err := http.NewRequest("POST", "/transactions", strings.NewReader(jsonRequestBody))
		req.Header = header
		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.CreateTransactions)

		//Call handler-controller
		pc.CreateTransactions(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusBadRequest, rr.Code, "Invalid Status code: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, errMsg.Error(), resp.Message, "Invalid Body")
	})

	t.Run("Test Required Account ID JSON when create", func(t *testing.T) {
		msgError := domain.ErrRequiredAccountID
		//Create a request
		jsonRequestBody := `{
			"operation_type_id": 4, 
			"amount": 123.45
		}`

		req, err := http.NewRequest("POST", "/transactions", strings.NewReader(jsonRequestBody))
		req.Header = nil
		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.CreateTransactions)

		//Call handler-controller
		pc.CreateTransactions(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusBadRequest, rr.Code, "Invalid Status code: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, msgError.Error(), resp.Message, "Message is wrong")
	})

	t.Run("Test Internal Server error when create", func(t *testing.T) {
		//Mocking
		expectedError := domain.ErrInternalServerError
		paymentsServiceMock.EXPECT().CreateTransactions(gomock.Any()).Return(int64(0), expectedError)

		//Create a request
		jsonRequestBody := `{
			"account_id": 1, 
			"operation_type_id": 4, 
			"amount": 123.45
		}`

		req, err := http.NewRequest("POST", "/transactions", strings.NewReader(jsonRequestBody))
		req.Header = header
		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.CreateTransactions)

		//Call handler-controller
		pc.CreateTransactions(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusInternalServerError, rr.Code, "Invalid Status code: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, expectedError.Error(), resp.Message, "Message is wrong")
	})

}

func TestEnd_CreateTransactions(t *testing.T) {
	groupTest := "Create_Transactions"
	testutil.PrintTestGroupEnd(groupTest)
}
