package state

import "fmt"

// состояние "Подан" (заказ подан клиенту)
type ServedState struct{}

func NewServedState() *ServedState {
	return &ServedState{}
}

// отправить на кухню (запрещено)
func (s *ServedState) SubmitToKitchen(order *Order) error {
	return fmt.Errorf("заказ #%d уже подан", order.ID)
}

// начать готовить (запрещено)
func (s *ServedState) StartCooking(order *Order) error {
	return fmt.Errorf("заказ #%d уже подан", order.ID)
}

// отметить как готовый (запрещено)
func (s *ServedState) MarkAsReady(order *Order) error {
	return fmt.Errorf("заказ #%d уже подан", order.ID)
}

// подать на стол (запрещено, уже подан)
func (s *ServedState) ServeToTable(order *Order) error {
	return fmt.Errorf("заказ #%d уже подан", order.ID)
}

// оплатить (разрешено)
func (s *ServedState) ProcessPayment(order *Order) error {
	fmt.Printf("Заказ #%d оплачен (после подачи)\n", order.ID)
	order.SetState(NewPaidState())
	return nil
}

// отменить заказ (запрещено, уже подан)
func (s *ServedState) Cancel(order *Order) error {
	return fmt.Errorf("нельзя отменить: заказ уже подан")
}

// возвращает название состояния
func (s *ServedState) GetName() string {
	return "Подан"
}

// проверяет, можно ли перейти в другое состояние
func (s *ServedState) CanTransitionTo(target OrderState) bool {
	switch target.(type) {
	case *PaidState:
		return true
	default:
		return false
	}
}
