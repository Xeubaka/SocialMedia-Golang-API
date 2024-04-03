package models

import (
	"errors"
	"strings"
	"time"
)

// Posts represents a post made by an user
type Post struct {
	ID         uint64    `json: "id, omitempty"`
	Title      string    `json: "title, omitempty"`
	Content    string    `json: "content, omitempty"`
	AuthorID   uint64    `json: "authorId, omitempty"`
	AuthorNick string    `json: "authorNick, omitempty"`
	Likes      uint64    `json: "likes"`
	CreatedAt  time.Time `json: "createdAt, omitempty"`
}

// Prepare post for database insertion
func (post *Post) Prepare() (err error) {
	if err = post.validate(); err != nil {
		return
	}

	post.format()
	return
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New(FieldisEmptyMessage("title"))
	}

	if post.Content == "" {
		return errors.New(FieldisEmptyMessage("content"))
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
