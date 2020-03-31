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
func (c *Controller) Save(w http.ResponseWriter, r *http.Request) {
	var request = server.Request{Data: r}

	var post Post
	request.ParseJSONBody(&post)

	var response = server.Response{Writer: w, Status: 200, Data: service.Save(post)}
	response.SendJSON()
}

// UpdateByID a post
func (c *Controller) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var request = server.Request{Data: r}
	var postID, err = request.GetPathParameterAndParseInt("id")

	if err != nil {
		var response = server.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		var response = server.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	var post Post
	request.ParseJSONBody(&post)

	var response = server.Response{Writer: w, Status: 200, Data: service.UpdateByID(postID, post)}
	response.SendJSON()
}

// DeleteByID a post
func (c *Controller) DeleteByID(w http.ResponseWriter, r *http.Request) {
	var request = server.Request{Data: r}
	var postID, err = request.GetPathParameterAndParseInt("id")

	if err != nil {
		var response = server.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		var response = server.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	service.DeleteByID(postID)

	var response = server.Response{Writer: w, Status: 200}
	response.SendEmpty()
}

// GetAll saved posts
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	var response = server.Response{Writer: w, Status: 200, Data: service.GetAll()}
	response.SendJSON()
}

// GetByID a post
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	var request = server.Request{Data: r}
	var postID, err = request.GetPathParameterAndParseInt("id")

	if err != nil {
		var response = server.Response{Writer: w, Status: 400, Data: libs.Message{Message: "Cannot convert id to integer"}}
		response.SendJSON()
		return
	}

	if postID < 0 || postID >= service.GetLength() {
		var response = server.Response{Writer: w, Status: 404, Data: libs.Message{Message: "Cannot find post with the provided id"}}
		response.SendJSON()
		return
	}

	var response = server.Response{Writer: w, Status: 200, Data: service.GetByID(postID)}
	response.SendJSON()
}
