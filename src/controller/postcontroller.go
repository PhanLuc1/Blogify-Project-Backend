package controller

import (
	"encoding/json"
	"net/http"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
)

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	var userId int
	query := "SELECT * FROM post"
	result, err := database.Client.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for result.Next() {
		var post models.Post
		result.Scan(&post.Id, &userId, &post.Caption, &post.CreateAt)

		post.PostImages, err = models.GetImageProduct(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.User, err = models.GetInfoUser(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Reaction, err = models.GetReactionPost(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		posts = append(posts, post)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(posts)
}
