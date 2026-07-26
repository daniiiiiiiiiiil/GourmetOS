package observer

import "fmt"

// KitchenDisplay — экран на кухне (подписчик)
// Реализует интерфейс Observer
type KitchenDisplay struct {
	Name string // Название кухни
}

// NewKitchenDisplay — конструктор
func NewKitchenDisplay(name string) *KitchenDisplay {
	return &KitchenDisplay{Name: name}
}

// Update — вызывается при наступлении события
func (k *KitchenDisplay) Update(event Event) {
	switch event.Type {
	case OrderCreated:
		fmt.Printf("КУХНЯ [%s]: НОВЫЙ ЗАКАЗ #%d!\n", k.Name, event.OrderID)
		fmt.Printf("Стол: %d\n", event.TableID)
		fmt.Printf("Блюда: %v\n", event.Items)
		fmt.Println("Начинаем готовить!")

	case OrderReady:
		fmt.Printf("\nКУХНЯ [%s]: ЗАКАЗ #%d ГОТОВ!\n", k.Name, event.OrderID)
		fmt.Println("Зовем официанта!")

	default:

	}
}
