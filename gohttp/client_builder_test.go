package gohttp

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSetConnectionTimeout(t *testing.T) {
	cBuilder := &clientBuilder{}
	cBuilder.SetConnectionTimeout(1 * time.Minute)
	assert.Equal(t, 1*time.Minute, cBuilder.connectionTimeout, "Connection timeout should be set to 1 Min.")
}

func TestSetRequestTimeout(t *testing.T) {
	cBuilder := &clientBuilder{}
	cBuilder.SetRequestTimeout(1 * time.Minute)
	assert.Equal(t, 1*time.Minute, cBuilder.requestTimeout, "Request timeout should be set to 1 Min.")
}

func TestDisableTimeOut(t *testing.T) {
	cBuilder := &clientBuilder{}
	cBuilder.DisableTimeouts(false)
	assert.Equal(t, false, cBuilder.disableTimeouts, "Connection timeout should be set to false.")
	cBuilder.DisableTimeouts(true)
	assert.Equal(t, true, cBuilder.disableTimeouts, "Connection timeout should be set to true.")
}
func TestSetHeaders(t *testing.T) {
	header := http.Header{}
	header.Set("Prop", "Value")
	cBuilder := &clientBuilder{}
	cBuilder.SetHeaders(header)
	assert.Equal(t, header, cBuilder.headers, "Headers should be set.")
	assert.Equal(t, "Value", cBuilder.headers.Get("Prop"), "Value must exists in set header")
}

func TestSetMaxIdleConnections(t *testing.T) {
	cBuilder := &clientBuilder{}
	cBuilder.SetMaxIdleConnections(10)
	assert.Equal(t, 10, cBuilder.maxIdleconnections, "Max idle connections should be set to 10.")
}

func TestSetUserAgent(t *testing.T) {
	cBuilder := &clientBuilder{}
	cBuilder.SetUserAgent("User-Agent-Value")
	assert.Equal(t, "User-Agent-Value", cBuilder.userAgent, `User Agent should be "User-Agent-Value".`)
}

func TestSetHTTPClient(t *testing.T) {
	client := http.Client{
		Timeout: 1 * time.Minute,
	}
	cBuilder := &clientBuilder{}
	cBuilder.SetHTTPClient(&client)
	assert.True(t, cmp.Equal(client, *cBuilder.client), "Client must be correctly set in builder.")
}

func TestBuild(t *testing.T) {
	cBuilder := NewBuilder()
	cBuilder.SetConnectionTimeout(1 * time.Minute).
		SetMaxIdleConnections(2).
		SetRequestTimeout(3 * time.Minute)
	client := (cBuilder.Build()).(*httpClient)
	firstBuilder := cBuilder.(*clientBuilder)
	assert.Equal(t, firstBuilder.connectionTimeout, client.builder.connectionTimeout, "Builder values should match")
	assert.Equal(t, firstBuilder.maxIdleconnections, client.builder.maxIdleconnections, "Builder values should match")
	assert.Equal(t, firstBuilder.requestTimeout, client.builder.requestTimeout, "Builder values should match")
}
