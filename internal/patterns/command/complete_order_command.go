package command

import "fmt"

// CompleteOrderCommand — команда завершения заказа
type CompleteOrderCommand struct {
	manager      *OrderManager
	orderID      int
	wasCompleted bool
}

// NewCompleteOrderCommand — конструктор
func NewCompleteOrderCommand(manager *OrderManager, orderID int) *CompleteOrderCommand {
	return &CompleteOrderCommand{
		manager: manager,
		orderID: orderID,
	}
}

// Execute — выполняет команду
func (c *CompleteOrderCommand) Execute() error {
	fmt.Printf("   Завершение заказа #%d\n", c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}
	if !order.IsActive {
		return fmt.Errorf("заказ #%d не активен", c.orderID)
	}
	if order.Status == "completed" {
		return fmt.Errorf("заказ #%d уже завершен", c.orderID)
	}

	order.Status = "completed"
	order.IsActive = false
	c.wasCompleted = true

	fmt.Printf("Заказ #%d завершен! Итого: %.2f руб.\n", c.orderID, order.Total)
	return nil
}

// Undo — отменяет команду
func (c *CompleteOrderCommand) Undo() error {
	if !c.wasCompleted {
		return fmt.Errorf("нельзя отменить: заказ не был завершен")
	}

	fmt.Printf("Отмена завершения заказа #%d\n", c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}

	order.Status = "created"
	order.IsActive = true
	c.wasCompleted = false

	fmt.Printf("Заказ #%d возвращен в работу!\n", c.orderID)
	return nil
}

// GetName — название команды
func (c *CompleteOrderCommand) GetName() string {
	return fmt.Sprintf("Завершение заказа #%d", c.orderID)
}
