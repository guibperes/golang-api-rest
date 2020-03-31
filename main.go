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

// ResponseJSON parameters to create JSON response
type ResponseJSON struct {
	Writer http.ResponseWriter
	Status int
	Data   interface{}
}

// ResponseEmpty parameters to create empty response
type ResponseEmpty struct {
	Writer http.ResponseWriter
	Status int
}

func sendJSON(response ResponseJSON) {
	response.Writer.WriteHeader(response.Status)
	response.Writer.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(response.Writer).
		Encode(response.Data)
}

func sendEmpty(response ResponseEmpty) {
	response.Writer.WriteHeader(response.Status)
}

func save(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	database = append(database, post)

	sendJSON(ResponseJSON{Writer: w, Status: 200, Data: database})
}

func getAll(w http.ResponseWriter, r *http.Request) {
	sendJSON(ResponseJSON{Writer: w, Status: 200, Data: database})
}

func getByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		sendJSON(ResponseJSON{Writer: w, Status: 400, Data: Message{"Cannot convert id to integer"}})
		return
	}

	if postID < 0 || postID >= len(database) {
		sendJSON(ResponseJSON{Writer: w, Status: 404, Data: Message{"Cannot find post with the provided id"}})
		return
	}

	sendJSON(ResponseJSON{Writer: w, Status: 200, Data: database[postID]})
}

func updateByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		sendJSON(ResponseJSON{Writer: w, Status: 400, Data: Message{"Cannot convert id to integer"}})
		return
	}

	if postID < 0 || postID >= len(database) {
		sendJSON(ResponseJSON{Writer: w, Status: 404, Data: Message{"Cannot find post with the provided id"}})
		return
	}

	var updatedPost Post
	json.
		NewDecoder(r.Body).
		Decode(&updatedPost)

	database[postID] = updatedPost

	sendJSON(ResponseJSON{Writer: w, Status: 200, Data: database[postID]})
}

func patchByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		sendJSON(ResponseJSON{Writer: w, Status: 400, Data: Message{"Cannot convert id to integer"}})
		return
	}

	if postID < 0 || postID >= len(database) {
		sendJSON(ResponseJSON{Writer: w, Status: 404, Data: Message{"Cannot find post with the provided id"}})
		return
	}

	var post = &database[postID]
	json.
		NewDecoder(r.Body).
		Decode(post)

	sendJSON(ResponseJSON{Writer: w, Status: 200, Data: post})
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		sendJSON(ResponseJSON{Writer: w, Status: 400, Data: Message{"Cannot convert id to integer"}})
		return
	}

	if postID < 0 || postID >= len(database) {
		sendJSON(ResponseJSON{Writer: w, Status: 404, Data: Message{"Cannot find post with the provided id"}})
		return
	}

	database = append(database[:postID], database[postID+1:]...)

	sendEmpty(ResponseEmpty{Writer: w, Status: 200})
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
