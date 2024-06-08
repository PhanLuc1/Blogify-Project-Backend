package models

import (
	"time"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
)

type Post struct {
	Id           int         `json:"id"`
	User         User        `json:"user"`
	Caption      string      `json:"caption"`
	PostImages   []PostImage `json:"postImages"`
	CreateAt     time.Time   `json:"createAt"`
	CountComment int         `json:"countCommet"`
	Comments     []Comment   `json:"comments"`
	Reaction     Reaction    `json:"reaction"`
}
type PostImage struct {
	Id          int    `json:"id"`
	ImageURL    string `json:"imageURL"`
	Description string `json:"Description"`
}

func GetImagePost(postId int) (postImages []PostImage, err error) {
	result1, err := database.Client.Query("SELECT postimage.id, postimage.imageURL, postimage.description FROM postimage WHERE postId = ?", postId)
	if err != nil {
		return nil, err
	}
	for result1.Next() {
		var imageTemp PostImage
		err := result1.Scan(&imageTemp.Id, &imageTemp.ImageURL, &imageTemp.Description)
		if err != nil {
			return nil, err
		}
		postImages = append(postImages, imageTemp)
	}
	return postImages, nil
}
