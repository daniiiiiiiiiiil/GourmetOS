package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	Pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{Pool: pool}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, conn *pgx.Conn, order domain.Order) (*domain.Order, error) {
	sqlQuery := `
	INSERT INTO orders(table_id, customer_id, waiter_id, status, total_amount, 
					   discount_amount, final_amount, payment_method, payment_status, notes, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING order_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		order.TableID,
		order.CustomerID,
		order.WaiterID,
		order.Status,
		order.TotalAmount,
		order.DiscountAmount,
		order.FinalAmount,
		order.PaymentMethod,
		order.PaymentStatus,
		order.Notes,
		time.Now(),
	).Scan(&order.OrderID, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetByIdOrder(ctx context.Context, conn *pgx.Conn, orderID int) (*domain.Order, error) {
	sqlQuery := `
	SELECT *
	FROM orders
	WHERE order_id = $1
`
	rows, err := conn.Query(ctx, sqlQuery, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	order, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, conn *pgx.Conn, order domain.Order, id int) (*domain.Order, error) {
	sqlQuery := `
	UPDATE orders
	SET table_id = $1, customer_id = $2, waiter_id = $3, status = $4, 
		total_amount = $5, discount_amount = $6, final_amount = $7, 
		payment_method = $8, payment_status = $9, notes = $10, updated_at = $11
	WHERE order_id = $12
	RETURNING updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		order.TableID, order.CustomerID, order.WaiterID,
		order.Status, order.TotalAmount, order.DiscountAmount, order.FinalAmount,
		order.PaymentMethod, order.PaymentStatus, order.Notes, time.Now(), id,
	).Scan(&order.UpdatedAt)

	if err != nil {
		return nil, err
	}
	order.OrderID = id
	return &order, nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, conn *pgx.Conn, orderID int) error {
	sqlQuery := `
		DELETE FROM orders
		WHERE order_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, orderID)
	return err
}

func (r *OrderRepository) GetAllOrders(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
	SELECT * FROM orders
	LIMIT $1 OFFSET $2
	`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil

}

func (r *OrderRepository) GetByStatusOrder(ctx context.Context, conn *pgx.Conn, status string, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
			SELECT *
			FROM orders
			WHERE status = $1
			LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetByTableOrder(ctx context.Context, conn *pgx.Conn, tableID, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
	SELECT *
	FROM orders
	WHERE table_id = $1
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, tableID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetByCustomerOrder(ctx context.Context, conn *pgx.Conn, customerID, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
		SELECT *
		FROM orders
		WHERE customer_id = $1
		LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, customerID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetByWaiterOrder(ctx context.Context, conn *pgx.Conn, waiterID, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
		SELECT *
		FROM orders
		WHERE waiter_id = $1
		LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, waiterID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetByActiveOrder(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, error) {
	sqlQuery := `
		SELECT *
		FROM orders
		WHERE status = 'active'
		LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Order])
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateStatusOrder(ctx context.Context, conn *pgx.Conn, orderID int, status string) error {
	sqlQuery := `UPDATE orders SET status = $1 WHERE order_id = $2`
	_, err := conn.Exec(ctx, sqlQuery, status, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) UpdateTotalOrder(ctx context.Context, conn *pgx.Conn, orderID int, total float64) error {
	sqlQuery := `UPDATE orders SET total_amount = $1 WHERE order_id = $2`
	_, err := conn.Exec(ctx, sqlQuery, total, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) CountByStatus(ctx context.Context, conn *pgx.Conn, status string) (int, error) {
	var count int
	err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM orders WHERE status = $1", status).Scan(&count)
	return count, err
}

func (r *OrderRepository) AddOrderItem(ctx context.Context, conn *pgx.Conn, orderItem domain.OrderItem) error {
	sqlQuery := `
	INSERT INTO order_items (order_id, dish_id, quantity, price)
	VALUES ($1, $2, $3, $4)`
	_, err := conn.Exec(ctx, sqlQuery,
		orderItem.OrderID,
		orderItem.DishID,
		orderItem.Quantity,
		orderItem.Price,
	)
	return err
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, conn *pgx.Conn, orderID int) ([]domain.OrderItem, error) {
	sqlQuery := `
	SELECT order_item_id, order_id, dish_id, quantity, price, created_at, updated_at
	FROM order_items
	WHERE order_id = $1`
	rows, err := conn.Query(ctx, sqlQuery, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		if err := rows.Scan(
			&item.OrderItemID,
			&item.OrderID,
			&item.DishID,
			&item.Quantity,
			&item.Price,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *OrderRepository) RemoveOrderItem(ctx context.Context, conn *pgx.Conn, orderID, dishID int) error {
	sqlQuery := `DELETE FROM order_items WHERE order_id = $1 AND dish_id = $2`
	_, err := conn.Exec(ctx, sqlQuery, orderID, dishID)
	return err
}

func (r *OrderRepository) ClearOrderItems(ctx context.Context, conn *pgx.Conn, orderID int) error {
	sqlQuery := `DELETE FROM order_items WHERE order_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, orderID)
	return err
}

func (r *OrderRepository) UpdateOrderItemQuantity(ctx context.Context, conn *pgx.Conn, orderID, dishID int, quantity int) error {
	sqlQuery := `
	UPDATE order_items
	SET quantity = $1, updated_at = NOW()
	WHERE order_id = $2 AND dish_id = $3`
	_, err := conn.Exec(ctx, sqlQuery, quantity, orderID, dishID)
	return err
}

func (r *OrderRepository) UpdateDiscount(ctx context.Context, conn *pgx.Conn, orderID int, discountAmount, finalAmount float64) error {
	sqlQuery := `
	UPDATE orders
	SET discount_amount = $1, final_amount = $2, updated_at = NOW()
	WHERE order_id = $3`
	_, err := conn.Exec(ctx, sqlQuery, discountAmount, finalAmount, orderID)
	return err
}
