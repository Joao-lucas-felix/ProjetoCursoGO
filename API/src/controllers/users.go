package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Joao-lucas-felix/DevBook/API/src/auth"
	"github.com/Joao-lucas-felix/DevBook/API/src/database"
	"github.com/Joao-lucas-felix/DevBook/API/src/models"
	"github.com/Joao-lucas-felix/DevBook/API/src/repositories"
	"github.com/Joao-lucas-felix/DevBook/API/src/responses"
	"github.com/Joao-lucas-felix/DevBook/API/src/security"
)

// CreateUser Create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
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

	if err = repository.Create(user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

// GetAllUser Get all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

// GetUser Get an user with ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userID"])
	if err != nil {
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
	user, err := repository.FindById(userId)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser  Update  user with ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userID"])
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	userRequestId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userRequestId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("forbiden to update anoter user"))
		return
	}

	fmt.Println(userRequestId)
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := user.Prepare("update"); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Update(userId, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "The user are updated successfully!")

}

// DeleteUser Delete a user with ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userID"])
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRequestId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userRequestId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("forbiden to delete anoter user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "The User are deleted Successfuly")
}

// FollowUser thats function is used to start to following another user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userID"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if userId == followerId {
		responses.Error(w, http.StatusForbidden, errors.New("you can not follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.FollowUser(userId, followerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK,
		struct{
			Message string
		}{
			Message: "You started to follow the user sucessfully",
		},
	)

}

// FollowUser thats function is used to unfollow a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userID"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if userId == followerId {
		responses.Error(w, http.StatusForbidden, errors.New("you can not unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.UnfollowUser(userId, followerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK,
		struct{
			Message string
		}{
			Message: "You unfollow the user sucessfully",
		},
	)

}
// GetFollowers find and return all followers of an user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userID"])
	if err != nil {
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
	followers, err := repository.GetFollowers(userId)
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)

}
// GetFollowers find and return all followers of an user
func GetFollowings(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userID"])
	if err != nil {
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
	followers, err := repository.GetFollowing(userId)
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)

}
// RedifinePassword is the endpoint to redifine a user password
func RedifinePassword(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userID"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if userIdInToken != userId {
		responses.Error(w, http.StatusForbidden, errors.New("impossible to update anoter user"))
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	var password models.Password
	if err := json.Unmarshal(requestBody, &password); err != nil{
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
	passwordSavedInDb, err := repository.GetPasswordById(userId)
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err := security.VerifyPassword(password.Password, passwordSavedInDb); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	newPasswordHash, err := security.Hash(password.NewPassword)
	if err != nil{
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.UpdatePassword(userId, string(newPasswordHash)); err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	
	responses.JSON(w, http.StatusOK,
		struct{
			Message string
		}{
			Message: "You password are updated sucessfully",
		},
	)
}