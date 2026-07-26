package service

import (
	"context"
	"time"

	"GourmetOS/internal/domain"
	"GourmetOS/internal/patterns/observer"
	"GourmetOS/internal/patterns/state"
	"GourmetOS/internal/patterns/template_method"
	"GourmetOS/internal/repository/interfaceRepository"
	"GourmetOS/pkg/errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type KitchenService struct {
	orderRepo interfaceRepository.OrderRepository
	dishRepo  interfaceRepository.DishRepository
	subject   *observer.OrderSubject
	logger    *zap.Logger
}

func NewKitchenService(
	orderRepo interfaceRepository.OrderRepository,
	dishRepo interfaceRepository.DishRepository,
	subject *observer.OrderSubject,
	logger *zap.Logger,
) *KitchenService {
	return &KitchenService{
		orderRepo: orderRepo,
		dishRepo:  dishRepo,
		subject:   subject,
		logger:    logger,
	}
}

func limitOffset(limit, offset int) (int, int) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func convertDomainOrderToStateOrder(domainOrder *domain.Order) *state.Order {
	stateOrder := state.NewOrder(domainOrder.OrderID, domainOrder.TableID)
	stateOrder.Total = domainOrder.TotalAmount
	stateOrder.IsActive = true

	switch domainOrder.Status {
	case "created":
		stateOrder.SetState(state.NewCreatedState())
	case "in_kitchen":
		stateOrder.SetState(state.NewInKitchenState())
	case "cooking":
		stateOrder.SetState(state.NewCookingState())
	case "ready":
		stateOrder.SetState(state.NewReadyState())
	case "served":
		stateOrder.SetState(state.NewServedState())
	case "paid":
		stateOrder.SetState(state.NewPaidState())
	default:
		stateOrder.SetState(state.NewCreatedState())
	}

	return stateOrder
}

func (s *KitchenService) ReceiveOrder(ctx context.Context, conn *pgx.Conn, orderID int) error {
	s.logger.Info("kitchen received order", zap.Int("order_id", orderID))

	domainOrder, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to get order", zap.Int("order_id", orderID), zap.Error(err))
		return errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	order := convertDomainOrderToStateOrder(domainOrder)

	if err := order.State.SubmitToKitchen(order); err != nil {
		s.logger.Warn("failed to submit order to kitchen", zap.Int("order_id", orderID), zap.Error(err))
		return errors.BusinessError{
			Code:    "ErrSubmitToKitchen",
			Message: err.Error(),
		}
	}

	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, order.Status); err != nil {
		s.logger.Error("failed to update order status", zap.Int("order_id", orderID), zap.Error(err))
		return err
	}

	s.logger.Info("order received by kitchen", zap.Int("order_id", orderID), zap.String("status", order.Status))
	return nil
}

func (s *KitchenService) StartCooking(ctx context.Context, conn *pgx.Conn, orderID int) error {
	s.logger.Info("start cooking order", zap.Int("order_id", orderID))

	domainOrder, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to get order", zap.Int("order_id", orderID), zap.Error(err))
		return errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	order := convertDomainOrderToStateOrder(domainOrder)

	if err := order.State.StartCooking(order); err != nil {
		s.logger.Warn("failed to start cooking", zap.Int("order_id", orderID), zap.Error(err))
		return errors.BusinessError{
			Code:    "ErrStartCooking",
			Message: err.Error(),
		}
	}

	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, order.Status); err != nil {
		s.logger.Error("failed to update order status", zap.Int("order_id", orderID), zap.Error(err))
		return err
	}

	s.logger.Info("order cooking started", zap.Int("order_id", orderID))
	return nil
}

