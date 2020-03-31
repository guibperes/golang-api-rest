package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/app/libs"
)

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

// Save a new post
func Save(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	database = append(database, post)

	var response = libs.Response{Writer: w, Status: 200, Data: database}
	response.SendJSON()
}

// GetAll saved posts
func GetAll(w http.ResponseWriter, r *http.Request) {
	var response = libs.Response{Writer: w, Status: 200, Data: database}
	response.SendJSON()
}

// GetByID a post
func GetByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		var response = libs.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		var response = libs.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	var response = libs.Response{Writer: w, Status: 200, Data: database[postID]}
	response.SendJSON()
}

// UpdateByID a post
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		var response = libs.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		var response = libs.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	var post = &database[postID]
	json.
		NewDecoder(r.Body).
		Decode(post)

	var response = libs.Response{Writer: w, Status: 200, Data: post}
	response.SendJSON()
}

// DeleteByID a post
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		var response = libs.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= len(database) {
		var response = libs.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	database = append(database[:postID], database[postID+1:]...)

	var response = libs.Response{Writer: w, Status: 200}
	response.SendEmpty()
}
