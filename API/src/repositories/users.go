package repositories

import (
	"database/sql"
	"errors"

	"github.com/Joao-lucas-felix/DevBook/API/src/models"
)

// User represents an repository to interact with the users table in the database
type Users struct {
	db *sql.DB
}



// NewUserRepository recivies a DB connection and returns a new User repository 
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db: db}
}
// Create a new user in the database
func (repository Users) Create(user models.User) (error) {
	statment, err := repository.db.Prepare("INSERT INTO usuarios (name, nick, email, password_hash) VALUES ($1, $2, $3, $4);")
	if err != nil {
		return err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Nick, user.Email ,user.Password)
	if err != nil{
		return err
	}
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}  
	if rowsAffect != 1{
		return errors.New("error while trying to create an user")
	}
	return nil
}