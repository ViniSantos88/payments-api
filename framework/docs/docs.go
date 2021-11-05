// Package classification Payments.
//
// Payments API
//
//     Schemes: [http, https]
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//
// swagger:meta
package docs

import (
	"github.com/ViniSantos88/payments-api/domain"
)

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Not Found
// swagger:response errorNotFound
type errorValidationNotFoundWrapper struct {
	// Description of error
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Bad Request
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Description of validation error
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Data structure representing the id
// swagger:response accountIDResponse
type createAccountIDResponseWrapper struct {
	// Return with id
	// in: body
	Body struct {
		AccountID int64 `json:"account_id"`
	}
}

// swagger:parameters accountRequest createAccount
type AccountParamsWrapper struct {
	// Data structure representing the account
	// in: body
	// required: true
	Body struct {
		DocumentNumber string `json:"document_number"`
	}
}

// Data structure representing the account
// swagger:response accountResponse
type accountResponseWrapper struct {
	// Return with account id
	// in: body
	Body struct {
		Data domain.Accounts `json:"data"`
	}
}

// swagger:parameters accountId getAccount
type accountIDParamsWrapper struct {
	// Data structure representing the account ID
	// in: path
	// required: true
	AcccountID string `json:"accountId"`
}

// Data structure representing the id
// swagger:response transactionIDResponse
type createTransactionIDResponseWrapper struct {
	// Return with id
	// in: body
	Body struct {
		TransactionID int64 `json:"transaction_id"`
	}
}

// swagger:parameters transactionRequest createTransactions
type TransactionParamsWrapper struct {
	// Data structure representing the transaction
	// in: body
	// required: true
	Body struct {
		AccountID       int64   `json:"account_id"`
		OperationTypeID int64   `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}
}
