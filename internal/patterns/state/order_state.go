package state

// OrderState — интерфейс для всех состояний
type OrderState interface {
	SubmitToKitchen(order *Order) error
	StartCooking(order *Order) error
	MarkAsReady(order *Order) error
	ServeToTable(order *Order) error
	ProcessPayment(order *Order) error
	Cancel(order *Order) error
	GetName() string
	CanTransitionTo(target OrderState) bool
}

// Order — структура заказа из паттерна State
// Используется ТОЛЬКО внутри паттерна, не в domain
type Order struct {
	ID       int
	TableID  int
	Items    []string
	Total    float64
	State    OrderState
	Status   string
	IsActive bool
}

func NewOrder(id, tableID int) *Order {
	order := &Order{
		ID:       id,
		TableID:  tableID,
		Items:    []string{},
		Total:    0,
		IsActive: true,
	}
	order.SetState(NewCreatedState())
	return order
}

func (o *Order) SetState(state OrderState) {
	o.State = state
	o.Status = state.GetName()
}
