package adapter

import (
	"fmt"
	"math/rand"
	"time"
)

// внешнее API PayPal
type PayPalAPI struct{}

// конструктор
func NewPayPalAPI() *PayPalAPI {
	return &PayPalAPI{}
}

// создание платежа через PayPal
func (p *PayPalAPI) CreatePayment(email string, amount float64, currency string) (string, error) {
	fmt.Printf("   [PayPal API] Создание платежа: %s, %.2f %s\n", email, amount, currency)

	time.Sleep(300 * time.Millisecond)

	if rand.Float64() < 0.1 {
		return "", fmt.Errorf("PayPal API: ошибка авторизации аккаунта")
	}

	txnID := fmt.Sprintf("PP_%d", time.Now().UnixNano())
	fmt.Printf("Платеж создан: %s\n", txnID)
	return txnID, nil
}

// статус платежа
func (p *PayPalAPI) GetPayPalStatus(txnID string) (string, error) {
	fmt.Printf("   [PayPal API] Статус: %s\n", txnID)
	return "completed", nil
}

// RefundPayPal — возврат
func (p *PayPalAPI) RefundPayPal(txnID string) error {
	fmt.Printf("   [PayPal API] Возврат: %s\n", txnID)
	return nil
}

// адаптер для PayPal
type PayPalAdapter struct {
	api   *PayPalAPI
	email string
}

// конструктор
func NewPayPalAdapter(email string) *PayPalAdapter {
	return &PayPalAdapter{
		api:   NewPayPalAPI(),
		email: email,
	}
}

// оплата через PayPal
func (p *PayPalAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Оплата через PayPal Adapter: %.2f %s\n", amount, currency)
	return p.api.CreatePayment(p.email, amount, currency)
}

func (p *PayPalAdapter) GetPaymentStatus(transactionID string) (string, error) {
	fmt.Printf("Проверка статуса PayPal: %s\n", transactionID)
	return p.api.GetPayPalStatus(transactionID)
}

func (p *PayPalAdapter) RefundPayment(transactionID string) error {
	fmt.Printf("Возврат через PayPal: %s\n", transactionID)
	return p.api.RefundPayPal(transactionID)
}

func (p *PayPalAdapter) GetAdapterName() string {
	return "PayPal Adapter"
}
