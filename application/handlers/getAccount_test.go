package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	servicesmocks "github.com/ViniSantos88/payments-api/__test__/mocks/services"
	"github.com/ViniSantos88/payments-api/__test__/testutil"
	"github.com/ViniSantos88/payments-api/application/handlers"
	"github.com/ViniSantos88/payments-api/domain"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

/**********************************************************************/
/*                             GET                                   */
/**********************************************************************/
func TestStart_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestGetAccountFunction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	paymentsServiceMock := servicesmocks.NewMockPaymentsService(mockCtrl)
	pc := handlers.NewPaymentsCtrlImpl(paymentsServiceMock)

	accountID := int64(12345)

	//Create a response
	accountResponse :=
		`{
			"account_id": 12345,
			"document_number": "12345678900" 
		}`

	t.Run("Test Get Account OK", func(t *testing.T) {
		//Mocking

		responseMock := &domain.Accounts{}
		err := json.NewDecoder(strings.NewReader(accountResponse)).Decode(&responseMock)
		testutil.ValidateError(t, err)

		paymentsServiceMock.EXPECT().GetAccount(gomock.Any()).Return(responseMock, nil)

		url := fmt.Sprintf("/accounts/%d", accountID)
		req, err := http.NewRequest("GET", url, nil)

		req = mux.SetURLVars(req, map[string]string{"accountId": strconv.FormatInt(accountID, 10)})

		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.GetAccount)

		//Call handler-controller
		pc.GetAccount(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusOK, rr.Code, "Returned wrong status code: JSON %s", rr.Body.String())

		var data domain.Accounts

		err = json.NewDecoder(rr.Body).Decode(&data)
		testutil.ValidateError(t, err)

		require.Equal(t, accountID, data.AccountID, "Invalid account ID")
	})

	t.Run("Test Invalid path when get", func(t *testing.T) {
		msgError := domain.ErrRequiredAccountID
		url := fmt.Sprintf("/accounts/%s", "")
		req, err := http.NewRequest("GET", url, nil)

		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.GetAccount)

		//Call handler-controller
		pc.GetAccount(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusBadRequest, rr.Code, "Returned wrong status code: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, msgError.Error(), resp.Message, "Invalid ID")
	})

	t.Run("Test Invalid account ID path when get", func(t *testing.T) {
		msgError := domain.ErrInvalidAccountID
		url := fmt.Sprintf("/accounts/%s", "asdfg4")
		req, err := http.NewRequest("GET", url, nil)
		req = mux.SetURLVars(req, map[string]string{"accountId": "asdfg4"})

		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.GetAccount)

		//Call handler-controller
		pc.GetAccount(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusBadRequest, rr.Code, "Returned wrong status code: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, msgError.Error(), resp.Message, "Invalid Account ID")
	})

	t.Run("Test error service account when get", func(t *testing.T) {
		//Mocking
		expectedError := domain.ErrInternalServerError
		paymentsServiceMock.EXPECT().GetAccount(gomock.Any()).Return(nil, expectedError)

		url := fmt.Sprintf("/accounts/%d", accountID)
		req, err := http.NewRequest("GET", url, nil)
		req = mux.SetURLVars(req, map[string]string{"accountId": strconv.FormatInt(accountID, 10)})
		testutil.ValidateError(t, err)

		//Initialize response
		rr, _ := testutil.PrepareHandler(pc.GetAccount)

		//Call handler-controller
		pc.GetAccount(rr, req)

		// Check the status code is what we expect
		require.Equal(t, http.StatusInternalServerError, rr.Code, "Status code incorreto: JSON %s", rr.Body.String())

		resp := &handlers.GenericResponse{}

		err = json.NewDecoder(rr.Body).Decode(&resp)
		testutil.ValidateError(t, err)

		require.Equal(t, expectedError.Error(), resp.Message, "Message is wrong")
	})

}

func TestEnd_GetAccount(t *testing.T) {
	groupTest := "GetAccount"
	testutil.PrintTestGroupEnd(groupTest)
}
