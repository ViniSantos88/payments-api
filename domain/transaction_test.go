package domain_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ViniSantos88/payments-api/__test__/testutil"
	"github.com/ViniSantos88/payments-api/domain"
	"github.com/stretchr/testify/require"
)

/**********************************************************************/
/*                Validate Domain Transaction                         */
/**********************************************************************/
func TestStart_DomainTransaction(t *testing.T) {
	groupTest := "Domain Transaction"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestTransactionDomain(t *testing.T) {

	t.Run("Test account domain ok", func(t *testing.T) {
		jsonRequestBody := `{
			"account_id": 213452,
			"operation_type_id": 1 ,
			"amount": 120.90 
		}`

		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}

		account, _ := domain.ValidateTransaction(httpRequest)

		require.NotNil(t, account.AccountID, "Account ID is wrong")
		require.NotNil(t, account.OperationTypeID, "Operation Type number is wrong")
		require.NotNil(t, account.Amount, "Amount is wrong")

	})

	t.Run("Test Empty body", func(t *testing.T) {
		errMsg := domain.ErrEmptyBody

		httpRequest := &http.Request{
			Body: http.NoBody,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Body is wrong")
	})

	t.Run("Test invalid body", func(t *testing.T) {
		errMsg := domain.ErrInvalidBody

		jsonRequestBody := `{
			"account_id": 213452,
			"operation_type_id": 1 ,
			"amount": 120.90,
		}`

		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test required account id field", func(t *testing.T) {
		errMsg := domain.ErrRequiredAccountID

		jsonRequestBody := `{
			"operation_type_id": 1 ,
			"amount": 120.90
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test required operation type field", func(t *testing.T) {
		errMsg := domain.ErrRequiredOperationType

		jsonRequestBody := `{
			"account_id": 213452,
			"amount": 120.90
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test required amount field", func(t *testing.T) {
		errMsg := domain.ErrRequiredAmount

		jsonRequestBody := `{
			"account_id": 213452,
			"operation_type_id": 1 
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test invalid operationType", func(t *testing.T) {
		errMsg := domain.ErrInvalidOperationType

		jsonRequestBody := `{
			"account_id": 213452,
			"operation_type_id": 5 ,
			"amount": 120.90
		}`

		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateTransaction(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test amount by operationType Compra a Vista", func(t *testing.T) {
		amount := -120.90
		jsonRequestBody := `{
			"account_id": 213452,
			"operation_type_id": 1 ,
			"amount": 120.90
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		transaction, _ := domain.ValidateTransaction(httpRequest)

		require.Equal(t, amount, transaction.Amount, "Validation message is wrong")
	})

}

func TestEnd_DomainTransaction(t *testing.T) {
	groupTest := "Domain Transaction"
	testutil.PrintTestGroupEnd(groupTest)
}
