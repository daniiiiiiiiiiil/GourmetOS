// payment_adapter.go
package adapter

type PaymentAdapter interface {
	ProcessPayment(amount float64, currency string) (string, error)
	GetPaymentStatus(transactionID string) (string, error)
	RefundPayment(transactionID string) error
	GetAdapterName() string
}
