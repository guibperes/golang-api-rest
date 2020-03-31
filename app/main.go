package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/app/api/post"
)

var (
	postController = post.Controller{}
)

func main() {
	var router = mux.NewRouter()

	router.
		HandleFunc("/posts", postController.Save).
		Methods("POST")
	router.
		HandleFunc("/posts/{id}", postController.UpdateByID).
		Methods("PATCH")
	router.
		HandleFunc("/posts/{id}", postController.DeleteByID).
		Methods("DELETE")
	router.
		HandleFunc("/posts", postController.GetAll).
		Methods("GET")
	router.
		HandleFunc("/posts/{id}", postController.GetByID).
		Methods("GET")

	http.ListenAndServe(":5000", router)
}
