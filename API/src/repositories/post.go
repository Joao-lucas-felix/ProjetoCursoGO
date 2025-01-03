package repositories

import (
	"database/sql"

	"github.com/Joao-lucas-felix/DevBook/API/src/models"
)

type Post struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Post {
	return &Post{db: db}
}

// CreatePost insert a new post in the database
func (repository Post) CreatePost(userId int, post models.Post) error {
	statement, err := repository.db.Prepare(`
		insert into post (title, content, author_id) values ($1, $2, $3)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(post.Title, post.Content, userId); err != nil {
		return err
	}

	return nil
}
