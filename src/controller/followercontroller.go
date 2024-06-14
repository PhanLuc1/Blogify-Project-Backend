package controller

import (
	"net/http"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	"github.com/gorilla/mux"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	claims, err := auth.GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	query := "INSERT INTO follower (followerId, followedId) VALUES(?, ?)"
	_, err = database.Client.Query(query, claims.UserId, userid)
	if err != nil {
		query = "DELETE FROM follower WHERE followerId = ? AND followedId = ?"
		_, err = database.Client.Query(query, claims.UserId, userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
}
