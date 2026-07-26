package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	Pool *pgxpool.Pool
}

func NewPaymentRepository(pool *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{Pool: pool}
}

func (p *PaymentRepository) CreatePayment(ctx context.Context, conn *pgx.Conn, pay domain.Payment) (*domain.Payment, error) {
	sqlQuery := `
	INSERT INTO payments(order_id, amount, method, status, transaction_id, card_last4, 
						 crypto_address, receipt_url, paid_at, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING payment_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		pay.OrderID,
		pay.Amount,
		pay.Method,
		pay.Status,
		pay.TransactionID,
		pay.CardLast4,
		pay.CryptoAddress,
		pay.ReceiptURL,
		pay.PaidAt,
		time.Now(),
		pay.UpdatedAt,
	).Scan(&pay.PaymentID, &pay.CreatedAt, &pay.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &pay, nil
}

func (p *PaymentRepository) GetByIDPayment(ctx context.Context, conn *pgx.Conn, paymentID int) (*domain.Payment, error) {
	sqlQuery := `
	SELECT *
	FROM payments
	WHERE payment_id=$1`
	rows, err := conn.Query(ctx, sqlQuery, paymentID)
	if err != nil {
		return nil, err
	}
	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	return &payment, nil
}

func (p *PaymentRepository) UpdatePayment(ctx context.Context, conn *pgx.Conn, payment domain.Payment, id int) (*domain.Payment, error) {
	sqlQuery := `
	UPDATE payments
	SET order_id = $1, amount = $2, method = $3, status = $4, 
		transaction_id = $5, card_last4 = $6, crypto_address = $7,
		receipt_url = $8, paid_at = $9, updated_at = $10
	WHERE payment_id = $11
	RETURNING payment_id, order_id, amount, method, status, transaction_id, 
			  card_last4, crypto_address, receipt_url, paid_at, created_at, updated_at`

	rows, err := conn.Query(ctx, sqlQuery,
		payment.OrderID,
		payment.Amount,
		payment.Method,
		payment.Status,
		payment.TransactionID,
		payment.CardLast4,
		payment.CryptoAddress,
		payment.ReceiptURL,
		payment.PaidAt,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (p *PaymentRepository) GetByOrderIDPayment(ctx context.Context, conn *pgx.Conn, orderID int) (*domain.Payment, error) {
	sqlQuery := `
	SELECT *
	FROM payments
	WHERE order_id=$1`
	rows, err := conn.Query(ctx, sqlQuery, orderID)
	if err != nil {
		return nil, err
	}
	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	return &payment, nil
}

func (p *PaymentRepository) UpdateStatusPayment(ctx context.Context, conn *pgx.Conn, id int, status string) error {
	sqlQuery := `
	UPDATE payments
	SET status = $1
	WHERE payment_id=$2`
	_, err := conn.Exec(ctx, sqlQuery, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentRepository) UpdateTransactionID(ctx context.Context, conn *pgx.Conn, id int, transactionID string) error {
	sqlQuery := `
	UPDATE payments
	SET transaction_id = $1
	WHERE payment_id=$2`
	_, err := conn.Exec(ctx, sqlQuery, transactionID, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentRepository) GetByTransactionIDPayment(ctx context.Context, conn *pgx.Conn, transactionID string) (*domain.Payment, error) {
	sqlQuery := `SELECT * FROM payments WHERE transaction_id = $1`
	rows, err := conn.Query(ctx, sqlQuery, transactionID)
	if err != nil {
		return nil, err
	}
	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	return &payment, err
}

func (p *PaymentRepository) GetByMethodPayment(ctx context.Context, conn *pgx.Conn, method string) ([]domain.Payment, error) {
	sqlQuery := `
	SELECT * FROM payments
	WHERE method = $1`

	rows, err := conn.Query(ctx, sqlQuery, method)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Payment])
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (p *PaymentRepository) GetCompletedPayment(ctx context.Context, conn *pgx.Conn) (*domain.Payment, error) {
	sqlQuery := `
		SELECT *
		FROM payments
		WHERE status = 'completed'`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	return &payment, nil
}

func (p *PaymentRepository) GetPendingPayment(ctx context.Context, conn *pgx.Conn, limit, offset int) (*domain.Payment, error) {
	sqlQuery := `
	SELECT *
	FROM payments
	WHERE status = 'pending'
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Payment])
	return &payment, nil
}
