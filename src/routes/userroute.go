package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {

	router.HandleFunc("/user", controller.GetUserInfo).Methods("GET")
	router.HandleFunc("/user/registration", controller.Signup).Methods("POST")
	router.HandleFunc("/user/sign-in", controller.Login).Methods("POST")
	router.HandleFunc("/auth/code", auth.GetCodeSendMail).Methods("POST")
	router.HandleFunc("/auth", auth.AuthenticateCode).Methods("POST")
	router.HandleFunc(`/user/{userid:\d+}/follower`, controller.FollowUser).Methods("GET")
	router.HandleFunc(`/user/{userid:\d+}`, controller.GetUserById).Methods("GET")
	router.HandleFunc(`/user/{userid:\d+}/posts`, controller.GetUserPosts).Methods("GET")
	router.HandleFunc("/user/update", controller.UpdateProfile).Methods("POST")
	router.HandleFunc("/user/update/avatar", controller.UploadAvatarImage).Methods("POST")
	router.HandleFunc("/user/posts", controller.GetCurrentUserPosts).Methods("GET")
	router.HandleFunc("/user/update/state", controller.SetUpStateAccount).Methods("GET")
	router.HandleFunc("/users", controller.GetOtherUsers).Methods("GET")
	router.HandleFunc("/user/auth/forgetpassword", auth.GetResetPasswordCode).Methods("POST")
}
