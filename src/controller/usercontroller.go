package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	generate "github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
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
	query := "INSERT INTO user (email, username, password, avataImage) VALUES (?, ?, ?, ?)"
	_, err = database.Client.Query(query, user.Email, user.Username, user.Password, user.AvataImage)
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

	err = database.Client.QueryRow("SELECT user.id,user.email, user.password FROM user WHERE email =?", user.Email).Scan(
		&foundUser.Id,
		&foundUser.Email,
		&foundUser.Password,
	)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"Message": "Email is not avilable"}`))
		return
	}

	PasswordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
	if !PasswordIsValid {
		http.Error(w, msg, 401)
		return
	}
	token, refreshToken, _ := generate.TokenGeneration(foundUser.Id)
	tokenUser.Token = token
	tokenUser.Refreshtoken = refreshToken

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tokenUser)
}
