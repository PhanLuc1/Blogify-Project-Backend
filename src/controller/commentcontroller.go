package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	var query string
	vars := mux.Vars(r)
	postid := vars["postid"]
	postId, err := strconv.Atoi(postid)
	if err != nil {
		panic(err)
	}
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
		_, err = database.Client.Query(query, claims.UserId, postId, comment.ParentCommentID.Int64, comment.Content, createAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		query = "INSERT INTO comment (userId, postId, content, createAt) VALUES (?, ?, ?, ?)"
		_, err = database.Client.Query(query, claims.UserId, postId, comment.Content, createAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	comments, err := models.GetCommentsForPost(postId, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(comments)
}
func CommentReact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	query := "INSERT INTO comment_reaction (userId, comentId) VALUES (?, ?)"
	_, err = database.Client.Query(query, claims.UserId, commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var postId int
	err = database.Client.QueryRow("SELECT postId FROM comment WHERE id = ?", commentId).Scan(&postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query := "DELETE FROM comment_reaction WHERE commentId = ?"
	_, err = database.Client.Query(query, commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query = "DELETE FROM comment WHERE id = ? AND userId = ?"
	_, err = database.Client.Query(query, commentId, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	comments, err := models.GetCommentsForPost(postId, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comments)
}
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	//commentId := vars["commentid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if comment.Content == "" {
		w.WriteHeader(400)
		return
	}
	query := "UPDATE comment SET comment.content = ? WHERE id = ? AND userId = ?"
	_, err = database.Client.Query(query, comment.Content, comment.Id, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	var postId int
	err = database.Client.QueryRow("SELECT postId FROM comment WHERE id = ?", comment.Id).Scan(&postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := models.GetCommentsForPost(postId, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comments)
}
