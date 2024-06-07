package routes

import (
	"github.com/PhanLuc1/Blogify-Project-Backend/src/controller"
	"github.com/gorilla/mux"
)

var RegisterImageRoute = func(router *mux.Router) {
	router.HandleFunc("/image", controller.ImageHandle).Methods("GET")
	//router.HandleFunc("/upload", controller.UploadeHandle).Methods("POST")
}
