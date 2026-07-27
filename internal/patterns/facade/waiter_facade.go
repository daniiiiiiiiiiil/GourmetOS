package facade

import (
	"GourmetOS/internal/patterns/command"
	"GourmetOS/internal/patterns/observer"
	"fmt"
)

// упрощенный интерфейс для официанта
// Официант не знает о сложностях: командах, наблюдателях, заказах
type WaiterFacade struct {
	manager *command.OrderManager   `json:"manager,omitempty"`
	invoker *command.CommandInvoker `json:"invoker,omitempty"`
	subject *observer.OrderSubject  `json:"subject,omitempty"`
}

// конструктор фасада для официанта
func NewWaiterFacade(
	manager *command.OrderManager,
	invoker *command.CommandInvoker,
	subject *observer.OrderSubject,
) *WaiterFacade {
	return &WaiterFacade{
		manager: manager,
		invoker: invoker,
		subject: subject,
	}
}

// официант создает заказ (все сложности скрыты)
func (w *WaiterFacade) CreateOrder(tableID int) (int, error) {
	fmt.Printf("\n Официант: создаю заказ для стола %d\n", tableID)

	// Создаем команду
	cmd := command.NewCreateOrderCommand(w.manager, tableID)

	// Выполняем команду через инвокер
	err := w.invoker.Execute(cmd)
	if err != nil {
		return 0, fmt.Errorf("не удалось создать заказ: %w", err)
	}

	//  Получаем созданный заказ
	order, err := w.manager.GetOrder(1)
	if err != nil {
		return 0, err
	}

	// Оповещаем всех о новом заказе (Observer)
	event := observer.NewEvent(
		observer.OrderCreated,
		order.ID,
		order.TableID,
		order.Items,
	)
	w.subject.NotifyObservers(event)

	fmt.Printf("Заказ #%d создан!\n", order.ID)
	return order.ID, nil
}

// официант добавляет блюдо в заказ
func (w *WaiterFacade) AddDishToOrder(orderID int, dishName string, price float64) error {
	fmt.Printf("Официант: добавляю блюдо '%s' в заказ #%d\n", dishName, orderID)

	cmd := command.NewAddDishCommand(orderID, dishName, price)
	err := w.invoker.Execute(cmd)
	if err != nil {
		return fmt.Errorf("не удалось добавить блюдо: %w", err)
	}

	return nil
}

// официант отправляет заказ на кухню
func (w *WaiterFacade) SendToKitchen(orderID int) error {
	fmt.Printf("Официант: отправляю заказ #%d на кухню\n", orderID)

	cmd := command.NewSendToKitchenCommand(w.manager, orderID)
	err := w.invoker.Execute(cmd)
	if err != nil {
		return fmt.Errorf("не удалось отправить заказ на кухню: %w", err)
	}

	// Оповещаем кухню о новом заказе
	order, _ := w.manager.GetOrder(orderID)
	event := observer.NewEvent(
		observer.OrderReady,
		order.ID,
		order.TableID,
		order.Items,
	)
	w.subject.NotifyObservers(event)

	return nil
}

// официант завершает заказ (выдан клиенту)
func (w *WaiterFacade) CompleteOrder(orderID int) error {
	fmt.Printf("Официант: завершаю заказ #%d\n", orderID)

	cmd := command.NewCompleteOrderCommand(w.manager, orderID)
	err := w.invoker.Execute(cmd)
	if err != nil {
		return fmt.Errorf("не удалось завершить заказ: %w", err)
	}

	// Оповещаем клиента
	order, _ := w.manager.GetOrder(orderID)
	event := observer.NewEvent(
		observer.OrderServed,
		order.ID,
		order.TableID,
		order.Items,
	)
	w.subject.NotifyObservers(event)

	return nil
}

// официант отменяет заказ
func (w *WaiterFacade) CancelOrder(orderID, tableID int) error {
	fmt.Printf("Официант: отменяю заказ #%d\n", orderID)

	cmd := command.NewCancelOrderCommand(orderID, tableID)
	err := w.invoker.Execute(cmd)
	if err != nil {
		return fmt.Errorf("не удалось отменить заказ: %w", err)
	}

	return nil
}

// официант отменяет последнее действие
func (w *WaiterFacade) Undo() error {
	fmt.Println("Официант: отмена последнего действия")
	return w.invoker.Undo()
}

// официант повторяет отмененное действие
func (w *WaiterFacade) Redo() error {
	fmt.Println("Официант: повтор отмененного действия")
	return w.invoker.Redo()
}

// показывает текущий заказ
func (w *WaiterFacade) ShowOrder(orderID int) error {
	order, err := w.manager.GetOrder(orderID)
	if err != nil {
		return err
	}

	fmt.Printf("\n ЗАКАЗ #%d:\n", order.ID)
	fmt.Printf("   Стол: %d\n", order.TableID)
	fmt.Printf("   Блюда: %v\n", order.Items)
	fmt.Printf("   Итого: %.2f руб.\n", order.Total)
	fmt.Printf("   Статус: %s\n", order.Status)
	return nil
}
