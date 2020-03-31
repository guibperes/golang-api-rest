package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/src/response"
)

// Message Data Transfer Object(DTO)
type Message struct {
	Message string `json:"message"`
}

// Post entity model
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  User   `json:"author"`
}

// User entity model
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var database = []Post{}

func save(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	database = append(database, post)

	response.Builder().
		SetWriter(w).
		SetBody(database).
		SendJSON()
}

func getAll(w http.ResponseWriter, r *http.Request) {
	response.Builder().
		SetWriter(w).
		SetBody(database).
		SendJSON()
}

func getByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		response.Builder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		response.Builder().
			SetWriter(w).
			SetStatus(404).
			SetBody(Message{"Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	response.Builder().
		SetWriter(w).
		SetBody(database[postID]).
		SendJSON()
}

func patchByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		response.Builder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		response.Builder().
			SetWriter(w).
			SetStatus(404).
			SetBody(Message{"Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	var post = &database[postID]
	json.
		NewDecoder(r.Body).
		Decode(post)

	response.Builder().
		SetWriter(w).
		SetBody(post).
		SendJSON()
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		response.Builder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		response.Builder().
			SetWriter(w).
			SetStatus(404).
			SetBody(Message{"Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	database = append(database[:postID], database[postID+1:]...)

	response.Builder().
		SetWriter(w).
		SendEmpty()
}

func main() {
	var router = mux.NewRouter()

	router.
		HandleFunc("/posts", save).
		Methods("POST")
	router.
		HandleFunc("/posts/{id}", patchByID).
		Methods("PATCH")
	router.
		HandleFunc("/posts/{id}", deleteByID).
		Methods("DELETE")
	router.
		HandleFunc("/posts", getAll).
		Methods("GET")
	router.
		HandleFunc("/posts/{id}", getByID).
		Methods("GET")

	http.ListenAndServe(":5000", router)
}
