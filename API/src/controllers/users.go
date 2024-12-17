package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/database"
	"github.com/Joao-lucas-felix/DevBook/API/src/models"
	"github.com/Joao-lucas-felix/DevBook/API/src/repositories"
	"github.com/Joao-lucas-felix/DevBook/API/src/responses"
)

// Create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil{
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if err = repository.Create(user); err != nil{
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, "The user are created Successfully!")

}

// Get all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

// Get an user with ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get an User"))
}

// Update an user with ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an user"))
}

// Delete an user with ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete an user"))
}
