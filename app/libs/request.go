package libs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Request handler to parse JSON body and parse URL parameters
type Request struct {
	Data *http.Request
}

// ParseJSONBody return object of request body
func (request *Request) ParseJSONBody(data interface{}) {
	json.
		NewDecoder(request.Data.Body).
		Decode(&data)
}

// GetPathParameter return specified parameter of URL path
func (request *Request) GetPathParameter(value string) string {
	return mux.Vars(request.Data)[value]
}

// GetPathParameterAndParseInt return specified parameter of URL path and parse to int
func (request *Request) GetPathParameterAndParseInt(value string) (int, error) {
	return strconv.Atoi(request.GetPathParameter(value))
}
