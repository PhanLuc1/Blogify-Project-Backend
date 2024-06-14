package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

		post.PostImages, err = models.GetImagePost(post.Id)
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
		post.CountComment, err = models.GetAmountCommentPost(post.Id)
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
func UploadeHandle(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	caption := r.FormValue("caption")
	post := models.Post{
		User:     models.User{Id: claims.UserId},
		Caption:  caption,
		CreateAt: time.Now(),
	}
	var postImages []models.PostImage
	files := r.MultipartForm.File["images"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileName := filepath.Join("C:\\Users\\Admin\\Desktop\\image-blogify", fileHeader.Filename)
		dst, err := os.Create(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		postImages = append(postImages, models.PostImage{
			ImageURL:    fileHeader.Filename,
			Description: "",
		})
	}
	post.PostImages = postImages

	query := "INSERT INTO post (userId, caption, createAt) VALUES (?, ? ,? )"
	result, err := database.Client.Exec(query, post.User.Id, post.Caption, post.CreateAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, img := range post.PostImages {
		query = "INSERT INTO postimage (imageURL, description, postId) VALUES (?, ? ,?)"
		_, err := database.Client.Exec(query, img.ImageURL, img.Description, postId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
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
		query = "DELETE FROM reaction WHERE reaction.userId = ? AND reaction.postId = ?"
		_, err = database.Client.Query(query, claims.UserId, postId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
	post.PostImages, err = models.GetImagePost(post.Id)
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
	post.CountComment, err = models.GetAmountCommentPost(post.Id)
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
func GetPostsByUserId(w http.ResponseWriter, userId int) {
	var posts []models.Post
	query := "SELECT post.id, post.caption, post.createAt FROM post WHERE userId = ?"
	result, err := database.Client.Query(query, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for result.Next() {
		var post models.Post
		err = result.Scan(
			&post.Id,
			&post.Caption,
			&post.CreateAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post.User, err = models.GetInfoUser(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.PostImages, err = models.GetImagePost(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.CountComment, err = models.GetAmountCommentPost(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments, err = models.GetCommentsForPost(post.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(posts)
}
func GetCurrentUserPosts(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	GetPostsByUserId(w, claims.UserId)
}
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	userId, err := strconv.Atoi(userid)
	if err != nil {
		panic(err)
	}

	GetPostsByUserId(w, userId)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	query := "DELETE FROM comment WHERE postId = ?"
	_, err = database.Client.Query(query, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query = "DELETE FROM reaction WHERE postId = ?"
	_, err = database.Client.Query(query, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query = "DELETE FROM postimage WHERE postId = ?"
	_, err = database.Client.Query(query, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query = "DELETE FROM post WHERE id = ? AND userId = ?"
	_, err = database.Client.Query(query, postId, claims.UserId)
	if err != nil {
		http.Error(w, "You do not have permission to delete this post", http.StatusForbidden)
		return
	}
	w.WriteHeader(200)
}
