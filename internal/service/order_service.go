package service

import (
	"GourmetOS/internal/domain"
	"GourmetOS/internal/patterns/command"
	"GourmetOS/internal/patterns/observer"
	"GourmetOS/internal/patterns/strategy"
	"GourmetOS/internal/repository/interfaceRepository"
	"GourmetOS/pkg/errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type OrderService struct {
	orderRepo    interfaceRepository.OrderRepository
	paymentRepo  interfaceRepository.PaymentRepository
	tableRepo    interfaceRepository.TableRepository
	dishRepo     interfaceRepository.DishRepository
	customerRepo interfaceRepository.CustomerRepository
	employeeRepo interfaceRepository.EmployeeRepository
	subject      *observer.OrderSubject
	invoker      *command.CommandInvoker
	logger       *zap.Logger
}

func NewOrderService(
	orderRepo interfaceRepository.OrderRepository,
	paymentRepo interfaceRepository.PaymentRepository,
	tableRepo interfaceRepository.TableRepository,
	dishRepo interfaceRepository.DishRepository,
	customerRepo interfaceRepository.CustomerRepository,
	employeeRepo interfaceRepository.EmployeeRepository,
	subject *observer.OrderSubject,
	invoker *command.CommandInvoker,
	logger *zap.Logger,
) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		paymentRepo:  paymentRepo,
		tableRepo:    tableRepo,
		dishRepo:     dishRepo,
		customerRepo: customerRepo,
		employeeRepo: employeeRepo,
		subject:      subject,
		invoker:      invoker,
		logger:       logger,
	}
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	conn *pgx.Conn,
	tableID int,
	customerID int,
	waiterID int,
	dishIDs []int,
	quantities []int,
) (*domain.Order, error) {
	s.logger.Info("create order started",
		zap.Int("table_id", tableID),
		zap.Int("customer_id", customerID),
		zap.Int("waiter_id", waiterID),
		zap.Ints("dish_ids", dishIDs),
	)
	orderManager := command.NewOrderManager()
	table, err := s.tableRepo.GetByIDTable(ctx, conn, tableID)
	if err != nil {
		s.logger.Error("failed to get table", zap.Int("table_id", tableID), zap.Error(err))
		return nil, errors.NotFoundError{
			Entity: "Table",
			ID:     tableID,
		}
	}
	if table.IsOccupied {
		s.logger.Warn("table is occupied", zap.Int("table_id", tableID))
		return nil, errors.BusinessError{
			Code:    "ErrTableOccupied",
			Message: fmt.Sprintf("Стол #%d уже занят", tableID),
		}
	}

	if customerID > 0 {
		_, err = s.customerRepo.GetByIDCustomer(ctx, conn, customerID)
		if err != nil {
			s.logger.Error("failed to get customer", zap.Int("customer_id", customerID), zap.Error(err))
			return nil, errors.NotFoundError{
				Entity: "Customer",
				ID:     customerID,
			}
		}
	}

	if waiterID > 0 {
		_, err = s.employeeRepo.GetByIDEmployee(ctx, conn, waiterID)
		if err != nil {
			s.logger.Error("failed to get waiter", zap.Int("waiter_id", waiterID), zap.Error(err))
			return nil, errors.NotFoundError{
				Entity: "Employee",
				ID:     waiterID,
			}
		}
	}

	var totalAmount float64
	var dishNames []string

	for i, dishID := range dishIDs {
		dish, err := s.dishRepo.GetByIDDish(ctx, conn, dishID)
		if err != nil {
			s.logger.Error("failed to get dish", zap.Int("dish_id", dishID), zap.Error(err))
			return nil, errors.NotFoundError{
				Entity: "Dish",
				ID:     dishID,
			}
		}
		if !dish.IsAvailable {
			s.logger.Warn("dish is not available", zap.Int("dish_id", dishID), zap.String("dish_name", dish.Name))
			return nil, errors.BusinessError{
				Code:    "ErrDishNotAvailable",
				Message: fmt.Sprintf("Блюдо '%s' не доступно", dish.Name),
			}
		}

		quantity := 1
		if i < len(quantities) {
			quantity = quantities[i]
		}

		totalAmount += dish.Price * float64(quantity)
		dishNames = append(dishNames, dish.Name)
	}

	newOrder := &domain.Order{
		TableID:        tableID,
		CustomerID:     customerID,
		WaiterID:       waiterID,
		Status:         "created",
		TotalAmount:    totalAmount,
		DiscountAmount: 0,
		FinalAmount:    totalAmount,
		PaymentStatus:  "pending",
		Notes:          nil,
	}

	createdOrder, err := s.orderRepo.CreateOrder(ctx, conn, *newOrder)
	if err != nil {
		s.logger.Error("failed to create order", zap.Error(err))
		return nil, errors.BusinessError{
			Code:    "ErrCreateOrder",
			Message: "Не удалось создать заказ: " + err.Error(),
		}
	}

	stateOrder := convertDomainOrderToStateOrder(createdOrder)
	s.logger.Debug("order state set", zap.String("state", stateOrder.State.GetName()))

	if err := s.tableRepo.UpdateOccupiedTable(ctx, conn, tableID, true); err != nil {
		s.logger.Error("failed to occupy table", zap.Int("table_id", tableID), zap.Error(err))
	}

	event := observer.NewEvent(
		observer.OrderCreated,
		createdOrder.OrderID,
		createdOrder.TableID,
		dishNames,
	)
	s.subject.NotifyObservers(event)

	createCmd := command.NewCreateOrderCommand(orderManager, tableID)
	if err := s.invoker.Execute(createCmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	s.logger.Info("order created successfully",
		zap.Int("order_id", createdOrder.OrderID),
		zap.Float64("total_amount", createdOrder.TotalAmount),
		zap.String("status", createdOrder.Status),
	)

	return createdOrder, nil
}

func (s *OrderService) GetOrderService(ctx context.Context, conn *pgx.Conn, id int) (*domain.Order, error) {
	if id <= 0 {
		return nil, errors.ValidationError{
			Field:   "id",
			Message: "ID не может быть меньше или равно нулю",
		}
	}
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, id)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "Order",
			ID:     id,
		}
	}
	return order, nil
}

