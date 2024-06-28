package models

import "time"

type Post struct {
	Id          string    `json:"id"` //para cuando se serialice el json
	PostContent string    `json:"postContent"`
	CreatedAt   time.Time `json:"createdAt"`
	UserId      string    `json:"userId"`
}
