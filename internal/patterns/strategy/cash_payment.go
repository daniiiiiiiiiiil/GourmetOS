package strategy

import "fmt"

type CashPayment struct {
	Amount float64
}

func NewCashPayment() *CashPayment {
	return &CashPayment{}
}

func (c *CashPayment) Pay(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("сумма должна быть больше 0")
	}

	fmt.Printf("Оплачено наличными: %.2f руб.\n", amount)
	fmt.Println("Кассир принял деньги, сдача выдана.")

	return nil
}

func (c *CashPayment) GetName() string {
	return "Наличные"
}
