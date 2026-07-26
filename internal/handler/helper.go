package handler

import (
	"GourmetOS/internal/handler/dto"
	"GourmetOS/internal/middleware"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func getConnOrError(w http.ResponseWriter, r *http.Request) (*pgx.Conn, bool) {
	conn, err := middleware.GetConnFromContext(r)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "db_error", err.Error())
		return nil, false
	}
	return conn, true
}

func sendError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto.NewErrorResponse(code, message))
}

func sendSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto.NewSuccessResponse(data))
}
