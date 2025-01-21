package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//ApiUrl represents the url to use the API
	APIURL = ""
	// Port represents the port where this service is running
	Port = 0
	// HashKey is used to authenticate the cookie
	HashKey []byte
	// BlockKey is used to encript the cookie
	BlockKey []byte
)

// Load loads the env var
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
