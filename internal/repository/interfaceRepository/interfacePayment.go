package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, conn *pgx.Conn, pay domain.Payment) (*domain.Payment, error)
	GetByIDPayment(ctx context.Context, conn *pgx.Conn, paymentID int) (*domain.Payment, error)
	UpdatePayment(ctx context.Context, conn *pgx.Conn, payment domain.Payment, id int) (*domain.Payment, error)
	GetByOrderIDPayment(ctx context.Context, conn *pgx.Conn, orderID int) (*domain.Payment, error)
	UpdateStatusPayment(ctx context.Context, conn *pgx.Conn, id int, status string) error
	UpdateTransactionID(ctx context.Context, conn *pgx.Conn, id int, transactionID string) error
	GetByTransactionIDPayment(ctx context.Context, conn *pgx.Conn, transactionID string) (*domain.Payment, error)
	GetByMethodPayment(ctx context.Context, conn *pgx.Conn, method string) ([]domain.Payment, error)
	GetCompletedPayment(ctx context.Context, conn *pgx.Conn) (*domain.Payment, error)
	GetPendingPayment(ctx context.Context, conn *pgx.Conn, limit, offset int) (*domain.Payment, error)
}
