package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controller.GetAllPost).Methods("GET")
	router.HandleFunc("/posts/creating", controller.CreateNewPost).Methods("POST")
	router.HandleFunc("/posts/:postId/comment", controller.CreateComment).Methods("POST")
	router.HandleFunc("/posts/:postId/reaction", controller.PostReact).Methods("GET")
	router.HandleFunc("/posts/:postId/comment/reaction", controller.CommentReact).Methods("GET")
}
