package command

import "fmt"

type AddDishCommand struct {
	orderID  int
	dishName string
	price    float64
	wasAdded bool
}

func NewAddDishCommand(orderID int, dishName string, price float64) *AddDishCommand {
	return &AddDishCommand{
		orderID:  orderID,
		dishName: dishName,
		price:    price,
	}
}

func (c *AddDishCommand) Execute() error {
	fmt.Printf("Добавление блюда '%s' в заказ #%d\n", c.dishName, c.orderID)
	c.wasAdded = true
	return nil
}

func (c *AddDishCommand) Undo() error {
	if !c.wasAdded {
		return fmt.Errorf("нельзя отменить: блюдо не было добавлено")
	}

	fmt.Printf("Отмена добавления блюда '%s' из заказа #%d\n", c.dishName, c.orderID)
	return nil
}

func (c *AddDishCommand) GetName() string {
	return fmt.Sprintf("Добавление блюда '%s' в заказ #%d", c.dishName, c.orderID)
}
