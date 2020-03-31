package server

import (
	"encoding/json"
	"net/http"
)

// Response handler to empty and JSON response
type Response struct {
	Writer http.ResponseWriter
	Status int
	Data   interface{}
}

// ResponseBuilder build the response object
func ResponseBuilder(writer http.ResponseWriter, status int, data interface{}) Response {
	return Response{Writer: writer, Status: status, Data: data}
}

// SendJSON create a JSON response
func (response *Response) SendJSON() {
	response.
		Writer.
		WriteHeader(response.Status)
	response.
		Writer.
		Header().
		Set("Content-Type", "application/json")
	json.
		NewEncoder(response.Writer).
		Encode(response.Data)
}

// SendEmpty create a empty response
func (response *Response) SendEmpty() {
	response.
		Writer.
		WriteHeader(response.Status)
}
