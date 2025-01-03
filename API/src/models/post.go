package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post of an user
type Post struct {
	Id         int64     `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   int       `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
// Prepare validate and format the post before save or update
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil{
		return err
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title can not be blank")
	}
	if post.Content == "" {
		return errors.New("the content can not be blank")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
