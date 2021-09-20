package models

import (
	"errors"
	"strings"
	"time"
)

//Post represents a post made by a user in the social media
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	PosterID   uint64    `json:"posterId,omitempty"`
	PosterNick string    `json:"posterNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedOn  time.Time `json:"createdOn,omitempty"`
}

//Prepare will call the methods to validate and format the received post
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the post must have a title and it cannot be empty")
	}
	if post.Content == "" {
		return errors.New("the post must have a content and it cannot be empty")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
