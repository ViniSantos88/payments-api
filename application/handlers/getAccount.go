package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniSantos88/payments-api/domain"
	"github.com/gorilla/mux"
)

// swagger:route GET /accounts/{accountId} getAccount getAccount
// Get bank account detail
//
// responses:
//	200: accountResponse
//  400: errorValidation
//  404: errorNotFound
//  500: description: Internal Server Error

//GetAccount is the get bank account detail
func (pc *PaymentsCtrlImpl) GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)
	accountID := params["accountId"]
	if accountID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&GenericResponse{Message: domain.ErrRequiredAccountID.Error()})
		return
	}

	id, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&GenericResponse{Message: domain.ErrInvalidAccountID.Error()})
		return
	}

	accountResponse, err := pc.PaymentsService.GetAccount(id)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		json.NewEncoder(w).Encode(&GenericResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(accountResponse); err != nil {
		panic(err)
	}
}
