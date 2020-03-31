package post

import (
	"net/http"

	"github.com/guibperes/golang-api-rest/app/libs"
	"github.com/guibperes/golang-api-rest/app/libs/server"
)

var (
	service = Service{}
)

// Controller definition to access post functions
type Controller struct{}

// Save a new post
func (c Controller) Save(w http.ResponseWriter, r *http.Request) {
	var post Post
	server.
		RequestBuilder(r).
		ParseJSONBody(&post)

	server.
		ResponseBuilder(w, 200, service.Save(post)).
		SendJSON()
}

// UpdateByID a post
func (c Controller) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var request = server.RequestBuilder(r)
	var postID, err = request.GetPathParameterAndParseInt("id")
	var post Post
	request.ParseJSONBody(&post)

	if err != nil {
		server.
			ResponseBuilder(w, 400, libs.Message{Message: "Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		server.
			ResponseBuilder(w, 404, libs.Message{Message: "Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	server.
		ResponseBuilder(w, 200, service.UpdateByID(postID, post)).
		SendJSON()
}

// DeleteByID a post
func (c Controller) DeleteByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = server.
		RequestBuilder(r).
		GetPathParameterAndParseInt("id")

	if err != nil {
		server.
			ResponseBuilder(w, 400, libs.Message{Message: "Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		server.
			ResponseBuilder(w, 404, libs.Message{Message: "Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	service.DeleteByID(postID)

	server.
		ResponseBuilder(w, 200, nil).
		SendEmpty()
}

// GetAll saved posts
func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	server.
		ResponseBuilder(w, 200, service.GetAll()).
		SendJSON()
}

// GetByID a post
func (c Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	var postID, err = server.
		RequestBuilder(r).
		GetPathParameterAndParseInt("id")

	if err != nil {
		server.
			ResponseBuilder(w, 400, libs.Message{Message: "Cannot convert id to integer"}).
			SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		server.
			ResponseBuilder(w, 404, libs.Message{Message: "Cannot find post with the provided id"}).
			SendJSON()
		return
	}

	server.
		ResponseBuilder(w, 200, service.GetByID(postID)).
		SendJSON()
}
