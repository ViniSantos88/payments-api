package domain

import (
	"encoding/json"
	"net/http"
)

type Accounts struct {
	AccountID      int64  `json:"account_id" db:"account_id"`
	DocumentNumber string `json:"document_number,omitempty" db:"document_number"`
}

func ValidateAccount(r *http.Request) (*Accounts, error) {
	if r.Body == http.NoBody || r.Body == nil {
		return nil, ErrEmptyBody
	}

	var account Accounts
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		return nil, ErrInvalidBody
	}

	if account.DocumentNumber == "" {
		return nil, ErrRequiredDocNumber
	}

	return &account, nil
}
