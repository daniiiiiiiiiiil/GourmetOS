package strategy

import (
	"fmt"
	"math/rand"
	"time"
)

type CryptoPayment struct {
	WalletAddress string
	CryptoType    string // BTC, ETH, USDT
}

func NewCryptoPayment(walletAddress, cryptoType string) *CryptoPayment {
	return &CryptoPayment{
		WalletAddress: walletAddress,
		CryptoType:    cryptoType,
	}
}

func (c *CryptoPayment) Pay(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть больше 0")
	}

	cryptoAmount := amount / 5000 // 1 BTC = 5000 руб.

	fmt.Printf("🪙 Оплата криптовалютой: %.2f руб. (~%.8f %s)\n", amount, cryptoAmount, c.CryptoType)
	fmt.Printf("   Кошелек: %s\n", c.WalletAddress)
	fmt.Println("   Ожидание подтверждения...")

	time.Sleep(1500 * time.Millisecond)

	if rand.Float64() < 0.05 {
		return fmt.Errorf("ошибка: недостаточно средств на кошельке")
	}

	fmt.Printf("Транзакция подтверждена! Оплачено %.8f %s\n", cryptoAmount, c.CryptoType)
	return nil
}

func (c *CryptoPayment) GetName() string {
	return fmt.Sprintf("Криптовалюта (%s)", c.CryptoType)
}
