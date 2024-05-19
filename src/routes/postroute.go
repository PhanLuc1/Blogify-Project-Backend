package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controller.GetAllPost).Methods("GET")
	router.HandleFunc("/posts/creating", controller.CreateNewPost).Methods("POST")
	//router.HandleFunc("/posts/{postid}", controller.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postid}/comment", controller.CreateComment).Methods("POST")
	router.HandleFunc("/posts/{postid}/reaction", controller.PostReact).Methods("GET")
	router.HandleFunc("/posts/{postId}/comment/reaction", controller.CommentReact).Methods("GET")
}
