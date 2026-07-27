package dto

import (
	"GourmetOS/internal/domain"
	"GourmetOS/pkg/errors"
	"GourmetOS/pkg/pagination"
	"time"
)

type CreateOrderRequest struct {
	TableID    int            `json:"table_id" binding:"required"`
	CustomerID int            `json:"customer_id"`
	WaiterID   int            `json:"waiter_id"`
	Items      []OrderItemReq `json:"items" binding:"required"`
	Notes      string         `json:"notes"`
}

func (r *CreateOrderRequest) ToDomain() *CreateOrderRequest {
	return &CreateOrderRequest{
		TableID:    r.TableID,
		CustomerID: r.CustomerID,
		WaiterID:   r.WaiterID,
		Items:      r.Items,
		Notes:      r.Notes,
	}
}
func (r *CreateOrderRequest) Validate() error {
	var errs errors.ValidationErrors

	if r.TableID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "table_id",
			Message: "ID стола не может быть меньше или равен 0",
		})
	}

	if len(r.Items) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "items",
			Message: "Список блюд не может быть пустым",
		})
	}

	for _, item := range r.Items {
		if item.DishID <= 0 {
			errs = append(errs, errors.ValidationError{
				Field:   "items.dish_id",
				Message: "ID блюда не может быть меньше или равен 0",
			})
		}
		if item.Quantity <= 0 {
			errs = append(errs, errors.ValidationError{
				Field:   "items.quantity",
				Message: "Количество не может быть меньше или равен 0",
			})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type OrderItemReq struct {
	DishID   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (r *OrderItemReq) ToDomain(order domain.Order) *OrderItemReq {
	return &OrderItemReq{
		DishID:   r.DishID,
		Quantity: r.Quantity,
	}
}

type AddDishRequest struct {
	DishID   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (r *AddDishRequest) ToDomain(order domain.Order) *AddDishRequest {
	return &AddDishRequest{
		DishID:   r.DishID,
		Quantity: r.Quantity,
	}
}

type RemoveDishRequest struct {
	OrderID  int `json:"order_id" binding:"required"`
	DishID   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity"`
}

func (r *RemoveDishRequest) ToDomain(order domain.Order) *RemoveDishRequest {
	return &RemoveDishRequest{
		OrderID:  r.OrderID,
		DishID:   r.DishID,
		Quantity: r.Quantity,
	}
}

type ProcessPaymentRequest struct {
	Method     string `json:"method" binding:"required,oneof=cash card crypto"`
	CardNumber string `json:"card_number"`
	Expiry     string `json:"expiry"`
	CVV        string `json:"cvv"`
	Wallet     string `json:"wallet"`
}

func (r *ProcessPaymentRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.Method == "" || len(r.Method) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "method",
			Message: "Метод не может быть пустым",
		})
	}
	if r.CardNumber == "" || len(r.CardNumber) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "card_number",
			Message: "Номер карты не может быть пустым",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func (r *ProcessPaymentRequest) ToDomain(order domain.Order) *ProcessPaymentRequest {
	return &ProcessPaymentRequest{
		Method:     r.Method,
		CardNumber: r.CardNumber,
		Expiry:     r.Expiry,
		CVV:        r.CVV,
		Wallet:     r.Wallet,
	}
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (r *UpdateStatusRequest) ToDomain(order domain.Order) *UpdateStatusRequest {
	return &UpdateStatusRequest{
		Status: r.Status,
	}
}

type OrderResponse struct {
	OrderID        int             `json:"order_id"`
	TableID        int             `json:"table_id"`
	CustomerID     int             `json:"customer_id"`
	WaiterID       int             `json:"waiter_id"`
	Status         string          `json:"status"`
	TotalAmount    float64         `json:"total_amount"`
	DiscountAmount float64         `json:"discount_amount"`
	FinalAmount    float64         `json:"final_amount"`
	PaymentMethod  string          `json:"payment_method"`
	PaymentStatus  string          `json:"payment_status"`
	Notes          *string         `json:"notes"`
	Items          []OrderItemResp `json:"items"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func OrderResponseFromDomain(order domain.Order) *OrderResponse {
	return &OrderResponse{
		OrderID:        order.OrderID,
		TableID:        order.TableID,
		CustomerID:     order.CustomerID,
		WaiterID:       order.WaiterID,
		Status:         order.Status,
		TotalAmount:    order.TotalAmount,
		DiscountAmount: order.DiscountAmount,
		FinalAmount:    order.FinalAmount,
		PaymentMethod:  order.PaymentMethod,
		PaymentStatus:  order.PaymentStatus,
		Notes:          order.Notes,
		CreatedAt:      order.CreatedAt,
		UpdatedAt:      order.UpdatedAt,
	}
}

type OrderItemResp struct {
	DishID   int     `json:"dish_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderListResponse struct {
	Orders     []OrderResponse       `json:"orders"`
	Pagination pagination.Pagination `json:"pagination"`
}

func NewOrderListResponse(orders []domain.Order, total, limit, offset int) *OrderListResponse {
	resp := &OrderListResponse{
		Orders:     make([]OrderResponse, 0, len(orders)),
		Pagination: pagination.NewPagination(total, limit, offset),
	}
	for _, order := range orders {
		resp.Orders = append(resp.Orders, *OrderResponseFromDomain(order))
	}
	return resp
}
