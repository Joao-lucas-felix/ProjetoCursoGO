package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Joao-lucas-felix/DevBook/API/src/auth"
	"github.com/Joao-lucas-felix/DevBook/API/src/database"
	"github.com/Joao-lucas-felix/DevBook/API/src/models"
	"github.com/Joao-lucas-felix/DevBook/API/src/repositories"
	"github.com/Joao-lucas-felix/DevBook/API/src/responses"
	"github.com/Joao-lucas-felix/DevBook/API/src/security"
)

// Login is the func that athenticated the user in the API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	savedInDbUser, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err := security.VerifyPassword(user.Password, savedInDbUser.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	token, err := auth.CreateToken(int(savedInDbUser.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	userId := strconv.FormatUint(savedInDbUser.ID, 10)

	responses.JSON(w, http.StatusOK, 
		models.AuthDatas{
		ID: userId,
		Token: token,
	})
}
