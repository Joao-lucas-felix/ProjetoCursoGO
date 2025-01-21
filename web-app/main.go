package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app/src/config"
	"web-app/src/cookies"
	"web-app/src/router"
	"web-app/src/utils"
)

func main() {

	config.Load()
	cookies.Config()

	log.Printf("Runing the DevBook Web App - In the Port: %d\n", config.Port)
	utils.LoadTemplates()
	r := router.Gen()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r); err != nil {
		log.Fatalln(err)
	}
}

// use that functions to generate your keys if you need
// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))

//		fmt.Println(hashKey)
//		fmt.Println(blockkey)
//	}
