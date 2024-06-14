package models

import "github.com/PhanLuc1/Blogify-Project-Backend/src/database"

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Biography   string `json:"biography"`
	State       bool   `json:"state"`
	AvatarImage string `json:"avatarImage"`
}
type Follower struct {
	Id         int
	FollowerID int `json:"followerId"`
	FolloweeID int `json:"folloeeID"`
}
type AnotherUser struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Followers   int    `json:"follower"`
	Following   int    `json:"following"`
	AvatarImage string `json:"avatarImage"`
	State       bool   `json:"state"`
}

type Token struct {
	Token string `json:"token"`
}

func GetInfoUser(userId int) (user User, err error) {
	err = database.Client.QueryRow("SELECT user.id, user.email, user.username, user.biography ,user.state, user.avatarImage FROM user WHERE id = ?", userId).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Biography,
		&user.State,
		&user.AvatarImage,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
func GetStateUser(userId int) (state bool, err error) {
	err = database.Client.QueryRow("SELECT user.state FROM user WHERE id = ?", userId).Scan(
		&state,
	)
	if err != nil {
		return state, err
	}
	return state, nil
}

type Response struct {
	TokenUser Token `json:"tokenUser"`
	User      User  `json:"user"`
}
