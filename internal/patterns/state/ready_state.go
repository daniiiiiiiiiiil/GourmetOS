package state

import "fmt"

// состояние "Готов" (заказ готов к подаче)
type ReadyState struct{}

func NewReadyState() *ReadyState {
	return &ReadyState{}
}

// отправить на кухню (запрещено)
func (s *ReadyState) SubmitToKitchen(order *Order) error {
	return fmt.Errorf("заказ #%d уже готов", order.ID)
}

// начать готовить (запрещено)
func (s *ReadyState) StartCooking(order *Order) error {
	return fmt.Errorf("заказ #%d уже готов", order.ID)
}

// отметить как готовый (запрещено, уже готов)
func (s *ReadyState) MarkAsReady(order *Order) error {
	return fmt.Errorf("заказ #%d уже готов", order.ID)
}

// подать на стол (разрешено)
func (s *ReadyState) ServeToTable(order *Order) error {
	fmt.Printf("Заказ #%d подан на стол %d\n", order.ID, order.TableID)
	order.SetState(NewServedState())
	return nil
}

// оплатить (разрешено, можно оплатить до подачи)
func (s *ReadyState) ProcessPayment(order *Order) error {
	fmt.Printf("Заказ #%d оплачен (до подачи)\n", order.ID)
	order.SetState(NewPaidState())
	return nil
}

// отменить заказ (запрещено, уже готов)
func (s *ReadyState) Cancel(order *Order) error {
	return fmt.Errorf("нельзя отменить: заказ уже готов")
}

func (s *ReadyState) GetName() string {
	return "Готов"
}

// проверяет, можно ли перейти в другое состояние
func (s *ReadyState) CanTransitionTo(target OrderState) bool {
	switch target.(type) {
	case *ServedState:
		return true
	case *PaidState:
		return true
	default:
		return false
	}
}
