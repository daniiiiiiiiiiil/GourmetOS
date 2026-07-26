package dto

import (
	"GourmetOS/pkg/errors"
)

type SubscribeRequest struct {
	ObserverType string   `json:"observer_type"`
	Name         string   `json:"name"`
	EventTypes   []string `json:"event_types"`
}

func (r *SubscribeRequest) Validate() error {
	var errs errors.ValidationErrors

	if r.ObserverType == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "observer_type",
			Message: "Тип наблюдателя не может быть пустым",
		})
	}

	if r.Name == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "name",
			Message: "Имя наблюдателя не может быть пустым",
		})
	}

	if len(r.EventTypes) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "event_types",
			Message: "Список событий не может быть пустым",
		})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type UnsubscribeRequest struct {
	ObserverType string `json:"observer_type"`
	Name         string `json:"name"`
}

func (r *UnsubscribeRequest) Validate() error {
	var errs errors.ValidationErrors

	if r.ObserverType == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "observer_type",
			Message: "Тип наблюдателя не может быть пустым",
		})
	}

	if r.Name == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "name",
			Message: "Имя наблюдателя не может быть пустым",
		})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type SendNotificationRequest struct {
	EventType  string   `json:"event_type"`
	OrderID    int      `json:"order_id"`
	TableID    int      `json:"table_id"`
	Items      []string `json:"items"`
	Recipients []string `json:"recipients"`
}

func (r *SendNotificationRequest) Validate() error {
	var errs errors.ValidationErrors

	if r.EventType == "" {
		errs = append(errs, errors.ValidationError{
			Field:   "event_type",
			Message: "Тип события не может быть пустым",
		})
	}

	if r.OrderID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "order_id",
			Message: "ID заказа не может быть меньше или равен 0",
		})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type NotifyOrderRequest struct {
	OrderID int      `json:"order_id"`
	TableID int      `json:"table_id"`
	Items   []string `json:"items"`
}

func (r *NotifyOrderRequest) Validate() error {
	var errs errors.ValidationErrors

	if r.OrderID <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "order_id",
			Message: "ID заказа не может быть меньше или равен 0",
		})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

type SubscriberInfo struct {
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	EventTypes []string `json:"event_types"`
}

type SubscriberListResponse struct {
	Subscribers []SubscriberInfo `json:"subscribers"`
	Total       int              `json:"total"`
}

type NotificationResponse struct {
	Message string `json:"message"`
	Sent    int    `json:"sent"`
	OrderID int    `json:"order_id"`
}

type SubscribeResponse struct {
	Message      string   `json:"message"`
	ObserverType string   `json:"observer_type"`
	Name         string   `json:"name"`
	EventTypes   []string `json:"event_types"`
}
