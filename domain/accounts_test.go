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
/*                    Validate Domain Account                         */
/**********************************************************************/
func TestStart_DomainAccount(t *testing.T) {
	groupTest := "Domain Account"
	testutil.PrintTestGroupStart(groupTest)
	//	testutil.Setup()
}

func TestAccountDomain(t *testing.T) {
	document_number := "31334493200"

	t.Run("Test account domain ok", func(t *testing.T) {
		jsonRequestBody := `{
			"document_number": "` + document_number + `" 
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}

		account, _ := domain.ValidateAccount(httpRequest)

		require.Equal(t, document_number, account.DocumentNumber, "Document number is wrong")
	})

	t.Run("Test Empty body", func(t *testing.T) {
		errMsg := domain.ErrEmptyBody

		httpRequest := &http.Request{
			Body: http.NoBody,
		}
		_, err := domain.ValidateAccount(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Body is wrong")
	})

	t.Run("Test invalid body", func(t *testing.T) {
		errMsg := domain.ErrInvalidBody

		jsonRequestBody := `{
			"document_number": "` + document_number + `",
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateAccount(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})

	t.Run("Test required field", func(t *testing.T) {
		errMsg := domain.ErrRequiredDocNumber

		jsonRequestBody := `{
			"document_number": ""
		}`
		// create a new reader with that JSON
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonRequestBody)))
		httpRequest := &http.Request{
			Body: r,
		}
		_, err := domain.ValidateAccount(httpRequest)

		require.Equal(t, errMsg.Error(), err.Error(), "Validation message is wrong")
	})
}

func TestEnd_DomainAccount(t *testing.T) {
	groupTest := "Domain Account"
	testutil.PrintTestGroupEnd(groupTest)
}
