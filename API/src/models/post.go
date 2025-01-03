package models

import "time"

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
