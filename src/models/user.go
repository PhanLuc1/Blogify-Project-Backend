package models

import "github.com/PhanLuc1/Blogify-Project-Backend/src/database"

type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	State      bool   `json:"state"`
	AvataImage string `json:"avataImage"`
}
type Follower struct {
	Id         int
	FollowerID int `json:"followerId"`
	FolloweeID int `json:"folloeeID"`
}
type VirtualUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	State      bool   `json:"state"`
	AvataImage string `json:"avataImage"`
}

type Token struct {
	Token        string `json:"token"`
	Refreshtoken string `json:"refreshToken"`
}

func GetInfoUser(userId int) (user User, err error) {
	err = database.Client.QueryRow("SELECT user.id, user.email, user.username, user.state, user.avataImage FROM user WHERE id = ?", userId).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.State,
		&user.AvataImage,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
