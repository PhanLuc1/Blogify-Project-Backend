package auth

import (
	"fmt"
	"net/http"

	generate "github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
)

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
