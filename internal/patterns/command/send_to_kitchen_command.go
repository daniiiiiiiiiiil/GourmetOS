package command

import "fmt"

// SendToKitchenCommand — команда отправки на кухню
type SendToKitchenCommand struct {
	manager *OrderManager
	orderID int
	wasSent bool
}

// NewSendToKitchenCommand — конструктор
func NewSendToKitchenCommand(manager *OrderManager, orderID int) *SendToKitchenCommand {
	return &SendToKitchenCommand{
		manager: manager,
		orderID: orderID,
	}
}

// Execute — выполняет команду
func (c *SendToKitchenCommand) Execute() error {
	fmt.Printf("   Отправка заказа #%d на кухню\n", c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}
	if !order.IsActive {
		return fmt.Errorf("заказ #%d не активен", c.orderID)
	}
	if order.Status == "in_kitchen" {
		return fmt.Errorf("заказ #%d уже на кухне", c.orderID)
	}

	order.Status = "in_kitchen"
	c.wasSent = true

	fmt.Printf("Заказ #%d отправлен на кухню!\n", c.orderID)
	fmt.Printf("Блюда: %v\n", order.Items)
	return nil
}

// Undo — отменяет команду
func (c *SendToKitchenCommand) Undo() error {
	if !c.wasSent {
		return fmt.Errorf("нельзя отменить: заказ не был отправлен")
	}

	fmt.Printf("   Отмена отправки заказа #%d на кухню\n", c.orderID)

	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}

	order.Status = "created"
	fmt.Printf("Заказ #%d возвращен с кухни!\n", c.orderID)
	return nil
}

// GetName — название команды
func (c *SendToKitchenCommand) GetName() string {
	return fmt.Sprintf("Отправка заказа #%d на кухню", c.orderID)
}
