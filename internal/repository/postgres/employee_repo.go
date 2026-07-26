package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	Pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{Pool: pool}
}

func (e *EmployeeRepository) CreateEmployee(ctx context.Context, conn *pgx.Conn, employee domain.Employees) (*domain.Employees, error) {
	sqlQuery := `
	INSERT INTO employees (name, email, phone, role, shift, hire_date, salary, is_active, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING employee_id`

	err := conn.QueryRow(ctx, sqlQuery,
		employee.Name,
		employee.Email,
		employee.Phone,
		employee.Role,
		employee.Shift,
		employee.HireDate,
		employee.Salary,
		employee.IsActive,
		time.Now(),
		employee.UpdatedAt,
	).Scan(&employee.EmployeeID)

	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *EmployeeRepository) GetByIDEmployee(ctx context.Context, conn *pgx.Conn, id int) (*domain.Employees, error) {
	sqlQuery := `
	SELECT *
	FROM employees
	WHERE employee_id = $1`

	rows, err := conn.Query(ctx, sqlQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employee, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *EmployeeRepository) GetAllEmployees(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Employees, error) {
	sqlQuery := `
	SELECT *
	FROM employees
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeRepository) GetByRoleEmployee(ctx context.Context, conn *pgx.Conn, role string, limit, offset int) ([]domain.Employees, error) {
	sqlQuery := `
		SELECT *
		FROM employees
		WHERE role = $1
		LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, role, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeRepository) GetActiveEmployees(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Employees, error) {
	sqlQuery := `
		SELECT *
		FROM employees
		WHERE is_active = true
		LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeRepository) UpdateEmployee(ctx context.Context, conn *pgx.Conn, employee domain.Employees, id int) (*domain.Employees, error) {
	sqlQuery := `
	UPDATE employees
	SET name = $1,
		email = $2,
		phone = $3,
		role = $4,
		shift = $5,
		hire_date = $6,
		salary = $7,
		is_active = $8,
		updated_at = $9
	WHERE employee_id = $10
	RETURNING employee_id, name, email, phone, role, shift, hire_date, 
			  salary, is_active, created_at, updated_at`

	rows, err := conn.Query(ctx, sqlQuery,
		employee.Name,
		employee.Email,
		employee.Phone,
		employee.Role,
		employee.Shift,
		employee.HireDate,
		employee.Salary,
		employee.IsActive,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (e *EmployeeRepository) DeleteEmployee(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `DELETE FROM employees WHERE employee_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) GetByEmailEmployee(ctx context.Context, conn *pgx.Conn, email string) (*domain.Employees, error) {
	sqlQuery := `
		SELECT *
		FROM employees
		WHERE email = $1`
	rows, err := conn.Query(ctx, sqlQuery, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}
	return &employees, nil
}

func (e *EmployeeRepository) GetByShiftEmployee(ctx context.Context, conn *pgx.Conn, shift string, limit, offset int) ([]domain.Employees, error) {
	sqlQuery := `
		SELECT *
		FROM employees
		WHERE shift = $1
		LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, shift, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Employees])
	if err != nil {
		return nil, err
	}
	return employees, nil
}
