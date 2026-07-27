package handler

import (
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TableHandler struct {
	service *service.TableService
}

func NewTableHandler(service *service.TableService) *TableHandler {
	return &TableHandler{service: service}
}

func (h *TableHandler) CreateTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	var req dto.CreateTableRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	table := req.ToDomain()
	created, err := h.service.CreateTableService(r.Context(), conn, table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.FromDomainTableResponse(*created)
	sendSuccess(w, http.StatusCreated, resp)
}

func (h *TableHandler) GetTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	table, err := h.service.GetTableService(r.Context(), conn, tableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.FromDomainTableResponse(*table)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) GetAllTables(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	tables, err := h.service.GetAllTablesService(r.Context(), conn, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.NewTableListResponse(tables, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) GetFreeTables(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	tables, err := h.service.GetFreeTablesService(r.Context(), conn, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.NewTableListResponse(tables, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) GetOccupiedTables(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	tables, err := h.service.GetOccupiedTablesService(r.Context(), conn, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.NewTableListResponse(tables, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) GetTablesByLocation(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	tables, err := h.service.GetFreeTablesService(r.Context(), conn, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.NewTableListResponse(tables, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) GetTablesByCapacity(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	capacity, _ := strconv.Atoi(r.URL.Query().Get("capacity"))

	tables, err := h.service.GetTablesByCapacityService(r.Context(), conn, limit, offset, capacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.NewTableListResponse(tables, 0, limit, offset)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) UpdateTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req dto.UpdateTableRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{})
	if req.Location != "" || len(req.Location) != 0 {
		updates["location"] = req.Location
	}
	if req.Number <= 0 {
		updates["number"] = req.Number
	}
	if req.Capacity <= 0 {
		updates["capacity"] = req.Capacity
	}

	table, err := h.service.UpdateTableService(r.Context(), conn, tableID, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.FromDomainTableResponse(*table)
	sendSuccess(w, http.StatusOK, resp)
}

func (h *TableHandler) DeleteTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.service.DeleteTableService(r.Context(), conn, tableID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *TableHandler) OccupyTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.service.OccupyTableService(r.Context(), conn, tableID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *TableHandler) FreeTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.service.FreeTableService(r.Context(), conn, tableID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *TableHandler) ReserveTable(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.service.ReserveTableService(r.Context(), conn, tableID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}

func (h *TableHandler) CancelReservation(w http.ResponseWriter, r *http.Request) {
	conn, ok := getConnOrError(w, r)
	if !ok {
		return
	}
	vars := mux.Vars(r)
	tableID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CancelReserveTableService(r.Context(), conn, tableID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendSuccess(w, http.StatusOK, nil)
}
