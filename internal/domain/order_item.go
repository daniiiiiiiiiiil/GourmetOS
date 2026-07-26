package domain

import "time"

type OrderItem struct {
	OrderItemID int       `json:"order_item_id"`
	OrderID     int       `json:"order_id"`
	DishID      int       `json:"dish_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
