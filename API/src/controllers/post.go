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
	"github.com/gorilla/mux"
)

// CreatePost is the endpoint to create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	var post models.Post
	if err := json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if err := json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := post.Prepare(); err != nil{
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewPostRepository(db)
	if err := repository.CreatePost(userId, post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK,
		struct {
			Message string
		}{
			Message: "Your Post are Created sucessfully",
		},
	)

}

// GetAllPost is the endpoint to get posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, err := repository.FindAll(userId)
	if err != nil{
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

// GetPostById is the enpoint to get a  post by id
func GetPostById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, err := strconv.Atoi(parameters["postId"])
	if err != nil{
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post, err :=  repository.FindById(int64(postId))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)

}

// UpdatePost is the endpoint to update a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost is the endpoint to delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
