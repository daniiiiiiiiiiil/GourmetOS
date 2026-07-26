package adapter

import (
	"fmt"
	"math/rand"
	"time"
)

// внешнее API Visa (Типо стороняя библиотека)
type VisaAPI struct{}

// конструктор внешнего API
func NewVisaAPI() *VisaAPI {
	return &VisaAPI{}
}

// списание средств через Visa
func (v *VisaAPI) Charge(cardNumber, expiry, cvv string, amount float64, currency string) (string, error) {
	fmt.Printf("   [Visa API] Запрос списания: карта %s, сумма %.2f %s\n",
		cardNumber[len(cardNumber)-4:], amount, currency)

	time.Sleep(500 * time.Millisecond)

	if rand.Float64() < 0.1 {
		return "", fmt.Errorf("Visa API: недостаточно средств на карте")
	}

	txnID := fmt.Sprintf("VISA_%d", time.Now().UnixNano())
	fmt.Printf("Транзакция одобрена: %s\n", txnID)
	return txnID, nil
}

// получение статуса транзакции
func (v *VisaAPI) GetVisaTransaction(txnID string) (string, error) {
	fmt.Printf("   [Visa API] Запрос статуса: %s\n", txnID)
	return "completed", nil
}

// возврат средств через Visa
func (v *VisaAPI) RefundVisa(txnID string) error {
	fmt.Printf("   [Visa API] Возврат средств: %s\n", txnID)
	return nil
}

// VisaAdapter — адаптер для внешнего Visa API
type VisaAdapter struct {
	api        *VisaAPI
	cardNumber string
	expiry     string
	cvv        string
}

// NewVisaAdapter — конструктор адаптера
func NewVisaAdapter(cardNumber, expiry, cvv string) *VisaAdapter {
	return &VisaAdapter{
		api:        NewVisaAPI(),
		cardNumber: cardNumber,
		expiry:     expiry,
		cvv:        cvv,
	}
}

func (v *VisaAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Оплата через Visa Adapter: %.2f %s\n", amount, currency)

	// Адаптер преобразует вызов к внешнему API
	return v.api.Charge(v.cardNumber, v.expiry, v.cvv, amount, currency)
}

func (v *VisaAdapter) GetPaymentStatus(transactionID string) (string, error) {
	fmt.Printf(" Проверка статуса Visa: %s\n", transactionID)
	return v.api.GetVisaTransaction(transactionID)
}

// возврат средств
func (v *VisaAdapter) RefundPayment(transactionID string) error {
	fmt.Printf("↩Возврат через Visa: %s\n", transactionID)
	return v.api.RefundVisa(transactionID)
}

func (v *VisaAdapter) GetAdapterName() string {
	return "Visa Adapter"
}
