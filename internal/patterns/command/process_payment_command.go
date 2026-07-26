package command

import "fmt"

type ProcessPaymentCommand struct {
	orderID int
	method  string
	amount  float64
}

func NewProcessPaymentCommand(orderID int, method string, amount float64) *ProcessPaymentCommand {
	return &ProcessPaymentCommand{
		orderID: orderID,
		method:  method,
		amount:  amount,
	}
}

func (c *ProcessPaymentCommand) Execute() error {
	fmt.Printf("Оплата заказа #%d на %.2f руб. методом %s\n", c.orderID, c.amount, c.method)
	return nil
}

func (c *ProcessPaymentCommand) Undo() error {
	fmt.Printf("Возврат средств за заказ #%d (%.2f руб.)\n", c.orderID, c.amount)
	return nil
}

func (c *ProcessPaymentCommand) GetName() string {
	return fmt.Sprintf("Оплата заказа #%d (%s)", c.orderID, c.method)
}
