package auth

import (
	"fmt"
	"log"
	"net/http"

	generate "github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func GetUserFromToken(r *http.Request) (*generate.SignedDetails, error) {
	token := r.Header.Get("Token")
	if token != "" {
		claims, msg := generate.ValidateToken(token)
		if claims == nil {
			return nil, fmt.Errorf(msg)
		}
		return claims, nil
	}
	return nil, fmt.Errorf("token not provided")
}