func (s *OrderService) AddDishService(ctx context.Context, conn *pgx.Conn, orderID, dishID, quantity int) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	dish, err := s.dishRepo.GetByIDDish(ctx, conn, dishID)
	if err != nil {
		return errors.NotFoundError{
			Entity: "Dish",
			ID:     dishID,
		}
	}

	if order.Status == "paid" || order.Status == "cancelled" {
		return errors.BusinessError{
			Code:    "Error",
			Message: "Статус paid или cancelled",
		}
	}
	orderItem := domain.OrderItem{
		OrderID:  orderID,
		DishID:   dishID,
		Quantity: quantity,
		Price:    dish.Price,
	}
	if err := s.orderRepo.AddOrderItem(ctx, conn, orderItem); err != nil {
		s.logger.Error("failed to add order item", zap.Error(err))
		return errors.BusinessError{
			Code:    "ErrAddOrderItem",
			Message: "Не удалось добавить блюдо в заказ",
		}
	}
	total, err := s.calculateOrderTotal(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to calculate order total", zap.Error(err))
		return errors.BusinessError{
			Code:    "ErrCalculateOrderTotal",
			Message: "Не удалось пересчитать сумму",
		}
	}
	if err := s.orderRepo.UpdateTotalOrder(ctx, conn, orderID, total); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateTotalOrder",
			Message: "Не удалось обновить сумму" + err.Error(),
		}
	}
	cmd := command.NewAddDishCommand(orderID, dish.Name, dish.Price)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	return nil
}
func (s *OrderService) calculateOrderTotal(ctx context.Context, conn *pgx.Conn, orderID int) (float64, error) {
	items, err := s.orderRepo.GetOrderItems(ctx, conn, orderID)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total, nil
}

