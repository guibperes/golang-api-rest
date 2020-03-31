package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/app/api/post"
)

func main() {
	var router = mux.NewRouter()

	router.
		HandleFunc("/posts", post.Save).
		Methods("POST")
	router.
		HandleFunc("/posts/{id}", post.UpdateByID).
		Methods("PATCH")
	router.
		HandleFunc("/posts/{id}", post.DeleteByID).
		Methods("DELETE")
	router.
		HandleFunc("/posts", post.GetAll).
		Methods("GET")
	router.
		HandleFunc("/posts/{id}", post.GetByID).
		Methods("GET")

	http.ListenAndServe(":5000", router)
}
