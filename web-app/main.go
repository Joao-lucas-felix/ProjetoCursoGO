package main

import (
	"log"
	"net/http"
	"web-app/src/router"
	"web-app/src/utils"
)

func main() {
	log.Println("Runing the DevBook Web App")
	utils.LoadTemplates()
	r := router.Gen()
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln(err)
	}
}