func (s *OrderService) RemoveDishFromOrder(ctx context.Context, conn *pgx.Conn, orderID, dishID int) error {
	if _, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID); err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}

	dish, err := s.dishRepo.GetByIDDish(ctx, conn, dishID)
	if err != nil {
		return errors.NotFoundError{Entity: "Dish", ID: dishID}
	}

	if err := s.orderRepo.RemoveOrderItem(ctx, conn, orderID, dishID); err != nil {
		return errors.BusinessError{
			Code:    "ErrRemoveOrderItem",
			Message: "Не удалось удалить блюдо: " + err.Error(),
		}
	}

	total, err := s.calculateOrderTotal(ctx, conn, orderID)
	if err != nil {
		return err
	}

	if err := s.orderRepo.UpdateTotalOrder(ctx, conn, orderID, total); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateTotalOrder",
			Message: "Не удалось обновить сумму: " + err.Error(),
		}
	}

	cmd := command.NewRemoveDishCommand(orderID, dish.Name, dish.Price)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	s.logger.Info("блюдо удалено", zap.Int("order_id", orderID), zap.Int("dish_id", dishID))
	return nil
}

func (s *OrderService) SubmitToKitchen(ctx context.Context, conn *pgx.Conn, orderID int) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}

	if order.Status != "created" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: "Нельзя отправить на кухню с таким статусом" + order.Status,
		}
	}

	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "in_kitchen"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось обновить статус" + err.Error(),
		}
	}
	event := observer.NewEvent(
		observer.OrderCreated,
		order.OrderID,
		order.TableID,
		[]string{},
	)

	s.subject.NotifyObservers(event)

	cmd := command.NewSubmitToKitchenCommand(orderID)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	return nil
}

func (s *OrderService) MarkAsReadyService(ctx context.Context, conn *pgx.Conn, orderID int) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}
	if order.Status != "cooking" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: "Не подходящий статус заказа" + order.Status,
		}
	}
	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "ready"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось обновить статус на ready" + err.Error(),
		}
	}
	event := observer.NewEvent(
		observer.OrderCreated,
		order.OrderID,
		order.TableID,
		[]string{})
	s.subject.NotifyObservers(event)

	cmd := command.NewSubmitToKitchenCommand(orderID)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}
	return nil
}

func (s *OrderService) ServeToTableService(ctx context.Context, conn *pgx.Conn, orderID int) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}
	if order.Status != "ready" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: "Не подходящий статус должен быть ready" + order.Status,
		}
	}
	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "served"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось обновить статус на served" + err.Error(),
		}
	}
	event := observer.NewEvent(
		observer.OrderCreated,
		order.OrderID,
		order.TableID,
		[]string{})
	s.subject.NotifyObservers(event)

	cmd := command.NewServeToTableCommand(orderID)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}
	return nil
}

func (s *OrderService) ProcessPayment(ctx context.Context, conn *pgx.Conn, orderID int, method string, cardNumber string, expiry string, cvv string) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}

	if order.Status != "served" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: "Не подходящий статус served должен быть,а у вас" + order.Status,
		}
	}
	//strategy
	var paymentStrategy strategy.PaymentStrategy

	switch method {
	case "cash":
		paymentStrategy = strategy.NewCashPayment()
	case "card":
		if cardNumber == "" || expiry == "" || cvv == "" {
			return errors.BusinessError{
				Code:    "ErrInvalidCardData",
				Message: "Не все данные карты заполнены",
			}
		}
		paymentStrategy = strategy.NewCardPayment(cardNumber, "CARD HOLDER", expiry, cvv)
	case "crypto":
		paymentStrategy = strategy.NewCryptoPayment("wallet_address", "BTC")

	default:
		return errors.BusinessError{
			Code:    "ErrUnknownPaymentMethod",
			Message: "Неизвестный способ оплаты",
		}
	}
	if err := paymentStrategy.Pay(order.FinalAmount); err != nil {
		return errors.BusinessError{
			Code:    "ErrPaymentMethod",
			Message: "Не удалось провести оплату" + err.Error(),
		}
	}
	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "paid"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось обновить статус заказа paid" + err.Error(),
		}
	}

	event := observer.NewEvent(
		observer.OrderCreated,
		order.OrderID,
		order.TableID,
		[]string{})
	s.subject.NotifyObservers(event)

	cmd := command.NewProcessPaymentCommand(orderID, method, order.FinalAmount)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}
	return nil
}

