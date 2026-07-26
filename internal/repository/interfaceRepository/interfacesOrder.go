package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, conn *pgx.Conn, order domain.Order) (*domain.Order, error)
	GetByIdOrder(ctx context.Context, conn *pgx.Conn, orderID int) (*domain.Order, error)
	UpdateOrder(ctx context.Context, conn *pgx.Conn, order domain.Order, id int) (*domain.Order, error)
	DeleteOrder(ctx context.Context, conn *pgx.Conn, orderID int) error
	GetAllOrders(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, error)
	GetByStatusOrder(ctx context.Context, conn *pgx.Conn, status string, limit, offset int) ([]domain.Order, error)
	GetByTableOrder(ctx context.Context, conn *pgx.Conn, tableID, limit, offset int) ([]domain.Order, error)
	GetByCustomerOrder(ctx context.Context, conn *pgx.Conn, customerID, limit, offset int) ([]domain.Order, error)
	GetByWaiterOrder(ctx context.Context, conn *pgx.Conn, waiterID, limit, offset int) ([]domain.Order, error)
	GetByActiveOrder(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, error)
	UpdateStatusOrder(ctx context.Context, conn *pgx.Conn, orderID int, status string) error
	UpdateTotalOrder(ctx context.Context, conn *pgx.Conn, orderID int, total float64) error
	CountByStatus(ctx context.Context, conn *pgx.Conn, status string) (int, error)
	AddOrderItem(ctx context.Context, conn *pgx.Conn, orderItem domain.OrderItem) error
	GetOrderItems(ctx context.Context, conn *pgx.Conn, orderID int) ([]domain.OrderItem, error)
	RemoveOrderItem(ctx context.Context, conn *pgx.Conn, orderID, dishID int) error
	ClearOrderItems(ctx context.Context, conn *pgx.Conn, orderID int) error
	UpdateOrderItemQuantity(ctx context.Context, conn *pgx.Conn, orderID, dishID int, quantity int) error
	UpdateDiscount(ctx context.Context, conn *pgx.Conn, orderID int, discountAmount, finalAmount float64) error
}
