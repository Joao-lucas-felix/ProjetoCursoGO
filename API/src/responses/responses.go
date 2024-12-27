package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON render a JSON message
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error render a JSON error message
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode,
		struct {
			Erro string `json:"error"`
		}{
			Erro: erro.Error(),
		},
	)
}
