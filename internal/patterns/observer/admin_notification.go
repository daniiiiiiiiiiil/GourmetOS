package observer

import "fmt"

type AdminNotifier struct {
	Name string
}

func NewAdminNotifier(name string) *AdminNotifier {
	return &AdminNotifier{Name: name}
}

func (a *AdminNotifier) Update(event Event) {
	switch event.Type {
	case OrderPaid:
		fmt.Printf("\n АДМИН %s: ЗАКАЗ #%d ОПЛАЧЕН!\n", a.Name, event.OrderID)
	}
}
