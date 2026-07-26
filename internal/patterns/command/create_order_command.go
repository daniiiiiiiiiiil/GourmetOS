package command

import "fmt"

// OrderManager — упрощенный менеджер заказов (заменяет БД/сервисы)
type OrderManager struct {
	orders      map[int]*Order
	nextOrderID int
}

// Order — структура заказа
type Order struct {
	ID       int
	TableID  int
	Items    []string
	Total    float64
	Status   string
	IsActive bool
}

// NewOrderManager — конструктор
func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders:      make(map[int]*Order),
		nextOrderID: 1,
	}
}

// CreateOrder — создает новый заказ
func (m *OrderManager) CreateOrder(tableID int) *Order {
	order := &Order{
		ID:       m.nextOrderID,
		TableID:  tableID,
		Items:    []string{},
		Total:    0,
		Status:   "created",
		IsActive: true,
	}
	m.orders[order.ID] = order
	m.nextOrderID++
	return order
}

// GetOrder — получает заказ по ID
func (m *OrderManager) GetOrder(id int) (*Order, error) {
	order, exists := m.orders[id]
	if !exists {
		return nil, fmt.Errorf("заказ #%d не найден", id)
	}
	return order, nil
}

// CreateOrderCommand — команда создания заказа
type CreateOrderCommand struct {
	manager *OrderManager
	tableID int
	orderID int
}

// NewCreateOrderCommand — конструктор
func NewCreateOrderCommand(manager *OrderManager, tableID int) *CreateOrderCommand {
	return &CreateOrderCommand{
		manager: manager,
		tableID: tableID,
	}
}

// Execute — выполняет команду
func (c *CreateOrderCommand) Execute() error {
	fmt.Printf("   Создание заказа для стола %d\n", c.tableID)
	order := c.manager.CreateOrder(c.tableID)
	c.orderID = order.ID
	fmt.Printf("Заказ #%d создан!\n", c.orderID)
	return nil
}

// Undo — отменяет команду
func (c *CreateOrderCommand) Undo() error {
	if c.orderID == 0 {
		return fmt.Errorf("нельзя отменить: заказ не был создан")
	}
	fmt.Printf("   Отмена создания заказа #%d\n", c.orderID)
	order, err := c.manager.GetOrder(c.orderID)
	if err != nil {
		return err
	}
	order.IsActive = false
	order.Status = "cancelled"
	fmt.Printf("Заказ #%d отменен!\n", c.orderID)
	return nil
}

// GetName — название команды
func (c *CreateOrderCommand) GetName() string {
	return fmt.Sprintf("Создание заказа (стол %d)", c.tableID)
}
