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
	if rows.Next() {
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNick,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

// FindAll find all posts based in the users that the user in request followings and returns the user posts
func (repository Post) FindAll(userId int) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.id, p.title, p.content, u.id, u.nick, p.likes, p.created_at 
		FROM post p JOIN seguidores s  on p.author_id = s.usuario_id
		JOIN usuarios u ON p.author_id = u.id
		WHERE seguidor_id = $1 OR u.id = $2
		order by p.created_at desc;
	`, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNick,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
// UpdatePost update a tile and content of a post 
func (repository Post) UpdatePost(postId int64, post models.Post) error {
	statement, err := repository.db.Prepare(`
		update post set title = $1, content = $2 where id = $3
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}
	return nil
}
// DeletePost delete a post of an user
func (repository Post) DeletePost(postId int64) error {
	statement, err := repository.db.Prepare(`
		delete from post where id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(postId); err != nil {
		return err
	} 
	return nil
}
// FindPostsByUser find all posts of a user
func (repository Post) FindPostsByUser(userId int) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.id, p.title, p.content, u.id, u.nick, p.likes, p.created_at 
		FROM post p JOIN usuarios u ON p.author_id = u.id
		WHERE p.author_id = $1
		order by p.created_at desc;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNick,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}