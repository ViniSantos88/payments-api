package domain

import "errors"

var (
	// ErrEmptyBody empty body
	ErrEmptyBody = errors.New("Empty body")
	// ErrInvalidBody invalid body
	ErrInvalidBody = errors.New("Invalid body")
	// ErrInvalidAccountID invalid account ID
	ErrInvalidAccountID = errors.New("Invalid Account ID")
	// ErrRequiredDocNumber required document number
	ErrRequiredDocNumber = errors.New("Required Document Number")
	// ErrRequiredAccountID required account id
	ErrRequiredAccountID = errors.New("Required Account ID")
	// ErrRequiredOperationType required operation type
	ErrRequiredOperationType = errors.New("Required Operation Type")
	// ErrRequiredAmount required document number
	ErrRequiredAmount = errors.New("Required amount")
	// ErrAccountNotFound not found
	ErrAccountNotFound = errors.New("Bank account not found")
	// ErrTransactionNotFound not found
	ErrTransactionNotFound = errors.New("Transaction not found")
	// ErrInternalServerError is internal server error
	ErrInternalServerError = errors.New("Internal Server Error")
)