func (s *KitchenService) MarkAsReady(ctx context.Context, conn *pgx.Conn, orderID int) error {
	s.logger.Info("mark order as ready", zap.Int("order_id", orderID))

	domainOrder, err := s.orderRepo.GetByIdOrder(ctx, conn, orderID)
	if err != nil {
		s.logger.Error("failed to get order", zap.Int("order_id", orderID), zap.Error(err))
		return errors.NotFoundError{
			Entity: "Order",
			ID:     orderID,
		}
	}

	order := convertDomainOrderToStateOrder(domainOrder)

	if err := order.State.MarkAsReady(order); err != nil {
		s.logger.Warn("failed to mark as ready", zap.Int("order_id", orderID), zap.Error(err))
		return errors.BusinessError{
			Code:    "ErrMarkAsReady",
			Message: err.Error(),
		}
	}

	if err := s.orderRepo.UpdateStatusOrder(ctx, conn, orderID, order.Status); err != nil {
		s.logger.Error("failed to update order status", zap.Int("order_id", orderID), zap.Error(err))
		return err
	}

	event := observer.NewEvent(
		observer.OrderReady,
		order.ID,
		order.TableID,
		[]string{},
	)
	s.subject.NotifyObservers(event)

	s.logger.Info("order marked as ready", zap.Int("order_id", orderID))
	return nil
}

func (s *KitchenService) GetCookingQueue(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Order, int, error) {
	limit, offset = limitOffset(limit, offset)
	s.logger.Debug("get cooking queue", zap.Int("limit", limit), zap.Int("offset", offset))

	orders, err := s.orderRepo.GetByStatusOrder(ctx, conn, "in_kitchen", limit, offset)
	if err != nil {
		s.logger.Error("failed to get cooking queue", zap.Error(err))
		return nil, 0, err
	}

	total := len(orders)

	s.logger.Debug("cooking queue retrieved", zap.Int("count", len(orders)), zap.Int("total", total))
	return orders, total, nil
}

func (s *KitchenService) GetCookingTime(ctx context.Context, conn *pgx.Conn, dishID int) (int, error) {
	s.logger.Debug("get cooking time", zap.Int("dish_id", dishID))

	dish, err := s.dishRepo.GetByIDDish(ctx, conn, dishID)
	if err != nil {
		s.logger.Error("failed to get dish", zap.Int("dish_id", dishID), zap.Error(err))
		return 0, errors.NotFoundError{
			Entity: "Dish",
			ID:     dishID,
		}
	}

	var cookingTime int

	switch dish.Category {
	case "pizza":
		prep := template_method.NewPizzaPreparation("средняя", dish.Name)
		cookingTime = prep.GetCookingTime()
	case "pasta":
		prep := template_method.NewPastaPreparation("спагетти", dish.Cuisine)
		cookingTime = prep.GetCookingTime()
	case "salad":
		prep := template_method.NewSaladPreparation(dish.Name)
		cookingTime = prep.GetCookingTime()
	default:
		prep := template_method.NewDishPreparation(dish.Name)
		cookingTime = prep.GetCookingTime()
	}

	s.logger.Debug("cooking time retrieved", zap.Int("dish_id", dishID), zap.Int("cooking_time", cookingTime))
	return cookingTime, nil
}

func (s *KitchenService) GetKitchenStatus(ctx context.Context, conn *pgx.Conn) (map[string]interface{}, error) {
	s.logger.Debug("get kitchen status")

	activeOrders, err := s.orderRepo.GetByStatusOrder(ctx, conn, "in_kitchen", 100, 0)
	if err != nil {
		s.logger.Error("failed to get active orders", zap.Error(err))
		return nil, err
	}

	cookingOrders, err := s.orderRepo.GetByStatusOrder(ctx, conn, "cooking", 100, 0)
	if err != nil {
		s.logger.Error("failed to get cooking orders", zap.Error(err))
		return nil, err
	}

	status := map[string]interface{}{
		"in_queue":  len(activeOrders),
		"cooking":   len(cookingOrders),
		"total":     len(activeOrders) + len(cookingOrders),
		"timestamp": time.Now().Format(time.RFC3339),
	}

	s.logger.Debug("kitchen status retrieved", zap.Int("in_queue", len(activeOrders)), zap.Int("cooking", len(cookingOrders)))
	return status, nil
}
