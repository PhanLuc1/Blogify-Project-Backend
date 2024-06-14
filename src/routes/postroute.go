package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controller.GetAllPost).Methods("GET")
	router.HandleFunc("/posts/creating", controller.UploadeHandle).Methods("POST")
	router.HandleFunc("/posts/{postid}", controller.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postid}/comment", controller.CreateComment).Methods("POST")
	router.HandleFunc("/posts/{postid}/reaction", controller.PostReact).Methods("GET")
	router.HandleFunc("/posts/{commentid}/reaction", controller.CommentReact).Methods("GET")
	router.HandleFunc("/posts/{postid}/delete", controller.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/{commentid}/delete", controller.DeleteComment).Methods("DELETE")
}
