package strategy

import (
	"fmt"
	"math/rand"
	"time"
)

type CardPayment struct {
	CardNumber string
	CardHolder string
	ExpiryDate string
	CVV        string
}

func NewCardPayment(cardNumber, cardHolder, expiryDate, cvv string) *CardPayment {
	return &CardPayment{
		CardNumber: cardNumber,
		CardHolder: cardHolder,
		ExpiryDate: expiryDate,
		CVV:        cvv,
	}
}

func (c *CardPayment) Pay(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть больше 0")
	}

	if len(c.CardNumber) < 16 {
		return fmt.Errorf("неверный номер карты")
	}

	fmt.Printf(" Оплата картой: %.2f руб.\n", amount)
	fmt.Printf("   Карта: ****-%s\n", c.CardNumber[len(c.CardNumber)-4:])
	fmt.Println("   Обработка платежа...")

	time.Sleep(500 * time.Millisecond)

	if rand.Float64() < 0.1 {
		return fmt.Errorf("ошибка оплаты: недостаточно средств")
	}

	fmt.Println("Платеж одобрен!")
	return nil
}

func (c *CardPayment) GetName() string {
	return "Банковская карта"
}
