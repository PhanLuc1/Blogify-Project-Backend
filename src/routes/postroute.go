package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/post", controller.GetAllPost).Methods("GET")
	router.HandleFunc("/post/create", controller.CreateNewPost).Methods("POST")
	router.HandleFunc("/comment", controller.CreateComment).Methods("POST")
	router.HandleFunc("/post/reaction", controller.PostReact).Methods("GET")
	router.HandleFunc("/post/comment/react", controller.CommentReact).Methods("GET")
}
