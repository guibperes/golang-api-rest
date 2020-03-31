package response

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

// Builder function builder to a Response object
func Builder() Response {
	return Response{Status: 200}
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
