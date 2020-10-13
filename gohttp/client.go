package gohttp

import (
	"net/http"
	"sync"

	"github.com/ToteEmmanuel/go-httpclient/core"
)

type httpClient struct {
	builder        *clientBuilder
	client         *http.Client
	clientCreation sync.Once
}

//HTTPClient to be used as main entrypoint of project.
type HTTPClient interface {
	Get(string, ...http.Header) (*core.Response, error)
	Post(string, interface{}, ...http.Header) (*core.Response, error)
	Put(string, interface{}, ...http.Header) (*core.Response, error)
	Patch(string, interface{}, ...http.Header) (*core.Response, error)
	Delete(string, ...http.Header) (*core.Response, error)
	Options(string, ...http.Header) (*core.Response, error)
}

func (hc *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodGet, url, nil, headers)
}
func (hc *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodPost, url, body, headers)
}
func (hc *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodPut, url, body, headers)
}
func (hc *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodPatch, url, body, headers)
}
func (hc *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodDelete, url, nil, headers)
}
func (hc *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return hc.do(http.MethodOptions, url, nil, headers)
}
