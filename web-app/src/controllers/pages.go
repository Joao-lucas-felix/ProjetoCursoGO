package controllers

import (
	"fmt"
	"log"
	"net/http"
	"web-app/requests"
	"web-app/src/config"
	"web-app/src/utils"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

// LoadCreateUserPage loads the page to create a new user
func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "create-user.html", nil)
}

// LoadHomePage loads the home  page
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.DoAuthenticateReuqest(r, http.MethodGet, url, nil)

	log.Println(response.StatusCode, err)
	utils.ExecTemplate(w, "home.html", nil)
}