func (s *OrderService) CancelOrder(ctx context.Context, conn *pgx.Conn, orderID int) error {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return errors.NotFoundError{Entity: "Order", ID: orderID}
	}
	if order.Status == "paid" {
		return errors.BusinessError{
			Code:    "ErrInvalidStatus",
			Message: "Статус не должен быть paid или completed, у вас " + order.Status,
		}
	}
	if order.Status == "cancelled" {
		return errors.BusinessError{
			Code:    "ErrAlreadyCancelled",
			Message: "Заказ уже отменён",
		}
	}
	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, "cancelled"); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateStatusOrder",
			Message: "Не удалось поменять статус заказа на cancelled" + err.Error(),
		}
	}
	if err := s.tableRepo.UpdateOccupiedTable(ctx, conn, order.TableID, false); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateOccupiedTable",
			Message: "Не удалось освободить стол" + err.Error(),
		}
	}
	cmd := command.NewCancelOrderCommand(orderID, order.TableID)
	if err := s.invoker.Execute(cmd); err != nil {
		s.logger.Warn("failed to save command to history", zap.Error(err))
	}

	return nil
}

func (s *OrderService) UndoLastAction(ctx context.Context, conn *pgx.Conn) error {
	if err := s.invoker.Undo(); err != nil {
		return errors.BusinessError{
			Code:    "ErrUndo",
			Message: "Не удалось отменить действие: " + err.Error(),
		}
	}
	return nil
}

func (s *OrderService) RedoLastAction(ctx context.Context, conn *pgx.Conn) error {
	if err := s.invoker.Redo(); err != nil {
		return errors.BusinessError{
			Code:    "ErrRedo",
			Message: "Не удалось повторить действие: " + err.Error(),
		}
	}
	return nil
}

func (s *OrderService) GetOrderHistory(ctx context.Context, conn *pgx.Conn, orderID int) ([]string, error) {
	order, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to get order", zap.Int("order_id", orderID), zap.Error(err))
		return nil, errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	history := []string{
		"created",
		order.Status,
	}

	return history, nil
}

func (s *OrderService) GetActiveOrders(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, int, error) {
	limit, offset = limitOffset(limit, offset)

	statuses := []string{"created", "in_kitchen", "cooking", "ready", "served"}
	var allOrders []domain.Order

	for _, status := range statuses {
		orders, err := s.orderRepo.GetByStatusOrder(ctx, conn, status, limit, offset)
		if err != nil {
			continue
		}
		allOrders = append(allOrders, orders...)
	}

	total := len(allOrders)

	return allOrders, total, nil
}

func (s *OrderService) GetOrderTotal(ctx context.Context, conn *pgx.Conn, orderID int) (float64, error) {
	_, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		return 0, errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

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

	return total, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, int, error) {
	limit, offset = limitOffset(limit, offset)
	s.logger.Debug("get all orders", zap.Int("limit", limit), zap.Int("offset", offset))

	orders, err := s.orderRepo.GetAllOrders(ctx, conn, limit, offset)
	if err != nil {
		s.logger.Error("failed to get all orders", zap.Error(err))
		return nil, 0, errors.BusinessError{
			Code:    "ErrGetAllOrders",
			Message: "Не удалось получить список заказов: " + err.Error(),
		}
	}

	total := len(orders)

	s.logger.Debug("all orders retrieved", zap.Int("count", len(orders)), zap.Int("total", total))
	return orders, total, nil
}
