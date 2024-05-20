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
	router.HandleFunc("/user/authentication", auth.GetNewTokenFromRefreshToken).Methods("GET")
	router.HandleFunc("/auth/code", auth.GetCodeSendMail).Methods("POST")
	router.HandleFunc("/auth", auth.AuthenticateCode).Methods("POST")
	router.HandleFunc("/user/update", controller.UpdateUser).Methods("POST")
}
