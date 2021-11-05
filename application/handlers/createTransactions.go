package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ViniSantos88/payments-api/domain"
)

// swagger:route POST /transacations createTransactions createTransactions
// Creation of payment method transactions
//
// responses:
//	201: transactionIDResponse
//  400: errorValidation
//  404: errorNotFound
//  500: description: Internal Server Error

//CreateTransactions is the creation of payment method transactions
func (pc *PaymentsCtrlImpl) CreateTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	transaction, err := domain.ValidateTransaction(r)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		json.NewEncoder(w).Encode(&GenericResponse{Message: err.Error()})
		return
	}

	response := &domain.Transactions{}
	response.TransactionID, err = pc.PaymentsService.CreateTransactions(*transaction)
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
