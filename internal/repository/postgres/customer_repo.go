package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	Pool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{Pool: pool}
}

func (c *CustomerRepository) CreateCustomer(ctx context.Context, conn *pgx.Conn, customer domain.Customer) (*domain.Customer, error) {
	sqlQuery := `
	INSERT INTO customers(name, phone, email, address, loyalty_level, is_active, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING customer_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		customer.Name,
		customer.Phone,
		customer.Email,
		customer.Address,
		customer.LoyaltyLevel,
		customer.IsActive,
		time.Now(),
		time.Now(),
	).Scan(&customer.CustomerID, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) GetByIDCustomer(ctx context.Context, conn *pgx.Conn, customerID int) (*domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	WHERE customer_id = $1`

	rows, err := conn.Query(ctx, sqlQuery, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customer, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) GetAllCustomers(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	LIMIT $1 OFFSET $2`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *CustomerRepository) GetByPhoneCustomer(ctx context.Context, conn *pgx.Conn, phone string) (*domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	WHERE phone = $1`

	rows, err := conn.Query(ctx, sqlQuery, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customer, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) GetByEmailCustomer(ctx context.Context, conn *pgx.Conn, email string) (*domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	WHERE email = $1`

	rows, err := conn.Query(ctx, sqlQuery, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customer, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) UpdateCustomer(ctx context.Context, conn *pgx.Conn, customer domain.Customer, id int) (*domain.Customer, error) {
	sqlQuery := `
	UPDATE customers
	SET name = $1, phone = $2, email = $3, address = $4,
		loyalty_level = $5, is_active = $6, updated_at = $7
	WHERE customer_id = $8
	RETURNING customer_id, name, phone, email, address, loyalty_level, is_active, created_at, updated_at`

	rows, err := conn.Query(ctx, sqlQuery,
		customer.Name,
		customer.Phone,
		customer.Email,
		customer.Address,
		customer.LoyaltyLevel,
		customer.IsActive,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (c *CustomerRepository) UpdateLoyaltyCustomer(ctx context.Context, conn *pgx.Conn, id int, loyaltyLevel string) error {
	sqlQuery := `
	UPDATE customers
	SET loyalty_level = $1, updated_at = $2
	WHERE customer_id = $3`

	_, err := conn.Exec(ctx, sqlQuery, loyaltyLevel, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) DeleteCustomer(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `DELETE FROM customers WHERE customer_id = $1`

	_, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) GetByNameCustomer(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	WHERE name ILIKE $1
	LIMIT $2 OFFSET $3`

	rows, err := conn.Query(ctx, sqlQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *CustomerRepository) GetActiveCustomers(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Customer, error) {
	sqlQuery := `
	SELECT *
	FROM customers
	WHERE is_active = true
	LIMIT $1 OFFSET $2`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}
	return customers, nil
}
