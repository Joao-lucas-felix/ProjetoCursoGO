package repositories

import (
	"database/sql"
	"errors"
	"fmt"

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
func (repository Users) Search(nameOrNick string) ([]models.User, error){
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	
	rows, err := repository.db.Query(
		"SELECT id, name, nick ,email, created_at, updated_at FROM usuarios u WHERE u.name LIKE $1 OR u.nick LIKE $2",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name,&user.Nick ,&user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}