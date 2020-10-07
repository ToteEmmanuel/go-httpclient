package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func (hc *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	fullHeaders := hc.getRequestHeaders(headers)
	requestBody, err := hc.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Unable to create http Request %s", err)
	}
	return client.Do(request)
}

func (hc *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (hc *httpClient) getRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add common headers to the request.
	for header, value := range hc.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	// Add custom headers to the request.
	for header, value := range customHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result

}
