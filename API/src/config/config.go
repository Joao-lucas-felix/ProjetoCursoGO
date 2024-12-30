package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//DatabaseStrConnection Str of Database connection
	DatabaseStrConnection = ""
	// Port API running port
	Port = 0
	// SecretKey is the key used to signed the JWT
	SecretKey []byte
)

// LoadENV load the env variables
func LoadENV() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error while trying to up the API:\n", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_DEFAULT_PORT"))
	if err != nil {
		Port = 9000
	}
	//"user=golang dbname=devbook password=golang host=localhost port=5432 sslmode=disable"
	DatabaseStrConnection = fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE_DEFAULT"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
