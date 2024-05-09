package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/auth"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controller.GetUserInfo).Methods("GET")
	router.HandleFunc("/user/registration", controller.Signup).Methods("POST")
	router.HandleFunc("/user/logging", controller.Login).Methods("POST")
	router.HandleFunc("/user/authentication", auth.GetNewTokenFromRefreshToken).Methods("GET")
	router.HandleFunc("/user/authentication/code", auth.GetCodeSendMail).Methods("POST")
	router.HandleFunc("/user/authentication/codeauthentication", auth.AuthenticateCode).Methods("POST")
}
