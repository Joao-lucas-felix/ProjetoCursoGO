package models
// Password represents the request body to update a user password
type Password struct {
	NewPassword string `json:"new_password"`
	Password string `json:"password"`
}