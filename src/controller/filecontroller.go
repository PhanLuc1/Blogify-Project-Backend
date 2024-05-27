package controller

import (
	"net/http"
	"os"
	"path/filepath"
)

func AvatarHandle(w http.ResponseWriter, r *http.Request) {
	avatar := r.URL.Query().Get("avatar")
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
