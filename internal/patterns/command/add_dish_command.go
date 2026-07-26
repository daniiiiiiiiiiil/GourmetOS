package command

import "fmt"

// AddDishCommand — команда добавления блюда
type AddDishCommand struct {
	manager  *OrderManager
	orderID  int
	dishName string
	price    float64
	wasAdded bool
}

// NewAddDishCommand — конструктор
func NewAddDishCommand(manager *OrderManager, orderID int, dishName string, price float64) *AddDishCommand {
	return &AddDishCommand{
		manager:  manager,
		orderID:  orderID,
		dishName: dishName,
		price:    price,
	}
}

// Execute — выполняет команду
func (c *AddDishCommand) Execute() error {
	fmt.Printf("Добавление блюда '%s' в заказ #%d\n", c.dishName, c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}
	if !order.IsActive {
		return fmt.Errorf("заказ #%d не активен", c.orderID)
	}

	order.Items = append(order.Items, c.dishName)
	order.Total += c.price
	c.wasAdded = true

	fmt.Printf("Блюдо '%s' добавлено! Итого: %.2f руб.\n", c.dishName, order.Total)
	return nil
}

// Undo — отменяет команду
func (c *AddDishCommand) Undo() error {
	if !c.wasAdded {
		return fmt.Errorf("нельзя отменить: блюдо не было добавлено")
	}

	fmt.Printf("Отмена добавления блюда '%s' из заказа #%d\n", c.dishName, c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}

	// Удаляем последнее добавленное блюдо
	for i, item := range order.Items {
		if item == c.dishName {
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
			order.Total -= c.price
			fmt.Printf("Блюдо '%s' удалено! Итого: %.2f руб.\n", c.dishName, order.Total)
			return nil
		}
	}

	return fmt.Errorf("блюдо '%s' не найдено в заказе", c.dishName)
}

// GetName — название команды
func (c *AddDishCommand) GetName() string {
	return fmt.Sprintf("Добавление блюда '%s' в заказ #%d", c.dishName, c.orderID)
}
