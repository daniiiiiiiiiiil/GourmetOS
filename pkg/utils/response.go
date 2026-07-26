package utils

import (
	"encoding/json"
	"net/http"
)

// используются не эти а в handler helper
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func SendError(w http.ResponseWriter, status int, message string) {
	SendJSON(w, status, Response{
		Success: false,
		Error:   message,
	})
}

func SendCreated(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func SendBadRequest(w http.ResponseWriter, message string) {
	SendError(w, http.StatusBadRequest, message)
}

func SendNotFound(w http.ResponseWriter, message string) {
	SendError(w, http.StatusNotFound, message)
}

func SendInternalError(w http.ResponseWriter, message string) {
	SendError(w, http.StatusInternalServerError, message)
}
