package utils

import (
	"encoding/json"
	"net/http"
)

func ToJson(w http.ResponseWriter, body interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}
