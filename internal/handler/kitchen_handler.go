package handler

import (
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KitchenHandler struct {
	service *service.KitchenService
}

func NewKitchenHandler(service *service.KitchenService) *KitchenHandler {
	return &KitchenHandler{
		service: service,
	}
}

func (h *KitchenHandler) ReceiveOrder(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Не правильный ID")
		return
	}

	if err := h.service.ReceiveOrder(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}

	sendSuccess(w, http.StatusCreated, nil)
}

func (h *KitchenHandler) StartCooking(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Не правильный ID")
		return
	}
	if err := h.service.StartCooking(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *KitchenHandler) MarkAsReady(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Не правильный ID")
		return
	}
	if err := h.service.MarkAsReady(r.Context(), conn, orderID); err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusCreated, nil)
}

func (h *KitchenHandler) GetCookingQueue(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	cooking, total, err := h.service.GetCookingQueue(r.Context(), conn, limit, offset)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}
	resp := dto.NewKitchenQueueResponse(cooking, limit, offset, total)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *KitchenHandler) GetCookingTime(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	dishID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidID", "Не правильный ID")
		return
	}

	time, err := h.service.GetCookingTime(r.Context(), conn, dishID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, time)
}

func (h *KitchenHandler) GetKitchenStatus(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}

	status, err := h.service.GetKitchenStatus(r.Context(), conn)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InvalidServer", err.Error())
		return
	}
	sendSuccess(w, http.StatusOK, status)
}
