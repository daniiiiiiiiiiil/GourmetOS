package adapter

import (
	"fmt"
	"math/rand"
	"time"
)

type CryptoAPI struct{}

func NewCryptoAPI() *CryptoAPI {
	return &CryptoAPI{}
}

// отправка криптовалюты
func (c *CryptoAPI) SendCrypto(fromWallet, toWallet string, amount float64, cryptoType string) (string, error) {
	fmt.Printf("   [Crypto API] Отправка %.8f %s с %s на %s\n",
		amount, cryptoType, fromWallet[:8], toWallet[:8])

	time.Sleep(2 * time.Second)

	if rand.Float64() < 0.05 {
		return "", fmt.Errorf("Crypto API: недостаточно средств на кошельке")
	}

	txnID := fmt.Sprintf("CRYPTO_%d", time.Now().UnixNano())
	fmt.Printf("Транзакция подтверждена: %s\n", txnID)
	return txnID, nil
}

func (c *CryptoAPI) GetCryptoStatus(txnID string) (string, error) {
	fmt.Printf("   [Crypto API] Статус транзакции: %s\n", txnID)
	return "confirmed", nil
}

func (c *CryptoAPI) RefundCrypto(txnID string) error {
	fmt.Printf("   [Crypto API] Возврат средств: %s\n", txnID)
	return nil
}

type CryptoAdapter struct {
	api        *CryptoAPI
	fromWallet string
	toWallet   string
	cryptoType string
}

func NewCryptoAdapter(fromWallet, toWallet, cryptoType string) *CryptoAdapter {
	return &CryptoAdapter{
		api:        NewCryptoAPI(),
		fromWallet: fromWallet,
		toWallet:   toWallet,
		cryptoType: cryptoType,
	}
}

// оплата криптовалютой
func (c *CryptoAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	cryptoAmount := amount / 5000
	fmt.Printf("Оплата через Crypto Adapter: %.2f %s (%.8f %s)\n",
		amount, currency, cryptoAmount, c.cryptoType)

	return c.api.SendCrypto(c.fromWallet, c.toWallet, cryptoAmount, c.cryptoType)
}

func (c *CryptoAdapter) GetPaymentStatus(transactionID string) (string, error) {
	fmt.Printf("Проверка статуса Crypto: %s\n", transactionID)
	return "confirmed", nil
}

func (c *CryptoAdapter) RefundPayment(transactionID string) error {
	fmt.Printf("Возврат через Crypto: %s\n", transactionID)
	return c.api.RefundCrypto(transactionID)
}

func (c *CryptoAdapter) GetAdapterName() string {
	return "Crypto Adapter"
}
