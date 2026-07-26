package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, conn *pgx.Conn, customer domain.Customer) (*domain.Customer, error)
	GetByIDCustomer(ctx context.Context, conn *pgx.Conn, customerID int) (*domain.Customer, error)
	GetAllCustomers(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Customer, error)
	GetByPhoneCustomer(ctx context.Context, conn *pgx.Conn, phone string) (*domain.Customer, error)
	GetByEmailCustomer(ctx context.Context, conn *pgx.Conn, email string) (*domain.Customer, error)
	UpdateCustomer(ctx context.Context, conn *pgx.Conn, customer domain.Customer, id int) (*domain.Customer, error)
	UpdateLoyaltyCustomer(ctx context.Context, conn *pgx.Conn, id int, loyaltyLevel string) error
	DeleteCustomer(ctx context.Context, conn *pgx.Conn, id int) error
	GetByNameCustomer(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Customer, error)
	GetActiveCustomers(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Customer, error)
}
