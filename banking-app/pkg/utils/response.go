package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func WriteJSON[T any](w http.ResponseWriter, status int, data T, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(ApiResponse[T]{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(ApiResponse[T]{
		Success: true,
		Data:    data,
	})
}
