package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Failed to encode response:", err)
	}
	return err
}
