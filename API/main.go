package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/config"
	"github.com/Joao-lucas-felix/DevBook/API/src/router"
)

func init(){
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil{
		log.Fatal(err)
	}
	strBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Print(strBase64)
}
func main() {
	config.LoadENV()
	fmt.Printf("\nRunning in the port:%d\nDatabase connection: %s\n", config.Port, config.DatabaseStrConnection)
	fmt.Println("-------Welcome to the DevBook API------")

	r := router.GenRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
