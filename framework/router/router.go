package router

import (
	"github.com/ViniSantos88/payments-api/application/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	pc := handlers.GetPaymentsCtrlImpl()

	r.Methods("POST").Path("/accounts").HandlerFunc(pc.CreateAccounts)
	r.Methods("GET").Path("/accounts/{accountId}").HandlerFunc(pc.GetAccount)
	r.Methods("POST").Path("/transactions").HandlerFunc(pc.CreateTransactions)

	return r

}
