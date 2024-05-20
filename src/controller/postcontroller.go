package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
		post.Comments, err = models.GetCommentsForPost(post.Id)
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
func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	creatAt := time.Now().Format("2006-01-02 15:04:05")
	query := "INSERT INTO post (post.userId, caption, createAt) VALUES (?, ?, ?)"
	result, err := database.Client.Exec(query, claims.UserId, post.Caption, creatAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Id = int(lastid)
	for _, image := range post.PostImages {
		query = "INSERT INTO postimage (imageURL, description, postId) VALUES (?, ? ,?)"
		_, err = database.Client.Query(query, image.ImageURL, image.Description, post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"message": "created post"`))
}
func PostReact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	query := "INSERT INTO reaction (userId, postId) VALUES (?, ?)"
	_, err = database.Client.Query(query, claims.UserId, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postid"]
	var userId int
	var post models.Post
	query := "SELECT * FROM post WHERE id = ?"
	err := database.Client.QueryRow(query, postId).Scan(
		&post.Id,
		&userId,
		&post.Caption,
		&post.CreateAt,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	post.Comments, err = models.GetCommentsForPost(post.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Reaction, err = models.GetReactionPost(post.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(post)
}
