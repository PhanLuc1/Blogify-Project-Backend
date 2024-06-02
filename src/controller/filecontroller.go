package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
)

func ImageHandle(w http.ResponseWriter, r *http.Request) {
	avatar := r.URL.Query().Get("image")
	if avatar == "" {
		http.Error(w, "Missing avatar parameter", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("C:\\Users\\Admin\\Desktop\\image-blogify", avatar)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
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
