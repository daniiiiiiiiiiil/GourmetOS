package state

import "fmt"

// PaidState — состояние "Оплачен" (финальное состояние)
type PaidState struct{}

// NewPaidState — конструктор состояния
func NewPaidState() *PaidState {
	return &PaidState{}
}

// SubmitToKitchen — отправить на кухню (запрещено)
func (s *PaidState) SubmitToKitchen(order *Order) error {
	return fmt.Errorf("заказ #%d уже оплачен", order.ID)
}

// StartCooking — начать готовить (запрещено)
func (s *PaidState) StartCooking(order *Order) error {
	return fmt.Errorf("заказ #%d уже оплачен", order.ID)
}

// отметить как готовый (запрещено)
func (s *PaidState) MarkAsReady(order *Order) error {
	return fmt.Errorf("заказ #%d уже оплачен", order.ID)
}

// подать на стол (запрещено)
func (s *PaidState) ServeToTable(order *Order) error {
	return fmt.Errorf("заказ #%d уже оплачен", order.ID)
}

// оплатить (запрещено, уже оплачен)
func (s *PaidState) ProcessPayment(order *Order) error {
	return fmt.Errorf("заказ #%d уже оплачен", order.ID)
}

// отменить заказ (запрещено, финальное состояние)
func (s *PaidState) Cancel(order *Order) error {
	return fmt.Errorf("нельзя отменить: заказ уже оплачен")
}

// возвращает название состояния
func (s *PaidState) GetName() string {
	return "Оплачен"
}

// проверяет, можно ли перейти в другое состояние
func (s *PaidState) CanTransitionTo(target OrderState) bool {
	// Финальное состояние — переходов нет
	return false
}
