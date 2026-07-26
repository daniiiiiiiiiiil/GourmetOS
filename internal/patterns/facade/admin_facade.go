package facade

import (
	"GourmetOS/internal/patterns/command"
	"fmt"
)

// AdminFacade — упрощенный интерфейс для администратора
// Админ управляет всем рестораном через простые методы
type AdminFacade struct {
	manager *command.OrderManager
	waiter  *WaiterFacade
	kitchen *KitchenFacade
}

// NewAdminFacade — конструктор фасада для администратора
func NewAdminFacade(
	manager *command.OrderManager,
	waiter *WaiterFacade,
	kitchen *KitchenFacade,
) *AdminFacade {
	return &AdminFacade{
		manager: manager,
		waiter:  waiter,
		kitchen: kitchen,
	}
}

// ПОКА ПРИНТЫ ПОТОМ НАДО ТАКЖЕ БУДЕТ ПОЛУЧАТЬ ДАННЫЕ ИЗ РЕПОЗИТОРИЯ
func (a *AdminFacade) Dashboard() {
	fmt.Println("\nДАШБОРД РЕСТОРАНА")
	fmt.Println("======================")
	fmt.Println("   Активных заказов: 4")
	fmt.Println("   Свободных столов: 6")
	fmt.Println("   Сегодняшняя выручка: 12 450 руб.")
	fmt.Println("   Активных сотрудников: 7")
	fmt.Println("======================")
}

// администратор видит все заказы
func (a *AdminFacade) ViewAllOrders() {
	fmt.Println("\n ВСЕ ЗАКАЗЫ:")
	//ЗАПРОС К БД
	fmt.Println("   #1 - стол 3 - 830 руб. - completed")
	fmt.Println("   #2 - стол 5 - 450 руб. - cooking")
	fmt.Println("   #3 - стол 1 - 1200 руб. - in_kitchen")
	fmt.Println("   #4 - стол 7 - 650 руб. - created")
}

// администратор управляет сотрудниками
func (a *AdminFacade) ManageEmployee(action string, name string) {
	fmt.Printf("\n Администратор: %s сотрудника %s\n", action, name)

	switch action {
	case "hire":
		fmt.Println("Сотрудник нанят")
	case "fire":
		fmt.Println("Сотрудник уволен")
	case "promote":
		fmt.Println("Сотрудник повышен")
	default:
		fmt.Println("Действие выполнено")
	}
}

// администратор смотрит отчеты
func (a *AdminFacade) ViewReports(period string) {
	fmt.Printf("\nОТЧЕТЫ за %s:\n", period)
	fmt.Println("Популярные блюда: Пицца, Паста")
	fmt.Println("Средний чек: 850 руб.")
	fmt.Println("Количество заказов: 120")
	fmt.Println("Выручка: 102 000 руб.")
}

// администратор экстренно отменяет заказ
func (a *AdminFacade) EmergencyCancel(orderID int) {
	fmt.Printf("\nАдминистратор: экстренная отмена заказа #%d\n", orderID)

	cmd := command.NewCancelOrderCommand(a.manager, orderID)
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Заказ #%d отменен!\n", orderID)
}

// администратор устанавливает скидку
func (a *AdminFacade) SetDiscount(orderID int, percent float64) {
	fmt.Printf("\nАдминистратор: скидка %.0f%% на заказ #%d\n", percent, orderID)

	order, err := a.manager.GetOrder(orderID)
	if err != nil {
		fmt.Printf("Ошибка: заказ не найден\n")
		return
	}

	discount := order.Total * (percent / 100)
	order.Total -= discount
	fmt.Printf("Скидка применена! Сумма: %.2f руб.\n", order.Total)
}
