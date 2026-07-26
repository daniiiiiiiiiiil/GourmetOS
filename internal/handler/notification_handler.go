package handler

import (
	"GourmetOS/internal/patterns/observer"
	"encoding/json"
	"net/http"

	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var req dto.SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.Subscribe(
		r.Context(),
		req.ObserverType,
		req.Name,
		req.EventTypes,
	)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusCreated, map[string]interface{}{
		"message":       "Подписка оформлена",
		"observer_type": req.ObserverType,
		"name":          req.Name,
		"event_types":   req.EventTypes,
	})
}

func (h *NotificationHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var req dto.UnsubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.Unsubscribe(r.Context(), req.ObserverType, req.Name)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, map[string]interface{}{
		"message":       "Отписка выполнена",
		"observer_type": req.ObserverType,
		"name":          req.Name,
	})
}

func (h *NotificationHandler) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	subscribers, count, err := h.service.GetSubscribers(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	var subscriberInfos []dto.SubscriberInfo
	for range subscribers {
		subscriberInfos = append(subscriberInfos, dto.SubscriberInfo{
			Type: "unknown",
			Name: "unknown",
		})
	}

	response := dto.SubscriberListResponse{
		Subscribers: subscriberInfos,
		Total:       count,
	}

	sendSuccess(w, http.StatusOK, response)
}

func (h *NotificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	var req dto.SendNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	event := observer.NewEvent(
		observer.EventType(req.EventType),
		req.OrderID,
		req.TableID,
		req.Items,
	)

	sent, err := h.service.SendNotification(r.Context(), event)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, dto.NotificationResponse{
		Message: "Уведомление отправлено",
		Sent:    sent,
		OrderID: req.OrderID,
	})
}

func (h *NotificationHandler) NotifyOrderCreated(w http.ResponseWriter, r *http.Request) {
	var req dto.NotifyOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.NotifyOrderCreated(r.Context(), req.OrderID, req.TableID, req.Items)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, dto.NotificationResponse{
		Message: "Уведомление о создании заказа отправлено",
		OrderID: req.OrderID,
	})
}

func (h *NotificationHandler) NotifyOrderReady(w http.ResponseWriter, r *http.Request) {
	var req dto.NotifyOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.NotifyOrderReady(r.Context(), req.OrderID, req.TableID, req.Items)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, dto.NotificationResponse{
		Message: "Уведомление о готовности заказа отправлено",
		OrderID: req.OrderID,
	})
}

func (h *NotificationHandler) NotifyOrderServed(w http.ResponseWriter, r *http.Request) {
	var req dto.NotifyOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.NotifyOrderServed(r.Context(), req.OrderID, req.TableID, req.Items)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, dto.NotificationResponse{
		Message: "Уведомление о подаче заказа отправлено",
		OrderID: req.OrderID,
	})
}

func (h *NotificationHandler) NotifyOrderPaid(w http.ResponseWriter, r *http.Request) {
	var req dto.NotifyOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidateError", err.Error())
		return
	}

	err := h.service.NotifyOrderPaid(r.Context(), req.OrderID, req.TableID, req.Items)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "ServerError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, dto.NotificationResponse{
		Message: "Уведомление об оплате заказа отправлено",
		OrderID: req.OrderID,
	})
}
