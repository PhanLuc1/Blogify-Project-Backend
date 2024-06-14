package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	generate "github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userpassword string, givenpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenpassword), []byte(userpassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Login Or Passowrd is Incorerct"
		valid = false
	}
	return valid, msg
}
func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user.Email) {
		w.WriteHeader(422)
		w.Write([]byte(`{"error": "Invalid email format"}`))
		return
	}

	password := HashPassword(user.Password)
	user.Password = password
	query := "INSERT INTO user (email, username, password, biography, avatarImage) VALUES (?, ?, ?, ?, ?)"
	_, err = database.Client.Query(query, user.Email, user.Username, user.Password, " ", user.AvatarImage)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(`{"message": "Your account has been created"}`))
}
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var foundUser models.User
	var tokenUser models.Token
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = database.Client.QueryRow("SELECT * FROM user WHERE email =?", user.Email).Scan(
		&foundUser.Id,
		&foundUser.Email,
		&foundUser.Username,
		&foundUser.Password,
		&foundUser.Biography,
		&foundUser.State,
		&foundUser.AvatarImage,
	)
	if err != nil {
		w.WriteHeader(404)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	PasswordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
	if !PasswordIsValid {
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	token, _ := generate.TokenGeneration(foundUser.Id)
	tokenUser.Token = token

	foundUser.Password = ""

	response := models.Response{
		TokenUser: tokenUser,
		User:      foundUser,
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	var user models.User
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err = models.GetInfoUser(claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	setClauses := []string{}
	args := []interface{}{}

	if user.Username != "" {
		setClauses = append(setClauses, "username = ?")
		args = append(args, user.Username)
	}
	if user.AvatarImage != "" {
		setClauses = append(setClauses, "avatarImage = ?")
		args = append(args, user.AvatarImage)
	}

	if len(setClauses) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("UPDATE user SET %s WHERE id = ?", strings.Join(setClauses, ", "))
	args = append(args, claims.UserId)

	_, err = database.Client.Exec(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	userId, err := strconv.Atoi(userid)
	if err != nil {
		panic(err)
	}
	_, err = auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err := models.GetInfoUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE user SET "
	params := []interface{}{}
	if user.Biography != "" {
		query += "biography = ?, "
		params = append(params, user.Biography)
	}
	if user.Username != "" {
		query += "username = ?, "
		params = append(params, user.Username)
	}

	if len(params) > 0 {
		query = query[:len(query)-2]
		query += " WHERE id = ?"
		params = append(params, claims.UserId)
	} else {
		http.Error(w, "No valid fields to update", http.StatusBadRequest)
		return
	}

	_, err = database.Client.Exec(query, params...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile updated successfully"))
}
func UploadAvatarImage(w http.ResponseWriter, r *http.Request) {
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

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		http.Error(w, "No image file uploaded", http.StatusBadRequest)
		return
	}

	fileHeader := files[0]
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

	user := models.User{
		Id:          claims.UserId,
		AvatarImage: fileHeader.Filename,
	}

	query := "UPDATE user SET avatarImage = ? WHERE id = ?"
	_, err = database.Client.Exec(query, user.AvatarImage, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Avatar image uploaded successfully"))
}
func SetUpStateAccount(w http.ResponseWriter, r *http.Request) {
	var query string
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	state, err := models.GetStateUser(claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if state {
		query = "UPDATE user SET state = 0 WHERE id = ?"
	} else {
		query = "UPDATE user SET state = 1 WHERE id = ?"
	}
	_, err = database.Client.Query(query, claims.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
