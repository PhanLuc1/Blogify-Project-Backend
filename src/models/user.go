package models

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
