package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, conn *pgx.Conn, employee domain.Employees) (*domain.Employees, error)
	GetByIDEmployee(ctx context.Context, conn *pgx.Conn, id int) (*domain.Employees, error)
	GetAllEmployees(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Employees, error)
	GetByRoleEmployee(ctx context.Context, conn *pgx.Conn, role string, limit, offset int) ([]domain.Employees, error)
	GetActiveEmployees(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Employees, error)
	UpdateEmployee(ctx context.Context, conn *pgx.Conn, employee domain.Employees, id int) (*domain.Employees, error)
	DeleteEmployee(ctx context.Context, conn *pgx.Conn, id int) error
	GetByEmailEmployee(ctx context.Context, conn *pgx.Conn, email string) (*domain.Employees, error)
	GetByShiftEmployee(ctx context.Context, conn *pgx.Conn, shift string, limit, offset int) ([]domain.Employees, error)
}
