package models

import "time"

//Post represents a post made by a user
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	PosterID   uint64    `json:"posterId,omitempty"`
	PosterNick string    `json:"posterNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedOn  time.Time `json:"createdOn,omitempty"`
}
