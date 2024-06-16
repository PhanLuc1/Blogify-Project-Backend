package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controller.GetAllPost).Methods("GET")
	router.HandleFunc("/posts/creating", controller.UploadeHandle).Methods("POST")
	router.HandleFunc("/posts/{postid}", controller.GetPostById).Methods("GET")
	router.HandleFunc(`/posts/{postid:\d+}/comment`, controller.CreateComment).Methods("POST")
	router.HandleFunc("/posts/{postid}/reaction", controller.PostReact).Methods("GET")
	router.HandleFunc("/posts/{commentid}/reaction", controller.CommentReact).Methods("GET")
	router.HandleFunc("/posts/{postid}/delete", controller.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/comment/{commentid}/delete", controller.DeleteComment).Methods("DELETE")
	router.HandleFunc("/posts/comment/update", controller.UpdateComment).Methods("POST")
	router.HandleFunc(`/posts/{postid:\d+}/update`, controller.UpdatePost).Methods("POST")
}
