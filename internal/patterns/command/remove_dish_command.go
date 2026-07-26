package command

import "fmt"

// RemoveDishCommand — команда удаления блюда из заказа
type RemoveDishCommand struct {
	orderID  int
	dishName string
	price    float64
	wasAdded bool
}

func NewRemoveDishCommand(orderID int, dishName string, price float64) *RemoveDishCommand {
	return &RemoveDishCommand{
		orderID:  orderID,
		dishName: dishName,
		price:    price,
	}
}

func (c *RemoveDishCommand) Execute() error {
	fmt.Printf("Удаление блюда '%s' из заказа #%d\n", c.dishName, c.orderID)
	return nil
}

func (c *RemoveDishCommand) Undo() error {
	fmt.Printf("Восстановление блюда '%s' в заказе #%d\n", c.dishName, c.orderID)
	return nil
}

func (c *RemoveDishCommand) GetName() string {
	return fmt.Sprintf("Удаление блюда '%s' из заказа #%d", c.dishName, c.orderID)
}
