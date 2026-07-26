package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TableRepository struct {
	Pool *pgxpool.Pool
}

func NewTableRepository(pool *pgxpool.Pool) *TableRepository {
	return &TableRepository{Pool: pool}
}

func (t *TableRepository) CreateTable(ctx context.Context, conn *pgx.Conn, table domain.Table) (*domain.Table, error) {
	sqlQuery := `
	INSERT INTO tables(number, capacity, location, is_occupied, is_reserved, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING table_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		table.Number,
		table.Capacity,
		table.Location,
		table.IsOccupied,
		table.IsReserved,
		time.Now(),
		table.UpdatedAt,
	).Scan(&table.TableID, &table.CreatedAt, &table.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (t *TableRepository) GetByIDTable(ctx context.Context, conn *pgx.Conn, id int) (*domain.Table, error) {
	sqlQuery := `
	SELECT * 
	FROM tables
	WHERE table_id = $1`
	rows, err := conn.Query(ctx, sqlQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	table, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Table])
	return &table, nil
}

func (t *TableRepository) UpdateTable(ctx context.Context, conn *pgx.Conn, id int, table domain.Table) (*domain.Table, error) {
	sqlQuery := `
	UPDATE tables
	SET number = $1, capacity = $2, location = $3, 
		is_occupied = $4, is_reserved = $5, updated_at = $6
	WHERE table_id = $7
	RETURNING table_id, number, capacity, location, is_occupied, is_reserved, 
			  created_at, updated_at`

	rows, err := conn.Query(ctx, sqlQuery,
		table.Number,
		table.Capacity,
		table.Location,
		table.IsOccupied,
		table.IsReserved,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (t *TableRepository) DeleteTable(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `DELETE FROM tables WHERE table_id = $1`
	if _, err := conn.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	return nil
}

func (t *TableRepository) GetAllTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	sqlQuery := `
		SELECT * 
		FROM tables
		LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *TableRepository) GetFreeTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	sqlQuery := `
	SELECT * 
	FROM tables
	WHERE is_occupied = FALSE
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *TableRepository) GetOccupiedTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	sqlQuery := `
	SELECT * 
	FROM tables
	WHERE is_occupied = TRUE`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *TableRepository) UpdateOccupiedTable(ctx context.Context, conn *pgx.Conn, id int, occopied bool) error {
	sqlQuery := `
	UPDATE tables
	SET is_occupied = $2
	WHERE table_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, id, occopied)
	if err != nil {
		return err
	}
	return nil
}

func (t *TableRepository) GetByLocationTable(ctx context.Context, conn *pgx.Conn, location string, limit, offset int) ([]domain.Table, error) {
	sqlQuery := `
	SELECT * 
	FROM tables
	WHERE location = $1
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, location, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *TableRepository) GetByCapacityTable(ctx context.Context, conn *pgx.Conn, capacity, limit, offset int) ([]domain.Table, error) {
	sqlQuery := `
	SELECT * 
	FROM tables
	WHERE $1 >= capacity
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, capacity, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Table])
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (t *TableRepository) UpdateReservedTable(ctx context.Context, conn *pgx.Conn, id int, isReserved bool) error {
	sqlQuery := `
	UPDATE tables
	SET is_reserved = $1, updated_at = $2
	WHERE table_id = $3`

	_, err := conn.Exec(ctx, sqlQuery, isReserved, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TableRepository) ReserveTable(ctx context.Context, conn *pgx.Conn, id int) error {
	return t.UpdateReservedTable(ctx, conn, id, true)
}

func (t *TableRepository) CancelReservation(ctx context.Context, conn *pgx.Conn, id int) error {
	return t.UpdateReservedTable(ctx, conn, id, false)
}
