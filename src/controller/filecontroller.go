package controller

import (
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
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
