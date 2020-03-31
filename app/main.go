package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/app/api/post"
	"github.com/guibperes/golang-api-rest/app/libs/log"
)

func main() {
	var router = mux.NewRouter()
	var log = log.Builder()

	var postController = post.ControllerBuilder()

	router.Use(log.GetLogMiddleware)

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

	log.Info("Server started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
