package service

import (
	"GourmetOS/internal/domain"
	"GourmetOS/internal/patterns/adapter"
	"GourmetOS/internal/patterns/command"
	"GourmetOS/internal/patterns/strategy"
	"GourmetOS/internal/repository/interfaceRepository"
	"GourmetOS/pkg/errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type PaymentService struct {
	paymentRepo interfaceRepository.PaymentRepository
	orderRepo   interfaceRepository.OrderRepository
	invoker     *command.CommandInvoker
	logger      *zap.Logger
}

func NewPaymentService(
	paymentRepo interfaceRepository.PaymentRepository,
	orderRepo interfaceRepository.OrderRepository,
	invoker *command.CommandInvoker,
	logger *zap.Logger,
) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		invoker:     invoker,
		logger:      logger,
	}
}

func (s *PaymentService) ProcessPayment(
	ctx context.Context,
	conn *pgx.Conn,
	orderID int,
	method string,
	cardNumber string,
	cardHolder string,
	expiry string,
	cvv string,
	wallet string,
) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	if order.Status != "served" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: fmt.Sprintf("Нельзя оплатить заказ в статусе '%s'", order.Status),
		}
	}

	amount := order.FinalAmount

	var paymentStrategy strategy.PaymentStrategy

	switch method {
	case "cash":
		paymentStrategy = strategy.NewCashPayment()
	case "card":
		if cardNumber == "" || cardHolder == "" || expiry == "" || cvv == "" {
			return errors.BusinessError{
				Code:    "ErrInvalidCardData",
				Message: "Не все данные карты заполнены",
			}
		}
		paymentStrategy = strategy.NewCardPayment(cardNumber, cardHolder, expiry, cvv)
	case "crypto":
		if wallet == "" {
			return errors.BusinessError{
				Code:    "ErrInvalidWallet",
				Message: "Адрес кошелька не может быть пустым",
			}
		}
		paymentStrategy = strategy.NewCryptoPayment(wallet, "BTC")
	default:
		return errors.BusinessError{
			Code:    "ErrUnknownPaymentMethod",
			Message: "Неизвестный способ оплаты",
		}
	}

	if err := paymentStrategy.Pay(amount); err != nil {
		return errors.BusinessError{
			Code:    "ErrPaymentFailed",
			Message: "Не удалось провести оплату: " + err.Error(),
		}
	}

	payment := &domain.Payment{
		OrderID: orderID,
		Amount:  amount,
		Method:  method,
		Status:  "completed",
	}
	if _, err := s.paymentRepo.CreatePayment(ctx, conn, *payment); err != nil {
		return errors.BusinessError{
			Code:    "ErrCreatePayment",
			Message: "Не удалось создать запись о платеже: " + err.Error(),
		}
	}

	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "paid"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось обновить статус заказа: " + err.Error(),
		}
	}

	cmd := command.NewProcessPaymentCommand(orderID, method, amount)
	if err := s.invoker.Execute(cmd); err != nil {
	}

	return nil
}

func (s *PaymentService) GetPaymentMethods(ctx context.Context, conn *pgx.Conn, method string) ([]domain.Payment, error) {
	pay, err := s.paymentRepo.GetByMethodPayment(ctx, conn, method)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrorGetPayment",
			Message: "Не удалось вернуть оплаты по этому методу" + err.Error(),
		}
	}
	return pay, nil
}

func (s *PaymentService) GetPaymentStatus(ctx context.Context, conn *pgx.Conn, paymentID int) (string, error) {

	payment, err := s.paymentRepo.GetByIDPayment(ctx, conn, paymentID)
	if err != nil {
		s.logger.Error("failed to get payment", zap.Int("payment_id", paymentID), zap.Error(err))
		return "", errors.NotFoundError{
			Entity: "Payment",
			ID:     paymentID,
		}
	}

	if payment.Method == "cash" {
		return payment.Status, nil
	}

	if payment.Method == "card" {
		visaAdapter := adapter.NewVisaAdapter("1234567890123456", "12/26", "123")

		txnID := payment.TransactionID

		status, err := visaAdapter.GetPaymentStatus(txnID)
		if err != nil {
			return "", errors.BusinessError{
				Code:    "ErrGetPaymentStatus",
				Message: "Не удалось получить статус: " + err.Error(),
			}
		}

		if status != payment.Status {
			if err := s.paymentRepo.UpdateStatusPayment(ctx, conn, paymentID, status); err != nil {
				return "", errors.BusinessError{
					Code:    "ErrUpdateStatusPayment",
					Message: "Не удалось обновить статус " + err.Error(),
				}
			}
		}

		return status, nil
	}

	if payment.Method == "crypto" {
		cryptoAdapter := adapter.NewCryptoAdapter(
			"from_wallet",
			payment.CryptoAddress,
			"BTC",
		)

		txnID := payment.TransactionID

		status, err := cryptoAdapter.GetPaymentStatus(txnID)
		if err != nil {
			return "", errors.BusinessError{
				Code:    "ErrGetPaymentStatus",
				Message: "Не удалось получить статус: " + err.Error(),
			}
		}

		if status != payment.Status {
			if err := s.paymentRepo.UpdateStatusPayment(ctx, conn, paymentID, status); err != nil {
				return "", errors.BusinessError{
					Code:    "ErrUpdateStatusPayment",
					Message: "Не удалось обновить статус " + err.Error(),
				}
			}
		}

		return status, nil
	}

	return "", errors.BusinessError{
		Code:    "ErrUnknownMethod",
		Message: fmt.Sprintf("Неизвестный метод оплаты: %s", payment.Method),
	}
}

