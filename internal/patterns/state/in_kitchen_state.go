package state

import "fmt"

// состояние "На кухне" (ожидает готовки)
type InKitchenState struct{}

func NewInKitchenState() *InKitchenState {
	return &InKitchenState{}
}

// отправить на кухню (запрещено, уже на кухне)
func (s *InKitchenState) SubmitToKitchen(order *Order) error {
	return fmt.Errorf("заказ #%d уже на кухне", order.ID)
}

// начать готовить (разрешено)
func (s *InKitchenState) StartCooking(order *Order) error {
	fmt.Printf("Заказ #%d начали готовить\n", order.ID)
	order.SetState(NewCookingState())
	return nil
}

// отметить как готовый (запрещено)
func (s *InKitchenState) MarkAsReady(order *Order) error {
	return fmt.Errorf("нельзя отметить как готовый: заказ ещё не готовится")
}

// подать на стол (запрещено)
func (s *InKitchenState) ServeToTable(order *Order) error {
	return fmt.Errorf("нельзя подать: заказ ещё не готов")
}

// оплатить (запрещено)
func (s *InKitchenState) ProcessPayment(order *Order) error {
	return fmt.Errorf("нельзя оплатить: заказ ещё не готов")
}

// отменить заказ (разрешено, но редко)
func (s *InKitchenState) Cancel(order *Order) error {
	fmt.Printf("Заказ #%d отменён (статус: На кухне)\n", order.ID)
	return nil
}

func (s *InKitchenState) GetName() string {
	return "На кухне"
}

// проверяет, можно ли перейти в другое состояние
func (s *InKitchenState) CanTransitionTo(target OrderState) bool {
	switch target.(type) {
	case *CookingState:
		return true
	default:
		return false
	}
}
