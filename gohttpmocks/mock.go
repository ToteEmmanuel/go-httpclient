package gohttpmocks

import (
	"fmt"
	"net/http"

	"github.com/ToteEmmanuel/go-httpclient/core"
)

//Represents a Mock Key.
type Mock struct {
	Method             string
	URL                string
	RequestBody        string
	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return &core.Response{
		Body:       []byte(m.ResponseBody),
		StatusCode: m.ResponseStatusCode,
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
	}, nil
}
