package controllers

import (
	"encoding/json"
	"errors"
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

	if err := post.Prepare(); err != nil {
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
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

// GetPostById is the enpoint to get a  post by id
func GetPostById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, err := strconv.Atoi(parameters["postId"])
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

	repository := repositories.NewPostRepository(db)
	post, err := repository.FindById(int64(postId))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)

}

// UpdatePost is the endpoint to update a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	postId, err := strconv.Atoi(parameters["postId"])
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

	repository := repositories.NewPostRepository(db)
	postInDatabe, err := repository.FindById(int64(postId))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if postInDatabe.AuthorId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("you can not edit a post of anoter user"))
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
	if err := post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.UpdatePost(int64(postId), post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK,
		struct {
			Message string
		}{
			Message: "Your Post are Updated sucessfully",
		},
	)
}

// DeletePost is the endpoint to delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	postId, err := strconv.Atoi(parameters["postId"])
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

	repository := repositories.NewPostRepository(db)
	postInDatabe, err := repository.FindById(int64(postId))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if postInDatabe.AuthorId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("you can not delete a post of anoter user"))
		return
	}
	if err := repository.DeletePost(int64(postId)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK,
		struct {
			Message string
		}{
			Message: "Your Post are Deleted sucessfully",
		},
	)
}

// GetUserPosts is the endpoint that returns all posts of a specific user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.Atoi(parameters["userId"])
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
	repository := repositories.NewPostRepository(db)
	posts, err := repository.FindPostsByUser(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, err := strconv.Atoi(parameters["postId"])
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

	repository := repositories.NewPostRepository(db)
	if err := repository.LikePost(int64(postId)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK,
		struct {
			Message string
		}{
			Message: "Your Like the Post sucessfully",
		},
	)

}
