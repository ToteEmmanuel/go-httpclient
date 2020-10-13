package gohttp

import (
	"net/http"

	"github.com/ToteEmmanuel/go-httpclient/mime"
)

func (hc *httpClient) getRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add common headers to the request.
	for header, value := range hc.builder.headers {
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
	//Set default user agent if none other has been provided
	if hc.builder.userAgent != "" {
		if result.Get(mime.HeaderUserAgent) != "" {
			result.Set(mime.HeaderUserAgent, hc.builder.userAgent)
		}
	}
	return result
}

func (hc *httpClient) reduceHeaders(headers ...http.Header) *http.Header {
	var reducedHeaders = make(http.Header)
	if len(headers) > 0 {
		for _, headerEntry := range headers {
			for hKey, hValues := range headerEntry {
				for _, singleValue := range hValues {
					reducedHeaders.Add(hKey, singleValue)
				}
			}
		}
	}
	return &reducedHeaders
}
