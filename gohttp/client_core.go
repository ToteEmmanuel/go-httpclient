package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ToteEmmanuel/go-httpclient/core"
	"github.com/ToteEmmanuel/go-httpclient/gohttpmocks"
	"github.com/ToteEmmanuel/go-httpclient/mime"
)

const (
	defaultConnectionTimeout  = 1 * time.Second
	defaultMaxIdleConnections = 5
	defaultRequestTimeout     = 5 * time.Second
)

func (hc *httpClient) do(method string, url string, body interface{}, headers []http.Header) (*core.Response, error) {
	fullHeaders := hc.getRequestHeaders(*hc.reduceHeaders(headers...))
	requestBody, err := hc.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Unable to create http Request %s", err)
	}
	response, err := hc.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	finalResponse := core.Response{
		Body:       responseBody,
		Headers:    response.Header,
		Status:     response.Status,
		StatusCode: response.StatusCode,
	}
	return &finalResponse, nil

}

func (hc *httpClient) getHTTPClient() core.HTTPClient {
	if gohttpmocks.MockedServer.IsEnabled() {
		return gohttpmocks.MockedServer.GetMockedClient()
	}
	hc.clientCreation.Do(func() {
		if hc.builder.client != nil {
			hc.client = hc.builder.client
			return
		}
		hc.client = &http.Client{
			Timeout: hc.getConnectionTimeout() + hc.getRequestTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   hc.getMaxIdleConnections(),
				ResponseHeaderTimeout: hc.getRequestTimeout(),
				DialContext: (&net.Dialer{
					Timeout: hc.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return hc.client
}

func (hc *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case mime.ContentTypeJSON:
		return json.Marshal(body)
	case mime.ContentTypeXML:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (hc *httpClient) getMaxIdleConnections() int {
	if hc.builder.maxIdleconnections > 0 {
		return hc.builder.maxIdleconnections
	}
	return defaultMaxIdleConnections
}

func (hc *httpClient) getRequestTimeout() time.Duration {
	if hc.builder.requestTimeout > 0 {
		return hc.builder.requestTimeout
	}
	return defaultRequestTimeout
}
func (hc *httpClient) getConnectionTimeout() time.Duration {
	if hc.builder.connectionTimeout > 0 {
		return hc.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}
