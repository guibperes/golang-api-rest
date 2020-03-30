package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(database)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(database)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot convert id to integer"})
		return
	}

	if postID < 0 || postID >= len(database) {
		w.WriteHeader(404)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot find post with the provided id"})
		return
	}

	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(database[postID])
}

func updateByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot convert id to integer"})
		return
	}

	if postID < 0 || postID >= len(database) {
		w.WriteHeader(404)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot find post with the provided id"})
		return
	}

	var updatedPost Post
	json.
		NewDecoder(r.Body).
		Decode(&updatedPost)

	database[postID] = updatedPost

	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(database[postID])
}

func patchByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot convert id to integer"})
		return
	}

	if postID < 0 || postID >= len(database) {
		w.WriteHeader(404)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot find post with the provided id"})
		return
	}

	var post = &database[postID]
	json.
		NewDecoder(r.Body).
		Decode(post)

	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(post)
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot convert id to integer"})
		return
	}

	if postID < 0 || postID >= len(database) {
		w.WriteHeader(404)
		json.
			NewEncoder(w).
			Encode(Message{"Cannot find post with the provided id"})
		return
	}

	database = append(database[:postID], database[postID+1:]...)

	w.WriteHeader(200)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(w).
		Encode(Message{"Route is not implemented"})
}

func main() {
	var router = mux.NewRouter()

	router.
		HandleFunc("/posts", save).
		Methods("POST")
	router.
		HandleFunc("/posts/{id}", updateByID).
		Methods("PUT")
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
