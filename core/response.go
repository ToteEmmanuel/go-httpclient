package core

import (
	"encoding/json"
	"net/http"
)

/*Response provides access to the http.response and has the most used attributes exposed as well as some
methods that help in reading the information. */
type Response struct {
	Headers    http.Header
	Body       []byte
	Status     string
	StatusCode int
}

//BodyBytes returns the byte slice that represents the body of the response
func (r *Response) BodyBytes() []byte {
	return r.Body
}

//BodyString returns the string representation of the response's body
func (r *Response) BodyString() string {
	return string(r.Body)
}

//UnmarshalJSON tries to unmarshal the body into an struct of the target's type
func (r *Response) UnmarshalJSON(target interface{}) error {
	return json.Unmarshal(r.BodyBytes(), target)
}
