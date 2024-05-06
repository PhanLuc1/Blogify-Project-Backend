package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	var query string
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createAt := time.Now().Format("2006-01-02 15:04:05")
	if comment.ParentCommentID.Int64 != 0 {
		query = "INSERT INTO comment (userId, postId, parentCommentId, content, createAt) VALUES (?, ?, ?, ?, ?)"
		_, err = database.Client.Query(query, claims.UserId, comment.PostId, comment.ParentCommentID.Int64, comment.Content, createAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		query = "INSERT INTO comment (userId, postId, content, createAt) VALUES (?, ?, ?, ?)"
		_, err = database.Client.Query(query, claims.UserId, comment.PostId, comment.Content, createAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func CommentReact(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	commentId := r.URL.Query().Get("commentId")
	query := "INSERT INTO comment_reaction (userId, postId) VALUES (?, ?)"
	_, err = database.Client.Query(query, claims.UserId, commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
