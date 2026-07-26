package facade

import (
	"GourmetOS/internal/patterns/observer"
	"fmt"
)

// упрощенный интерфейс для клиента
// заказах, командах, статусах
type CustomerFacade struct {
	subject *observer.OrderSubject
}

func NewCustomerFacade(subject *observer.OrderSubject) *CustomerFacade {
	return &CustomerFacade{
		subject: subject,
	}
}

func (c *CustomerFacade) PlaceOrder(tableID int, items []string) {
	fmt.Printf("\nКлиент: делаю заказ на столе %d\n", tableID)
	fmt.Printf("   Блюда: %v\n", items)

	// Клиент просто говорит Я хочу сделать заказ
	// Вся сложность скрыта в WaiterFacade
	// потом надо будет добавить вызов к сервису
	fmt.Println("Заказ передан официанту")
}

// клиент отслеживает статус заказа
func (c *CustomerFacade) TrackOrder(orderID int) {
	fmt.Printf("\n Клиент: проверяю статус заказа #%d\n", orderID)

	// Клиент не знает, как хранятся заказы
	// Он просто просит показать статус
	// потом надо будет добавить вызов к сервису
	fmt.Println("Статус: заказ готовится")
}

// клиент оплачивает счет
func (c *CustomerFacade) PayBill(orderID int, amount float64) {
	fmt.Printf("\nКлиент: оплачиваю счет за заказ #%d\n", orderID)
	fmt.Printf("Сумма: %.2f руб.\n", amount)
	fmt.Println("Оплата произведена")

	event := observer.NewEvent(
		observer.OrderPaid,
		orderID,
		0,
		[]string{},
	)
	c.subject.NotifyObservers(event)
}

func (c *CustomerFacade) LeaveReview(orderID int, rating int, comment string) {
	fmt.Printf("\nКлиент: оставляю отзыв на заказ #%d\n", orderID)
	fmt.Printf("Рейтинг: %d/5\n", rating)
	fmt.Printf("Комментарий: %s\n", comment)
	fmt.Println("Спасибо за отзыв!")
}
