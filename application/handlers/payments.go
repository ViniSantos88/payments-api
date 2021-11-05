package handlers

import (
	"net/http"

	"github.com/ViniSantos88/payments-api/application/services"
	"github.com/ViniSantos88/payments-api/domain"
)

var paymentsCtrlImpl *PaymentsCtrlImpl

//PaymentsCtrl is the interface for Payments
type PaymentsCtrl interface {
	CreateAccounts(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
	CreateTransactions(w http.ResponseWriter, r *http.Request)
}

//PaymentsCtrlImpl is the implementation for PaymentsCtrl
type PaymentsCtrlImpl struct {
	PaymentsService services.PaymentsService
}

//GetPaymentsCtrlImpl returns an instance of PaymentsCtrlImpl
func GetPaymentsCtrlImpl() *PaymentsCtrlImpl {
	if paymentsCtrlImpl == nil {
		paymentsCtrlImpl = NewPaymentsCtrlImpl(services.GetPaymentsServiceImpl())
	}

	return paymentsCtrlImpl
}

// NewPaymentsCtrlImpl is a dependency injection an instance of PaymentsCtrlImpl
func NewPaymentsCtrlImpl(paymentsService services.PaymentsService) *PaymentsCtrlImpl {
	return &PaymentsCtrlImpl{
		PaymentsService: paymentsService,
	}
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrEmptyBody, domain.ErrInvalidBody,
		domain.ErrRequiredDocNumber, domain.ErrRequiredAccountID,
		domain.ErrRequiredOperationType, domain.ErrRequiredAmount,
		domain.ErrInvalidAccountID:
		return http.StatusBadRequest
	case domain.ErrAccountNotFound, domain.ErrTransactionNotFound:
		return http.StatusNotFound
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

//CreateResponse is the response for account or transaction creation
type CreateResponse struct {
	ID int64 `json:"ID"`
}

// Generic Response for Errors
type GenericResponse struct {
	Message string        `json:"message,omitempty"`
	Data    DataInterface `json:"data,omitempty"`
}

// DataInterface returns the data related to the request
type DataInterface interface {
}
