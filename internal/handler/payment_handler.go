package handler

import (
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	var req dto.ProcessPaymentRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidValidate", err.Error())
		return
	}

	err := h.service.ProcessPayment(
		r.Context(),
		conn,
		req.OrderID,
		req.Method,
		req.CardNumber,
		req.CardHolder,
		req.Expiry,
		req.CVV,
		req.Wallet,
	)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}

	sendSuccess(w, http.StatusCreated, nil)
}

func (h *PaymentHandler) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["payment_id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidPaymentID", err.Error())
	}
	pay, err := h.service.GetPaymentStatus(r.Context(), conn, paymentID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, pay)
}

func (h *PaymentHandler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["payment_id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidPaymentID", err.Error())
		return
	}
	if err := h.service.RefundPayment(r.Context(), conn, paymentID); err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *PaymentHandler) GetPaymentByOrder(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidOrderID", err.Error())
		return
	}
	pay, err := h.service.GetPaymentByOrder(r.Context(), conn, orderID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, pay)
}

func (h *PaymentHandler) GetPaymentByTransaction(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]

	pay, err := h.service.GetPaymentByTransaction(r.Context(), conn, transactionID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	resp := dto.FromDomainPayments(*pay)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *PaymentHandler) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidOrderID", err.Error())
		return
	}
	pay, err := h.service.GetPaymentByOrder(r.Context(), conn, orderID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	resp := dto.FromDomainPayments(*pay)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *PaymentHandler) ApplyDiscount(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])

	discountType := vars["discount_type"]

	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidOrderID", err.Error())
		return
	}
	pay, err := h.service.ApplyDiscount(r.Context(), conn, orderID, discountType)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, pay)
}

func (h *PaymentHandler) CalculateFinalAmount(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["order_id"])
	discountType := vars["discount_type"]
	if err != nil {
		sendError(w, http.StatusNotFound, "InvalidOrderID", err.Error())
		return
	}
	pay, err := h.service.CalculateFinalAmount(r.Context(), conn, orderID, discountType)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ErrorServer", err.Error())
	}
	sendSuccess(w, http.StatusOK, pay)
}
