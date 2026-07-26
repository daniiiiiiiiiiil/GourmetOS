package command

import "fmt"

type ServeToTableCommand struct {
	orderID int
}

func NewServeToTableCommand(orderID int) *ServeToTableCommand {
	return &ServeToTableCommand{
		orderID: orderID,
	}
}

func (c *ServeToTableCommand) Execute() error {
	fmt.Printf("Подача заказа #%d на стол\n", c.orderID)
	return nil
}

func (c *ServeToTableCommand) Undo() error {
	fmt.Printf("Возврат заказа #%d с подачи\n", c.orderID)
	return nil
}

func (c *ServeToTableCommand) GetName() string {
	return fmt.Sprintf("Подача заказа #%d на стол", c.orderID)
}
