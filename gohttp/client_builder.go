package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	connectionTimeout  time.Duration
	disableTimeouts    bool
	maxIdleconnections int
	requestTimeout     time.Duration
	userAgent          string
	client             *http.Client
}

//ClientBuilder is exposed to create HTTPClients throught the implemented methods.
type ClientBuilder interface {
	DisableTimeouts(bool) ClientBuilder
	SetConnectionTimeout(time.Duration) ClientBuilder
	SetHeaders(http.Header) ClientBuilder
	SetMaxIdleConnections(int) ClientBuilder
	SetRequestTimeout(time.Duration) ClientBuilder
	SetHTTPClient(c *http.Client) ClientBuilder
	SetUserAgent(uAgent string) ClientBuilder
	Build() HTTPClient
}

func (cb *clientBuilder) Build() HTTPClient {
	client := httpClient{
		builder: cb,
	}
	return &client
}

func (cb *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	cb.disableTimeouts = disableTimeouts
	return cb
}

func (cb *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	cb.connectionTimeout = timeout
	return cb
}

func (cb *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	cb.headers = headers
	return cb
}

func (cb *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	cb.requestTimeout = timeout
	return cb
}
func (cb *clientBuilder) SetMaxIdleConnections(maxIdleConns int) ClientBuilder {
	cb.maxIdleconnections = maxIdleConns
	return cb
}

func (cb *clientBuilder) SetHTTPClient(client *http.Client) ClientBuilder {
	cb.client = client
	return cb
}
func (cb *clientBuilder) SetUserAgent(uAgent string) ClientBuilder {
	cb.userAgent = uAgent
	return cb
}

//NewBuilder returns an instance of a ClientBuilder struct.
func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}
