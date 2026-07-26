package dto

import (
	"GourmetOS/internal/domain"
	"GourmetOS/pkg/errors"
	"GourmetOS/pkg/pagination"
)

type KitchenActionRequest struct {
	OrderID int `json:"order_id" binding:"required"`
}

func (r *KitchenActionRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.OrderID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "order_id",
			Message: "ID заказа не может быть меньше или равен 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

type KitchenStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=in_kitchen cooking ready"`
}

func (r *KitchenStatusRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.Status == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "status",
			Message: "Статус не может быть пустым",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

type KitchenQueueItem struct {
	OrderID     int      `json:"order_id"`
	TableID     int      `json:"table_id"`
	Status      string   `json:"status"`
	Items       []string `json:"items"`
	Total       float64  `json:"total"`
	CookingTime int      `json:"cooking_time"`
}

type KitchenQueueResponse struct {
	Orders     []KitchenQueueItem    `json:"orders"`
	Pagination pagination.Pagination `json:"pagination"`
}

func NewKitchenQueueResponse(orders []domain.Order, total, limit, offset int) KitchenQueueResponse {
	resp := KitchenQueueResponse{
		Orders:     make([]KitchenQueueItem, 0, len(orders)),
		Pagination: pagination.NewPagination(total, limit, offset),
	}
	for _, order := range orders {
		resp.Orders = append(resp.Orders, KitchenQueueItem{
			OrderID: order.OrderID,
			TableID: order.TableID,
			Status:  order.Status,
			Total:   order.TotalAmount,
		})
	}
	return resp
}

type KitchenStatusResponse struct {
	InQueue   int    `json:"in_queue"`
	Cooking   int    `json:"cooking"`
	Total     int    `json:"total"`
	Timestamp string `json:"timestamp"`
}

type KitchenResponse struct {
	OrderID int    `json:"order_id"`
	TableID int    `json:"table_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func FromDomainKitchenResponse(order domain.Order, message string) KitchenResponse {
	return KitchenResponse{
		OrderID: order.OrderID,
		TableID: order.TableID,
		Status:  order.Status,
		Message: message,
	}
}
