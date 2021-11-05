package domain

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	ConstCompraAVista = iota + 1
	ConstCompraParcelada
	ConstSaque
	ConstPagamento
)

type Transactions struct {
	TransactionID   int64     `json:"transaction_id," db:"transaction_id"`
	AccountID       int64     `json:"account_id,omitempty" db:"account_id"`
	OperationTypeID int64     `json:"operation_type_id,omitempty" db:"operation_type_id"`
	Amount          float64   `json:"amount,omitempty" db:"amount"`
	EventDate       time.Time `json:"event_date,omitempty" db:"event_date"`
}

func ValidateTransaction(r *http.Request) (*Transactions, error) {
	if r.Body == http.NoBody || r.Body == nil {
		return nil, ErrEmptyBody
	}

	var transaction Transactions
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		return nil, ErrInvalidBody
	}

	if transaction.AccountID == 0 {
		return nil, ErrRequiredAccountID
	}

	if transaction.OperationTypeID == 0 {
		return nil, ErrRequiredOperationType
	}

	if transaction.Amount == 0 {
		return nil, ErrRequiredAmount
	}

	if transaction.OperationTypeID == ConstCompraAVista ||
		transaction.OperationTypeID == ConstCompraParcelada ||
		transaction.OperationTypeID == ConstSaque {
		transaction.Amount *= -1
	}

	return &transaction, nil
}
