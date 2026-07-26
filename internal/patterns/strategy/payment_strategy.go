package strategy

type PaymentStrategy interface {
	Pay(amount float64) error
	GetName() string
}
