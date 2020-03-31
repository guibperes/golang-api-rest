package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guibperes/golang-api-rest/app/libs"
)

var (
	service = Service{}
)

// Controller definition to access post functions
type Controller struct{}

// Save a new post
func (c *Controller) Save(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	var response = libs.Response{Writer: w, Status: 200, Data: service.Save(post)}
	response.SendJSON()
}

// UpdateByID a post
func (c *Controller) UpdateByID(w http.ResponseWriter, r *http.Request) {
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

	var post Post
	json.
		NewDecoder(r.Body).
		Decode(&post)

	var response = libs.Response{Writer: w, Status: 200, Data: service.UpdateByID(postID, post)}
	response.SendJSON()
}

// DeleteByID a post
func (c *Controller) DeleteByID(w http.ResponseWriter, r *http.Request) {
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

	service.DeleteByID(postID)

	var response = libs.Response{Writer: w, Status: 200}
	response.SendEmpty()
}

// GetAll saved posts
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	var response = libs.Response{Writer: w, Status: 200, Data: service.GetAll()}
	response.SendJSON()
}

// GetByID a post
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		var response = libs.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		var response = libs.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	var response = libs.Response{Writer: w, Status: 200, Data: service.GetByID(postID)}
	response.SendJSON()
}
