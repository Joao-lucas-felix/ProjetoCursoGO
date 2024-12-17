package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/config"
	"github.com/Joao-lucas-felix/DevBook/API/src/router"
)

func main() {
	config.LoadENV()
	fmt.Printf("Running in the port:%d\nDatabase connection: %s\n", config.Port, config.DatabaseStrConnection)
	fmt.Println("-------Welcome to the DevBook API------")

	r := router.GenRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
