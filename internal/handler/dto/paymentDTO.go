package dto

import (
	"GourmetOS/internal/domain"
	"GourmetOS/pkg/errors"
	"GourmetOS/pkg/pagination"
	"time"
)

type ProcessPaymentRequestDTO struct {
	OrderID    int    `json:"order_id" binding:"required"`
	Method     string `json:"method" binding:"required,oneof=cash card crypto"`
	CardNumber string `json:"card_number"`
	CardHolder string `json:"card_holder"`
	Expiry     string `json:"expiry"`
	CVV        string `json:"cvv"`
	Wallet     string `json:"wallet"`
}

func (r *ProcessPaymentRequestDTO) Validate() error {
	var errs errors.ValidationErrors

	if r.OrderID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "order_id",
			Message: "ID заказа не может быть меньше или равен 0",
		})
	}

	if r.Method == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "method",
			Message: "Способ оплаты не может быть пустым",
		})
	}

	if r.Method == "card" {
		if r.CardNumber == "" {
			errs = append(errs, errors.ValidationError{
				Field:   "card_number",
				Message: "Номер карты не может быть пустым",
			})
		}
		if r.CardHolder == "" {
			errs = append(errs, errors.ValidationError{
				Field:   "card_holder",
				Message: "Имя владельца карты не может быть пустым",
			})
		}
		if r.Expiry == "" {
			errs = append(errs, errors.ValidationError{
				Field:   "expiry",
				Message: "Срок действия карты не может быть пустым",
			})
		}
		if r.CVV == "" {
			errs = append(errs, errors.ValidationError{
				Field:   "cvv",
				Message: "CVV не может быть пустым",
			})
		}
	}

	if r.Method == "crypto" && r.Wallet == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "wallet",
			Message: "Адрес кошелька не может быть пустым",
		})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type RefundPaymentRequest struct {
	PaymentID int `json:"payment_id" binding:"required"`
}

func (r *RefundPaymentRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.PaymentID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "payment_id",
			Message: "ID платежа не может быть меньше или равен 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

type PaymentResponse struct {
	PaymentID     int       `json:"payment_id"`
	OrderID       int       `json:"order_id"`
	Amount        float64   `json:"amount"`
	Method        string    `json:"method"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"`
	CardLast4     string    `json:"card_last4"`
	CryptoAddress string    `json:"crypto_address"`
	ReceiptURL    string    `json:"receipt_url"`
	PaidAt        time.Time `json:"paid_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func PaymentResponseFromDomain(payment domain.Payment) PaymentResponse {
	return PaymentResponse{
		PaymentID:     payment.PaymentID,
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		Method:        payment.Method,
		Status:        payment.Status,
		TransactionID: payment.TransactionID,
		CardLast4:     payment.CardLast4,
		CryptoAddress: payment.CryptoAddress,
		ReceiptURL:    payment.ReceiptURL,
		PaidAt:        payment.PaidAt,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
}

type PaymentListResponse struct {
	Payments   []PaymentResponse     `json:"payments"`
	Pagination pagination.Pagination `json:"pagination"`
}

func NewPaymentListResponse(payments []domain.Payment, total, limit, offset int) PaymentListResponse {
	resp := PaymentListResponse{
		Payments:   make([]PaymentResponse, 0, len(payments)),
		Pagination: pagination.NewPagination(total, limit, offset),
	}
	for _, payment := range payments {
		resp.Payments = append(resp.Payments, PaymentResponseFromDomain(payment))
	}
	return resp
}

func FromDomainPayments(payments domain.Payment) PaymentResponse {
	return PaymentResponse{
		PaymentID:     payments.PaymentID,
		OrderID:       payments.OrderID,
		Amount:        payments.Amount,
		Method:        payments.Method,
		Status:        payments.Status,
		TransactionID: payments.TransactionID,
		CardLast4:     payments.CardLast4,
		CryptoAddress: payments.CryptoAddress,
		ReceiptURL:    payments.ReceiptURL,
		PaidAt:        payments.PaidAt,
		CreatedAt:     payments.CreatedAt,
		UpdatedAt:     payments.UpdatedAt,
	}
}
