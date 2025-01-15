package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"web-app/src/responses"
)

// CreateUser creates a user in the API
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{
			Error: err.Error(),
		})
		return
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{
			Error: err.Error(),
		})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
