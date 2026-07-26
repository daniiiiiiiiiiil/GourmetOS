package command

import "fmt"

type CancelOrderCommand struct {
	orderID int
	tableID int
}

func NewCancelOrderCommand(orderID, tableID int) *CancelOrderCommand {
	return &CancelOrderCommand{
		orderID: orderID,
		tableID: tableID,
	}
}

func (c *CancelOrderCommand) Execute() error {
	fmt.Printf("Отмена заказа #%d (стол %d)\n", c.orderID, c.tableID)
	return nil
}

func (c *CancelOrderCommand) Undo() error {
	fmt.Printf("осстановление заказа #%d (стол %d)\n", c.orderID, c.tableID)
	return nil
}

func (c *CancelOrderCommand) GetName() string {
	return fmt.Sprintf("Отмена заказа #%d", c.orderID)
}
