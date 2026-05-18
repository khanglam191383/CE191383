package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func WriteSuccess(w http.ResponseWriter, data interface{}) {

	response := SuccessResponse{
		Success: true,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {

	response := ErrorResponse{
		Success: false,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}