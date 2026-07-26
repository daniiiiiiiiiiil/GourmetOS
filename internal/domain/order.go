package domain

import (
	"GourmetOS/internal/patterns/state"
	"errors"
	"strings"
)

type Order struct {
	OrderID        int
	TableID        int
	CustomerID     int
	WaiterID       int
	Status         string
	TotalAmount    float64
	DiscountAmount float64
	FinalAmount    float64
	PaymentMethod  string
	PaymentStatus  string
	Notes          string
	CreatedAt      string
	UpdatedAt      string
	State          state.OrderState
}

func NewOrder(
	OrderID int,
	TableID int,
	CustomerID int,
	WaiterID int,
	Status string,
	TotalAmount float64,
	DiscountAmount float64,
	FinalAmount float64,
	PaymentMethod string,
	PaymentStatus string,
	Notes string,
	CreatedAt string,
	UpdatedAt string,
) *Order {
	order := &Order{
		OrderID:        OrderID,
		TableID:        TableID,
		CustomerID:     CustomerID,
		WaiterID:       WaiterID,
		Status:         Status,
		TotalAmount:    TotalAmount,
		DiscountAmount: DiscountAmount,
		FinalAmount:    FinalAmount,
		PaymentMethod:  PaymentMethod,
		PaymentStatus:  PaymentStatus,
		Notes:          Notes,
		CreatedAt:      CreatedAt,
		UpdatedAt:      UpdatedAt,
	}
	order.setStateFromStatus(Status)
	return order
}

func (o *Order) setStateFromStatus(status string) {
	switch status {
	case "created":
		o.State = state.NewCreatedState()
	case "in_kitchen":
		o.State = state.NewInKitchenState()
	case "cooking":
		o.State = state.NewCookingState()
	case "ready":
		o.State = state.NewReadyState()
	case "served":
		o.State = state.NewServedState()
	case "paid":
		o.State = state.NewPaidState()
	default:
		o.State = state.NewCreatedState()
	}
}

func (o *Order) SetState(newState state.OrderState) {
	o.State = newState
	o.Status = newState.GetName()
}

func (o *Order) GetStateName() string {
	if o.State == nil {
		return "unknown"
	}
	return o.State.GetName()
}

func (o *Order) Validate() error {
	var errs []string

	if o.TableID <= 0 {
		errs = append(errs, "table_id обязателен")
	}
	if o.CustomerID <= 0 {
		errs = append(errs, "customer_id обязателен")
	}
	if o.WaiterID <= 0 {
		errs = append(errs, "waiter_id обязателен")
	}
	if o.TotalAmount < 0 {
		errs = append(errs, "total_amount не может быть отрицательным")
	}
	if o.DiscountAmount < 0 {
		errs = append(errs, "discount_amount не может быть отрицательным")
	}
	if o.FinalAmount < 0 {
		errs = append(errs, "final_amount не может быть отрицательным")
	}
	if o.DiscountAmount > o.TotalAmount {
		errs = append(errs, "discount_amount не может быть больше total_amount")
	}

	if len(errs) > 0 {
		return errors.New("ошибка валидации: " + strings.Join(errs, "; "))
	}
	return nil
}
