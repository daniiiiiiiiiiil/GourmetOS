package service

import (
	"context"

	"GourmetOS/internal/patterns/observer"
	"GourmetOS/pkg/errors"

	"go.uber.org/zap"
)

type NotificationService struct {
	subject *observer.OrderSubject
	logger  *zap.Logger
}

func NewNotificationService(
	subject *observer.OrderSubject,
	logger *zap.Logger,
) *NotificationService {
	return &NotificationService{
		subject: subject,
		logger:  logger,
	}
}

func (s *NotificationService) SendNotification(ctx context.Context, event observer.Event) (int, error) {
	if event.Type == "" {
		return 0, errors.ValidationError{
			Field:   "event_type",
			Message: "Тип события не может быть пустым",
		}
	}

	s.subject.NotifyObservers(event)

	subscribersCount := s.subject.GetObserversCount()

	s.logger.Info("notification sent",
		zap.String("event_type", string(event.Type)),
		zap.Int("order_id", event.OrderID),
		zap.Int("subscribers", subscribersCount),
	)

	return subscribersCount, nil
}

func (s *NotificationService) Subscribe(
	ctx context.Context,
	observerType string,
	name string,
	eventTypes []string,
) error {
	s.logger.Info("subscribe observer",
		zap.String("observer_type", observerType),
		zap.String("name", name),
		zap.Strings("event_types", eventTypes),
	)

	var obs observer.Observer

	switch observerType {
	case "kitchen":
		obs = observer.NewKitchenDisplay(name)
	case "waiter":
		obs = observer.NewWaiterNotifier(name)
	case "customer":
		obs = observer.NewCustomerNotifier(name)
	case "admin":
		obs = observer.NewAdminNotifier(name)
	default:
		return errors.ValidationError{
			Field:   "observer_type",
			Message: "Неизвестный тип наблюдателя: " + observerType,
		}
	}

	s.subject.RegisterObserver(obs)

	s.logger.Info("observer subscribed successfully",
		zap.String("observer_type", observerType),
		zap.String("name", name),
	)

	return nil
}

func (s *NotificationService) Unsubscribe(ctx context.Context, observerType, name string) error {
	s.logger.Info("unsubscribe observer",
		zap.String("observer_type", observerType),
		zap.String("name", name),
	)

	var obs observer.Observer

	switch observerType {
	case "kitchen":
		obs = observer.NewKitchenDisplay(name)
	case "waiter":
		obs = observer.NewWaiterNotifier(name)
	case "customer":
		obs = observer.NewCustomerNotifier(name)
	case "admin":
		obs = observer.NewAdminNotifier(name)
	default:
		return errors.ValidationError{
			Field:   "observer_type",
			Message: "Неизвестный тип наблюдателя: " + observerType,
		}
	}

	s.subject.RemoveObserver(obs)

	s.logger.Info("observer unsubscribed successfully",
		zap.String("observer_type", observerType),
		zap.String("name", name),
	)

	return nil
}

func (s *NotificationService) GetSubscribers(ctx context.Context) ([]observer.Observer, int, error) {
	subscribers := s.subject.GetObservers()
	count := s.subject.GetObserversCount()

	s.logger.Debug("get subscribers", zap.Int("count", count))

	return subscribers, count, nil
}

func (s *NotificationService) NotifyOrderCreated(ctx context.Context, orderID, tableID int, items []string) error {
	event := observer.NewEvent(observer.OrderCreated, orderID, tableID, items)
	s.subject.NotifyObservers(event)
	s.logger.Info("notified: order created", zap.Int("order_id", orderID))
	return nil
}

func (s *NotificationService) NotifyOrderReady(ctx context.Context, orderID, tableID int, items []string) error {
	event := observer.NewEvent(observer.OrderReady, orderID, tableID, items)
	s.subject.NotifyObservers(event)
	s.logger.Info("notified: order ready", zap.Int("order_id", orderID))
	return nil
}

func (s *NotificationService) NotifyOrderServed(ctx context.Context, orderID, tableID int, items []string) error {
	event := observer.NewEvent(observer.OrderServed, orderID, tableID, items)
	s.subject.NotifyObservers(event)
	s.logger.Info("notified: order served", zap.Int("order_id", orderID))
	return nil
}

func (s *NotificationService) NotifyOrderPaid(ctx context.Context, orderID, tableID int, items []string) error {
	event := observer.NewEvent(observer.OrderPaid, orderID, tableID, items)
	s.subject.NotifyObservers(event)
	s.logger.Info("notified: order paid", zap.Int("order_id", orderID))
	return nil
}
