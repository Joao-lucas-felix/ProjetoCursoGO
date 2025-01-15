package models
// AuthDatas cotains the id and a token to a authenticated user
type AuthDatas struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
