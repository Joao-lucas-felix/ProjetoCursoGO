package responses

import (
	"encoding/json"
	"log"
	"net/http"
)
// Error represents the error response from API
type ErrorAPI struct{
	Error string `json:"error"`
}
// JSON returns a new json to a request
func JSON(w http.ResponseWriter, statusCode int, data interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err)
	}
}
// ErrorCode handle the responses that status doce greater than 400
func ErrorCode(w http.ResponseWriter, r *http.Response){
	var error ErrorAPI
	json.NewDecoder(r.Body).Decode(&error)
	JSON(w, r.StatusCode, error)
}