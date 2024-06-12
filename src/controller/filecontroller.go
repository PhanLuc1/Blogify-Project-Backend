package controller

import (
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func ImageHandle(w http.ResponseWriter, r *http.Request) {
	imageURL := r.URL.Query().Get("image")
	if imageURL == "" {
		http.Error(w, "Missing imageURL parameter", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("C:\\Users\\Admin\\Desktop\\image-blogify", imageURL)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
}
