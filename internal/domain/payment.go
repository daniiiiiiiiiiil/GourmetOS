package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Payment struct {
	PaymentID     int
	OrderID       int
	Amount        float64
	Method        string
	Status        string
	TransactionID string
	CardLast4     string
	CryptoAddress string
	ReceiptURL    string
	PaidAt        time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPayment(
	paymentID int,
	orderID int,
	amount float64,
	method string,
	status string,
	transactionID string,
	cardLast4 string,
	cryptoAddress string,
	receiptURL string,
	paidAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *Payment {
	return &Payment{
		PaymentID:     paymentID,
		OrderID:       orderID,
		Amount:        amount,
		Method:        method,
		Status:        status,
		TransactionID: transactionID,
		CardLast4:     cardLast4,
		CryptoAddress: cryptoAddress,
		ReceiptURL:    receiptURL,
		PaidAt:        paidAt,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func (p *Payment) Validate() error {
	var errs errors.ValidationErrors
	if p.OrderID < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "OrderID",
			Message: "OrderID не может быть меньше 0",
		})
	}
	if p.Amount < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Amount",
			Message: "amount не может быть меньше нуля",
		})
	}
	if p.Method == "" || len(p.Method) < 0 || len(p.Method) > 30 {
		errs = append(errs, errors.ValidationError{
			Field:   "Method",
			Message: "Метод не может быть пустым и не может быть больше 30 символов",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
