package gohttp

import (
	"net/http"
)

type httpClient struct {
	Headers http.Header
}

//HTTPClient to be used as main entrypoint of project.
type HTTPClient interface {
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)
	SetHeaders(http.Header)
}

func (hc *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, nil)
}
func (hc *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodPost, url, headers, body)
}
func (hc *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodPut, url, headers, body)
}
func (hc *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodPatch, url, headers, body)
}
func (hc *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return hc.do(http.MethodDelete, url, headers, nil)
}
func (hc *httpClient) SetHeaders(headers http.Header) {
	hc.Headers = headers
}

//New returns an instance of an HTTPClient Object.
func New() HTTPClient {
	client := &httpClient{}
	return client
}
