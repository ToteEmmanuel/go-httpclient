package core

import "net/http"

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}
