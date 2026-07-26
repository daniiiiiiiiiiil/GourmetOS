package command

import "fmt"

type SubmitToKitchenCommand struct {
	orderID int
}

func NewSubmitToKitchenCommand(orderID int) *SubmitToKitchenCommand {
	return &SubmitToKitchenCommand{
		orderID: orderID,
	}
}

func (c *SubmitToKitchenCommand) Execute() error {
	fmt.Printf("Отправка заказа #%d на кухню\n", c.orderID)
	return nil
}

func (c *SubmitToKitchenCommand) Undo() error {
	fmt.Printf("Возврат заказа #%d с кухни\n", c.orderID)
	return nil
}

func (c *SubmitToKitchenCommand) GetName() string {
	return fmt.Sprintf("Отправка заказа #%d на кухню", c.orderID)
}
