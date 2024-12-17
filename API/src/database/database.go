package database

import (
	"database/sql"

	"github.com/Joao-lucas-felix/DevBook/API/src/config"
	_ "github.com/lib/pq"
)

// Connect to the Database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DatabaseStrConnection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}