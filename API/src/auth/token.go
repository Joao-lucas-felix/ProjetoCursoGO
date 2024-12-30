package auth

import (
	"time"

	"github.com/Joao-lucas-felix/DevBook/API/src/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(usuerId int) (string, error) {
	permitions := jwt.MapClaims{}

	permitions["authorized"] = true
	permitions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permitions["user_id"] = usuerId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)
	return token.SignedString(config.SecretKey)

}
