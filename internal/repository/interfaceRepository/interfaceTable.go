package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type TableRepository interface {
	CreateTable(ctx context.Context, conn *pgx.Conn, table domain.Table) (*domain.Table, error)
	GetByIDTable(ctx context.Context, conn *pgx.Conn, id int) (*domain.Table, error)
	UpdateTable(ctx context.Context, conn *pgx.Conn, id int, table domain.Table) (*domain.Table, error)
	DeleteTable(ctx context.Context, conn *pgx.Conn, id int) error
	GetAllTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error)
	GetFreeTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error)
	GetOccupiedTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error)
	UpdateOccupiedTable(ctx context.Context, conn *pgx.Conn, id int, occopied bool) error
	GetByLocationTable(ctx context.Context, conn *pgx.Conn, location string, limit, offset int) ([]domain.Table, error)
	GetByCapacityTable(ctx context.Context, conn *pgx.Conn, capacity int, limit, offset int) ([]domain.Table, error)
	UpdateReservedTable(ctx context.Context, conn *pgx.Conn, id int, isReserved bool) error
	ReserveTable(ctx context.Context, conn *pgx.Conn, id int) error
	CancelReservation(ctx context.Context, conn *pgx.Conn, id int) error
}
