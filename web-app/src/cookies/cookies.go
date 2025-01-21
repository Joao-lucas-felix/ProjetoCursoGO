package cookies

import (
	"net/http"
	"web-app/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config make the cookie configuration
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}
// Save saves the cookie with the user auth datas as encrypted http only cookie
func Save(w http.ResponseWriter, ID, token string)  error {
	data := map[string]string{
		"id": ID, 
		"token": token,
	}

	codedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name: "data",
		Value: codedData,
		Path: "/",
		HttpOnly: true,
	})
	
	return nil
}	

func GetToken(r *http.Request) ( map[string]string ,error ){
	values, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}
	cookie := make(map[string]string)
	if err:= s.Decode("data", values.Value, &cookie); err != nil {
		return nil, err
	}
	return cookie, nil

}
