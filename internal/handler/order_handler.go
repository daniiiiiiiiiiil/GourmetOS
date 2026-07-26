package handler

import (
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
	"GourmetOS/pkg/pagination"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	var req dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidationError", err.Error())
		return
	}

	dishIDs := make([]int, len(req.Items))
	quantities := make([]int, len(req.Items))

	for i, item := range req.Items {
		dishIDs[i] = item.DishID
		quantities[i] = item.Quantity
	}

	created, err := h.service.CreateOrder(
		r.Context(),
		conn,
		req.TableID,
		req.CustomerID,
		req.WaiterID,
		dishIDs,
		quantities,
	)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "CreateOrderError", err.Error())
		return
	}

	response := dto.OrderResponseFromDomain(*created)
	sendSuccess(w, http.StatusCreated, response)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Не правильный id")
		return
	}
	order, err := h.service.GetOrderService(r.Context(), conn, orderID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "GetOrderError", err.Error())
		return
	}
	resp := dto.OrderResponseFromDomain(*order)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	pagination.LimitOffset(limit, offset)

	orders, total, err := h.service.GetAllOrders(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "GetAllOrdersError", err.Error())
		return
	}
	resp := dto.NewOrderListResponse(orders, total, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *OrderHandler) GetActiveOrders(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	orders, total, err := h.service.GetActiveOrders(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "GetActiveOrdersError", err.Error())
		return
	}
	resp := dto.NewOrderListResponse(orders, total, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *OrderHandler) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Неверный ввод id")
		return
	}
	order, err := h.service.GetOrderHistory(r.Context(), conn, orderID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "GetOrderHistoryError", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, order)
}

func (h *OrderHandler) AddDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	var req dto.AddDishRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", err.Error())
		return
	}
	err := h.service.AddDishService(r.Context(), conn, req.OrderID, req.DishID, req.Quantity)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "AddDishServiceError", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *OrderHandler) RemoveDish(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	var req dto.RemoveDishRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", err.Error())
		return
	}
	err := h.service.RemoveDishFromOrder(r.Context(), conn, req.OrderID, req.DishID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "RemoveDishFromOrderError", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *OrderHandler) SubmitToKitchen(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "OrderIDNotFound", err.Error())
		return
	}

	if err := h.service.SubmitToKitchen(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "SubmitToKitchenError", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *OrderHandler) MarkAsReady(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "OrderIDNotFound", err.Error())
		return
	}
	if err := h.service.MarkAsReadyService(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "MarkAsReadyError", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *OrderHandler) ServeToTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "OrderIDNotFound", err.Error())
		return
	}
	if err := h.service.ServeToTableService(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "ServeToTableError", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *OrderHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidOrderID", "Неверный ID заказа")
		return
	}

	var req dto.ProcessPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "InvalidRequest", "Неверный формат запроса")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, "ValidationError", err.Error())
		return
	}

	err = h.service.ProcessPayment(r.Context(), conn, orderID, req.Method, req.CardNumber, req.Expiry, req.CVV)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "PaymentError", err.Error())
		return
	}

	sendSuccess(w, http.StatusOK, map[string]interface{}{
		"message":  "Оплата прошла успешно",
		"order_id": orderID,
	})
}

func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "OrderIDNotFound", err.Error())
		return
	}
	if err := h.service.CancelOrder(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "CancelOrderError", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *OrderHandler) UndoLastAction(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	if err := h.service.UndoLastAction(r.Context(), conn); err != nil {
		sendError(w, http.StatusInternalServerError, "UndoLastActionError", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *OrderHandler) RedoLastAction(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	if err := h.service.RedoLastAction(r.Context(), conn); err != nil {
		sendError(w, http.StatusInternalServerError, "RedoLastActionError", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *OrderHandler) GetOrderTotal(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "OrderIDNotFound", err.Error())
		return
	}
	total, err := h.service.GetOrderTotal(r.Context(), conn, orderID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "GetOrderTotalError", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, total)
}
