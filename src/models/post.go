package models

import "time"

type Post struct {
	Id         int         `json:"id"`
	User       User        `json:"user"`
	Caption    string      `json:"caption"`
	PostImages []PostImage `json:"postImages"`
	CreateAt   time.Time   `json:"createAt"`
	Comments   []Comment   `json:"comments"`
}
type PostImage struct {
	Id       int    `json:"id"`
	ImageURL string `json:"imageURL"`
}
type Comment struct {
	Id            int       `json:"id"`
	User          User      `json:"user"`
	PostId        int       `json:"postId"`
	Content       string    `json:"content"`
	CreateAt      time.Time `json:"creatAt"`
	ParentComment *Comment  `json:"parentComment"`
	Reaction      Reaction  `json:"reaction"`
}
type Reaction struct {
	Id            int    `json:"id"`
	UserS         []User `json:"users"`
	CountReaction int    `json:"countReaction"`
}
