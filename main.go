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

// Response handler to empty and JSON response
type Response struct {
	Writer http.ResponseWriter
	Status int
	Data   interface{}
}

// SetWriter set the writer parameter
func (response Response) SetWriter(writer http.ResponseWriter) Response {
	response.Writer = writer
	return response
}

// SetStatus set the response status parameter
func (response Response) SetStatus(status int) Response {
	response.Status = status
	return response
}

// SetBody set the response body parameter
func (response Response) SetBody(data interface{}) Response {
	response.Data = data
	return response
}

// ResponseBuilder function builder to a Response object
func ResponseBuilder() Response {
	return Response{Status: 200}
}

// SendJSON create a JSON response
func (response Response) SendJSON() {
	response.Writer.WriteHeader(response.Status)
	response.Writer.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(response.Writer).
		Encode(response.Data)
}

// SendEmpty create a empty response
func (response Response) SendEmpty() {
	response.Writer.WriteHeader(response.Status)
}

func save(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	database = append(database, post)

	ResponseBuilder().
		SetWriter(w).
		SetBody(database).
		SendJSON()
}

func getAll(w http.ResponseWriter, r *http.Request) {
	ResponseBuilder().
		SetWriter(w).
		SetBody(database).
		SendJSON()
}

func getByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		ResponseBuilder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		ResponseBuilder().
			SetWriter(w).
			SetStatus(404).
			SetBody(Message{"Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	ResponseBuilder().
		SetWriter(w).
		SetBody(database[postID]).
		SendJSON()
}

func patchByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		ResponseBuilder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		ResponseBuilder().
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

	ResponseBuilder().
		SetWriter(w).
		SetBody(post).
		SendJSON()
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		ResponseBuilder().
			SetWriter(w).
			SetStatus(400).
			SetBody(Message{"Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		ResponseBuilder().
			SetWriter(w).
			SetStatus(404).
			SetBody(Message{"Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	database = append(database[:postID], database[postID+1:]...)

	ResponseBuilder().
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
