package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}){
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error){
	JSON(w, statusCode,
		 struct{
			Erro string `json:"error"`
		}{
			Erro: erro.Error(),
		},
	)
}