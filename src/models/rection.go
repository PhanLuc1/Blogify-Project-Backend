package models

import "github.com/PhanLuc1/Blogify-Project-Backend/src/database"

type Reaction struct {
	UserS         []User `json:"users"`
	CountReaction int    `json:"countReaction"`
}

func GetReactionPost(postId int) (Reaction, error) {
	var users []User
	CountReaction := 0
	query := "SELECT userId FROM reaction WHERE postId = ?"
	result, err := database.Client.Query(query, postId)
	if err != nil {
		return Reaction{}, err
	}
	for result.Next() {
		var userId int
		result.Scan(&userId)
		user, err := GetInfoUser(userId)
		if err != nil {
			return Reaction{} , err 
		}
		users = append(users, user)
		CountReaction ++
	}
	reaction := Reaction{users, CountReaction}
	return reaction, nil
}