func (s *PaymentService) RefundPayment(ctx context.Context, conn *pgx.Conn, paymentID int) error {

	payment, err := s.paymentRepo.GetByIDPayment(ctx, conn, paymentID)
	if err != nil {
		return errors.NotFoundError{
			Entity: "Payment",
			ID:     paymentID,
		}
	}

	if payment.Status != "completed" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: fmt.Sprintf("Нельзя вернуть деньги для платежа в статусе '%s'", payment.Status),
		}
	}

	if payment.Method == "cash" {
		return errors.BusinessError{
			Code:    "ErrCashRefund",
			Message: "Возврат наличных денег невозможен, только вручную",
		}
	}

	var payAdapter adapter.PaymentAdapter

	switch payment.Method {
	case "card":
		payAdapter = adapter.NewVisaAdapter("1234567890123456", "12/26", "123")

	case "crypto":
		payAdapter = adapter.NewCryptoAdapter(
			"from_wallet",
			payment.CryptoAddress,
			"BTC",
		)

	default:
		return errors.BusinessError{
			Code:    "ErrUnknownMethod",
			Message: fmt.Sprintf("Неизвестный метод оплаты: %s", payment.Method),
		}
	}

	txnID := payment.TransactionID

	if err := payAdapter.RefundPayment(txnID); err != nil {
		return errors.BusinessError{
			Code:    "ErrRefundPayment",
			Message: "Не удалось выполнить возврат: " + err.Error(),
		}
	}

	if err := s.paymentRepo.UpdateStatusPayment(ctx, conn, paymentID, "refunded"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatus",
			Message: "Не удалось обновить статус платежа: " + err.Error(),
		}
	}

	cmd := command.NewRefundPaymentCommand(paymentID, payment.Amount)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	return nil
}

func (s *PaymentService) ApplyDiscount(ctx context.Context, conn *pgx.Conn, orderID int, discountType string) (float64, error) {
	s.logger.Info("apply discount",
		zap.Int("order_id", orderID),
		zap.String("discount_type", discountType),
	)

	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to get order", zap.Int("order_id", orderID), zap.Error(err))
		return 0, errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	amount := order.TotalAmount

	var discountStrategy strategy.DiscountStrategy

	switch discountType {
	case "no_discount":
		discountStrategy = strategy.NewNoDiscountStrategy()
	case "seasonal":
		discountStrategy = strategy.NewSeasonalDiscountStrategy()
	case "loyalty":
		visits := 15
		discountStrategy = strategy.NewLoyaltyDiscountStrategy(visits)
	case "birthday":
		discountStrategy = strategy.NewBirthdayDiscountStrategy()
	case "order_discount":
		discountStrategy = strategy.NewOrderDiscountStrategy(1000, 50)
	default:
		return 0, errors.ValidationError{
			Field:   "discount_type",
			Message: fmt.Sprintf("Неизвестный тип скидки: %s", discountType),
		}
	}

	finalAmount := discountStrategy.ApplyDiscount(amount)

	discountAmount := amount - finalAmount
	if err := s.orderRepo.UpdateDiscount(ctx, conn, orderID, discountAmount, finalAmount); err != nil {
		return 0, errors.BusinessError{
			Code:    "ErrUpdateDiscount",
			Message: "Не удалось сохранить скидку: " + err.Error(),
		}
	}

	return finalAmount, nil
}

func (s *PaymentService) CalculateFinalAmount(ctx context.Context, conn *pgx.Conn, orderID int, discountType string) (float64, error) {

	items, err := s.orderRepo.GetOrderItems(ctx, conn, orderID)
	if err != nil {
		return 0, errors.BusinessError{
			Code:    "ErrGetOrderItems",
			Message: "Не удалось получить блюда заказа: " + err.Error(),
		}
	}

	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	var discountStrategy strategy.DiscountStrategy

	switch discountType {
	case "no_discount":
		discountStrategy = strategy.NewNoDiscountStrategy()
	case "seasonal":
		discountStrategy = strategy.NewSeasonalDiscountStrategy()
	case "loyalty":
		visits := 15
		discountStrategy = strategy.NewLoyaltyDiscountStrategy(visits)
	case "birthday":
		discountStrategy = strategy.NewBirthdayDiscountStrategy()
	case "order_discount":
		discountStrategy = strategy.NewOrderDiscountStrategy(1000, 50)
	}

	finalAmount := discountStrategy.ApplyDiscount(total)
	discountAmount := total - finalAmount
	_ = discountAmount

	if err := s.orderRepo.UpdateTotalOrder(ctx, conn, orderID, finalAmount); err != nil {
		return 0, errors.BusinessError{
			Code:    "ErrUpdateTotalOrder",
			Message: "Не удалось обновить итоговую сумму: " + err.Error(),
		}
	}

	return finalAmount, nil
}

func (s *PaymentService) GetPaymentByOrder(ctx context.Context, conn *pgx.Conn, orderID int) (*domain.Payment, error) {
	pay, err := s.paymentRepo.GetByOrderIDPayment(ctx, conn, orderID)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "Payment",
			ID:     orderID,
		}
	}
	return pay, nil
}

func (s *PaymentService) GetPaymentByTransaction(ctx context.Context, conn *pgx.Conn, transactionID string) (*domain.Payment, error) {
	pay, err := s.paymentRepo.GetByTransactionIDPayment(ctx, conn, transactionID)
	if err != nil {
		return nil, err
	}
	return pay, nil
}
