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
func (repository Users) Create(user models.User) error {
	statment, err := repository.db.Prepare("INSERT INTO usuarios (name, nick, email, password_hash) VALUES ($1, $2, $3, $4);")
	if err != nil {
		return err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffect != 1 {
		return errors.New("error while trying to create an user")
	}
	return nil
}

// Search returns one user by name or nick
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
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
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById returns a user by the id
func (repository Users) FindById(userId int) (models.User, error) {

	rows, err := repository.db.Query("SELECT id, name, email, nick, created_at, updated_at FROM usuarios u WHERE id = $1", userId)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Nick, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Update updates the user in the database
func (repository Users) Update(userId int, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE usuarios SET name = $1, nick = $2, email = $3 WHERE id = $4",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.Name, user.Nick, user.Email, userId)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an user in the database
func (repository Users) Delete(userId int) error {
	statement, err := repository.db.Prepare("DELETE FROM usuarios u WHERE u.id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(userId); err != nil {
		return err
	}
	return nil
}

// Find by email finds a user by email and returns the user Id and the password hash
func (repository Users) FindByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("SELECT id, password_hash FROM usuarios u WHERE u.email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// FollowUser adds the follow relation in the databse
func (repository Users) FollowUser(userId, followerId int) error {
	statement, err := repository.db.Prepare("insert into seguidores (usuario_id, seguidor_id) values ($1,$2);")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followerId, userId); err != nil {
		return err
	}

	return nil
}

// FollowUser unfollow a user
func (repository Users) UnfollowUser(userId, followerId int) error {
	statement, err := repository.db.Prepare("delete from seguidores where usuario_id = $1 and seguidor_id = $2;")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followerId, userId); err != nil {
		return err
	}

	return nil
}
// Get all followers of an user
func (repository Users) GetFollowers(userId int) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.created_at, u.updated_at FROM usuarios u JOIN seguidores s ON s.seguidor_id = u.id WHERE s.usuario_id=$1;`, 
		userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil{
			return nil, err
		}
		followers = append(followers, user)
	}

	return followers, nil
}

// Get all following of an user
func (repository Users) GetFollowing(userId int) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.created_at, u.updated_at FROM usuarios u JOIN seguidores s ON s.usuario_id = u.id WHERE s.seguidor_id=$1;`, 
		userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil{
			return nil, err
		}
		followers = append(followers, user)
	}

	return followers, nil
}
