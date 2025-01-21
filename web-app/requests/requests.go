package requests

import (
	"fmt"
	"io"
	"net/http"
	"web-app/src/cookies"
)

func DoAuthenticateReuqest(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	reuqest, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	cookie, _ := cookies.GetToken(r)
	bearer := fmt.Sprintf("Bearer %s", cookie["token"])
	reuqest.Header.Add("Authorization", bearer)
	client := &http.Client{}
	return client.Do(reuqest)
}
