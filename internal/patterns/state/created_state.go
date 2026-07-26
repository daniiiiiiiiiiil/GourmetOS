package state

import "fmt"

// состояние "Создан" (начальное состояние)
type CreatedState struct{}

func NewCreatedState() *CreatedState {
	return &CreatedState{}
}

// отправить на кухню (разрешено)
func (s *CreatedState) SubmitToKitchen(order *Order) error {
	fmt.Printf("Заказ #%d отправлен на кухню\n", order.ID)
	order.SetState(NewInKitchenState())
	return nil
}

// начать готовить (запрещено)
func (s *CreatedState) StartCooking(order *Order) error {
	return fmt.Errorf("нельзя начать готовить: заказ ещё не на кухне")
}

// отметить как готовый (запрещено)
func (s *CreatedState) MarkAsReady(order *Order) error {
	return fmt.Errorf("нельзя отметить как готовый: заказ ещё не готовится")
}

// подать на стол (запрещено)
func (s *CreatedState) ServeToTable(order *Order) error {
	return fmt.Errorf("нельзя подать: заказ ещё не готов")
}

// оплатить (запрещено)
func (s *CreatedState) ProcessPayment(order *Order) error {
	return fmt.Errorf("нельзя оплатить: заказ ещё не готов")
}

// отменить заказ (разрешено)
func (s *CreatedState) Cancel(order *Order) error {
	fmt.Printf("Заказ #%d отменён (статус: Создан)\n", order.ID)
	return nil
}

// возвращает название состояния
func (s *CreatedState) GetName() string {
	return "Создан"
}

// проверяет, можно ли перейти в другое состояние
func (s *CreatedState) CanTransitionTo(target OrderState) bool {
	switch target.(type) {
	case *InKitchenState:
		return true
	default:
		return false
	}
}
