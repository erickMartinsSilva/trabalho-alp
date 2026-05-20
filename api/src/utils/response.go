package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, status_code int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	json.NewEncoder(w).Encode(body)
}