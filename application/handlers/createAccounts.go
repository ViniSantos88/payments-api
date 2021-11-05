package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ViniSantos88/payments-api/domain"
)

// swagger:route POST /accounts createAccount createAccount
// Creation of bank account
//
// responses:
//	201: accountIDResponse
//  400: errorValidation
//  404: errorNotFound
//  500: description: Internal Server Error

//CreateAccounts is the creation of bank accounts
func (pc *PaymentsCtrlImpl) CreateAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	account, err := domain.ValidateAccount(r)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		json.NewEncoder(w).Encode(&GenericResponse{Message: err.Error()})
		return
	}

	response := &domain.Accounts{}
	response.AccountID, err = pc.PaymentsService.CreateAccounts(*account)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		json.NewEncoder(w).Encode(&GenericResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
