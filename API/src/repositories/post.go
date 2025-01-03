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
// FindById find the post with the ID - X in the database 
func (repository Post) FindById(postId int64) (models.Post, error) {
	rows, err := repository.db.Query(`
				select p.id, p.title, p.content, u.id, u.nick, p.likes, p.created_at 
				from post p join usuarios u on p.author_id = u.id
				where p.id = $1
			`, postId)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()
	var post models.Post
	if rows.Next(){
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNick,
			&post.Likes,
			&post.CreatedAt,
		); err != nil{
			return models.Post{}, err
		}
	}
	return post, nil
}
