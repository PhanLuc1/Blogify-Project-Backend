package main

import (
	"log"
	"net/http"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/middleware"
	"github.com/PhanLuc1/Blogify-Project-Backend/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.CORSMiddleware)
	http.Handle("/", router)
	routes.RegisterPostRoutes(router)
	routes.RegisterUserRoutes(router)
	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	
}
