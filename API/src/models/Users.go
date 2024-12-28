package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the aplication
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Prepare Validates and Prepare the user to persist in the database
func (u *User) Prepare(step string) error {
	err := u.validate(step)
	if err != nil {
		return err
	}
	u.formatFields()
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is a required parameter must not be blank")
	}
	if u.Nick == "" {
		return errors.New("nick is a required parameter must not be blank")
	}
	if u.Email == "" {
		return errors.New("email is a required parameter must not be blank")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil{
		errorMsg := fmt.Sprintf("email is in invalid format: %s", err.Error())
		return errors.New(errorMsg)
	}
	if u.Password == "" && step == "create"{
		return errors.New("password is a required parameter must not be blank")
	}
	return nil
}
func (u *User) formatFields() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
