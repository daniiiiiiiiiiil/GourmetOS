package command

import "fmt"

type RefundPaymentCommand struct {
	paymentID int
	amount    float64
}

func NewRefundPaymentCommand(paymentID int, amount float64) *RefundPaymentCommand {
	return &RefundPaymentCommand{
		paymentID: paymentID,
		amount:    amount,
	}
}

func (c *RefundPaymentCommand) Execute() error {
	fmt.Printf("Возврат денег за платёж #%d (%.2f руб.)\n", c.paymentID, c.amount)
	return nil
}

func (c *RefundPaymentCommand) Undo() error {
	fmt.Printf("Отмена возврата для платежа #%d (%.2f руб.)\n", c.paymentID, c.amount)
	return nil
}

func (c *RefundPaymentCommand) GetName() string {
	return fmt.Sprintf("Возврат денег за платёж #%d", c.paymentID)
}
