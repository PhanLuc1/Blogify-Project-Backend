package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/user/login", controller.Login).Methods("POST")
}
