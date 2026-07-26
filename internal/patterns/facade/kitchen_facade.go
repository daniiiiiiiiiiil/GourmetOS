package facade

import (
	"GourmetOS/internal/patterns/command"
	"GourmetOS/internal/patterns/observer"
	"fmt"
)

// упрощенный интерфейс для кухни
// Повара не знают о командах и заказах, они просто готовят
type KitchenFacade struct {
	manager *command.OrderManager
	subject *observer.OrderSubject
}

func NewKitchenFacade(
	manager *command.OrderManager,
	subject *observer.OrderSubject,
) *KitchenFacade {
	return &KitchenFacade{
		manager: manager,
		subject: subject,
	}
}

// кухня получает заказ (уведомление от Observer)
func (k *KitchenFacade) ReceiveOrder(orderID int) {
	fmt.Printf("\nКухня: получен заказ #%d\n", orderID)

	order, err := k.manager.GetOrder(orderID)
	if err != nil {
		fmt.Printf("Ошибка: заказ не найден\n")
		return
	}

	fmt.Printf("Блюда к приготовлению: %v\n", order.Items)
	fmt.Println("Начинаем готовить!")

	// Меняем статус на "готовится"
	order.Status = "cooking"
}

// кухня сообщает, что заказ готов
func (k *KitchenFacade) MarkAsReady(orderID int) {
	fmt.Printf("\n Кухня: заказ #%d готов!\n", orderID)

	order, err := k.manager.GetOrder(orderID)
	if err != nil {
		fmt.Printf("Ошибка: заказ не найден\n")
		return
	}

	order.Status = "ready"

	// Оповещаем всех, что заказ готов
	event := observer.NewEvent(
		observer.OrderReady,
		order.ID,
		order.TableID,
		order.Items,
	)
	k.subject.NotifyObservers(event)

	fmt.Printf("Заказ #%d готов к выдаче!\n", orderID)
}

// кухня показывает очередь заказов
func (k *KitchenFacade) ShowQueue() {
	fmt.Println("\n Очередь заказов на кухне:")

	//тут также сервис как и везде
	fmt.Println("Заказ #1 - готовится (15 мин)")
	fmt.Println("Заказ #2 - в очереди (10 мин)")
	fmt.Println("Заказ #3 - в очереди (20 мин)")
}

// кухня показывает время приготовления
func (k *KitchenFacade) CookingTime(dishName string) int {
	fmt.Printf("\n Время приготовления блюда '%s'\n", dishName)

	// В реальном проекте здесь был бы поиск в БД
	// подключение к репозитории
	times := map[string]int{
		"Пицца Маргарита": 15,
		"Паста Карбонара": 20,
		"Окономияки":      25,
		"Рамен":           30,
	}
	if time, ok := times[dishName]; ok {
		fmt.Printf("%d минут\n", time)
		return time
	}
	fmt.Println("15 минут (стандартное время)")
	return 15
}
