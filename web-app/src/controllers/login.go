package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web-app/src/config"
	"web-app/src/cookies"
	"web-app/src/models"
	"web-app/src/responses"
)

// Login uses the email and password to get the token in the API
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorCode(w, response)
	}

	var authData models.AuthDatas

	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}
	if err := cookies.Save(w, authData.ID, authData.Token ); err != nil{
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
