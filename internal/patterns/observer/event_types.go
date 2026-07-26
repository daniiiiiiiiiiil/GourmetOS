package observer

import "time"

// EventType — тип события (просто строка)
type EventType string

// Все возможные события в ресторане
const (
	OrderCreated   EventType = "order_created"   // Заказ создан
	OrderReady     EventType = "order_ready"     // Заказ готов
	OrderServed    EventType = "order_served"    // Заказ подан
	OrderPaid      EventType = "order_paid"      // Заказ оплачен
	OrderCancelled EventType = "order_cancelled" // Заказ отменен
)

type Event struct {
	Type      EventType // Что случилось?
	OrderID   int       // Какой заказ?
	TableID   int       // За каким столом?
	Items     []string  // Какие блюда?
	Timestamp time.Time // Когда случилось?
}

func NewEvent(eventType EventType, orderID, tableID int, items []string) Event {
	return Event{
		Type:      eventType,
		OrderID:   orderID,
		TableID:   tableID,
		Items:     items,
		Timestamp: time.Now(),
	}
}
