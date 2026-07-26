package observer

import "fmt"

type WaiterNotifier struct {
	Name string // Имя официанта
}

func NewWaiterNotifier(name string) *WaiterNotifier {
	return &WaiterNotifier{Name: name}
}

func (w *WaiterNotifier) Update(event Event) {
	switch event.Type {
	case OrderCreated:
		fmt.Printf("\nОФИЦИАНТ %s: НОВЫЙ ЗАКАЗ #%d!\n", w.Name, event.OrderID)
		fmt.Printf("   Стол: %d нужно обслужить!\n", event.TableID)

	case OrderReady:
		fmt.Printf("\nОФИЦИАНТ %s: ЗАКАЗ #%d ГОТОВ!\n", w.Name, event.OrderID)
		fmt.Printf("   Несу к столу %d!\n", event.TableID)

	case OrderServed:
		fmt.Printf("\nОФИЦИАНТ %s: ЗАКАЗ #%d ПОДАН!\n", w.Name, event.OrderID)

	default:
	}
}
