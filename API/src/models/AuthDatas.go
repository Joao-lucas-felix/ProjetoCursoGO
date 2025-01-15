package models
// AuthDatas contains the id and a token to a user
type AuthDatas struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
