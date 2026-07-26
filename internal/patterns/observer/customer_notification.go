package observer

import "fmt"

type CustomerNotifier struct {
	Phone string // Номер телефона
}

func NewCustomerNotifier(phone string) *CustomerNotifier {
	return &CustomerNotifier{Phone: phone}
}

func (c *CustomerNotifier) Update(event Event) {
	switch event.Type {
	case OrderCreated:
		fmt.Printf("\nКЛИЕНТ %s: ЗАКАЗ #%d ПРИНЯТ!\n", c.Phone, event.OrderID)
		fmt.Println("   Ожидайте, скоро начнут готовить.")

	case OrderReady:
		fmt.Printf("\nКЛИЕНТ %s: ЗАКАЗ #%d ГОТОВ!\n", c.Phone, event.OrderID)
		fmt.Println("   Официант скоро принесет!")

	case OrderServed:
		fmt.Printf("\nКЛИЕНТ %s: ПРИЯТНОГО АППЕТИТА!\n", c.Phone)

	default:
	}
}
