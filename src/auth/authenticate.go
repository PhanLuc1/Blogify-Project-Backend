package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	generate "github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
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
func GetNewTokenFromRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshtToken := r.Header.Get("RefreshToken")
	if refreshtToken != "" {
		claims, _ := generate.ValidateToken(refreshtToken)
		if claims == nil {
			w.WriteHeader(401)
			return
		}
		var tokenUser models.Token
		token, refreshToken, _ := generate.TokenGeneration(claims.UserId)
		tokenUser.Token = token
		tokenUser.Refreshtoken = refreshToken
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(token)
	}
}
