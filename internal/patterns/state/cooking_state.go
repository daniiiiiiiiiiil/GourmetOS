package state

import "fmt"

// состояние "Готовится" (процесс приготовления)
type CookingState struct{}

func NewCookingState() *CookingState {
	return &CookingState{}
}

// отправить на кухню (запрещено)
func (s *CookingState) SubmitToKitchen(order *Order) error {
	return fmt.Errorf("заказ #%d уже готовится", order.ID)
}

// начать готовить (запрещено, уже готовится)
func (s *CookingState) StartCooking(order *Order) error {
	return fmt.Errorf("заказ #%d уже готовится", order.ID)
}

// отметить как готовый (разрешено)
func (s *CookingState) MarkAsReady(order *Order) error {
	fmt.Printf("Заказ #%d готов!\n", order.ID)
	order.SetState(NewReadyState())
	return nil
}

// подать на стол (запрещено)
func (s *CookingState) ServeToTable(order *Order) error {
	return fmt.Errorf("нельзя подать: заказ ещё не готов")
}

// оплатить (запрещено)
func (s *CookingState) ProcessPayment(order *Order) error {
	return fmt.Errorf("нельзя оплатить: заказ ещё не готов")
}

// отменить заказ (запрещено, уже готовится)
func (s *CookingState) Cancel(order *Order) error {
	return fmt.Errorf("нельзя отменить: заказ уже готовится")
}

// возвращает название состояния
func (s *CookingState) GetName() string {
	return "Готовится"
}

// проверяет, можно ли перейти в другое состояние
func (s *CookingState) CanTransitionTo(target OrderState) bool {
	switch target.(type) {
	case *ReadyState:
		return true
	default:
		return false
	}
}
