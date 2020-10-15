package gohttp

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ToteEmmanuel/go-httpclient/mime"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	//Init
	testClient := httpClient{
		builder: &clientBuilder{},
	}
	commonHeaders := make(http.Header)
	commonHeaders.Set(mime.HeaderContentType, mime.ContentTypeJSON)
	commonHeaders.Set(mime.HeaderUserAgent, "cool-http-client")
	testClient.builder.headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	//Run
	finalHeaders := testClient.getRequestHeaders(requestHeaders)
	//Validate
	if len(finalHeaders) != 3 {
		t.Error("3 Headers expected by test.")
	}
}

func TestGetRequestHeadersTestify(t *testing.T) {
	//Init
	testClient := httpClient{
		builder: &clientBuilder{},
	}
	commonHeaders := make(http.Header)
	commonHeaders.Set(mime.HeaderContentType, mime.ContentTypeJSON)
	commonHeaders.Set(mime.HeaderUserAgent, "cool-http-client")
	testClient.builder.headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	//Run
	finalHeaders := testClient.getRequestHeaders(requestHeaders)
	//Validate
	assert.Equal(t, 3, len(finalHeaders), "3 Headers expected by test.")
	assert.Equal(t, "application/json", finalHeaders.Get("Content-Type"),
		"Content-Type expected to be appication/json.")
	assert.Equal(t, "cool-http-client", finalHeaders.Get("User-Agent"),
		"User-Agent expected to be cool-http-client")
	assert.Equal(t, "ABC-123", finalHeaders.Get("X-Request-Id"),
		"X-Requested-Id expected to be ABC-123")
}

func TestGetRequestBody(t *testing.T) {
	//Initialization
	client := httpClient{}
	t.Run("Nil Body", func(t *testing.T) {
		//Run
		body, err := client.getRequestBody("", nil)
		//Validation
		assert.Nil(t, body, "Nil body is expected from nil body parameter")
		assert.Nil(t, err, "No error should be returned when passing a nil body")
	})
	t.Run("JSON Body", func(t *testing.T) {
		//Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		assert.Equal(t, `["one","two"]`, string(body), "Incorrect parsed body.")
		assert.Nil(t, err, "No error should be returned when passing a correct json body")
	})
	t.Run("XML Body", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)
		assert.Equal(t, `<string>one</string><string>two</string>`, string(body), "Incorrect parsed body.")
		assert.Nil(t, err, "No error should be returned when passing a correct body")
	})
	t.Run("Default JSON Body", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)
		assert.Equal(t, `["one","two"]`, string(body), "Incorrect parsed body.")
		assert.Nil(t, err, "No error should be returned when passing a correct body")
	})
}

func TestBuilderValuesSet(t *testing.T) {
	cBuilder := NewBuilder()
	client := cBuilder.SetMaxIdleConnections(10).
		SetConnectionTimeout(3 * time.Minute).
		SetRequestTimeout(5 * time.Minute).
		Build()

	t.Run("Max idle connections", func(t *testing.T) {
		assert.Equal(t, 10, client.(*httpClient).getMaxIdleConnections(), "Max idle connection correctly set.")
	})
	t.Run("Connection timeout", func(t *testing.T) {
		assert.Equal(t, 3*time.Minute, client.(*httpClient).getConnectionTimeout(), "Connection timeout correctly set.")
	})
	t.Run("Request timeout", func(t *testing.T) {
		assert.Equal(t, 5*time.Minute, client.(*httpClient).getRequestTimeout(), "Request timeout correctly set.")
	})
}

func TestDefaultValuesSet(t *testing.T) {
	client := NewBuilder().Build()
	t.Run("Max idle connections", func(t *testing.T) {
		assert.Equal(t, defaultMaxIdleConnections, client.(*httpClient).getMaxIdleConnections(), "Default max idle connection correctly set.")
	})
	t.Run("Connection timeout", func(t *testing.T) {
		assert.Equal(t, defaultConnectionTimeout, client.(*httpClient).getConnectionTimeout(), "Default connection timeout correctly set.")
	})
	t.Run("Request timeout", func(t *testing.T) {
		assert.Equal(t, defaultRequestTimeout, client.(*httpClient).getRequestTimeout(), "Default request timeout correctly set.")
	})
}

func TestGetHTTPClient(t *testing.T) {
	cBuilder := NewBuilder()
	t.Run("Empty builder returns ok", func(t *testing.T) {
		client := cBuilder.Build()
		assert.NotNil(t, client.(*httpClient).getHTTPClient(), "Default max idle connection correctly set.")
		assert.Equal(t, "*http.Client", fmt.Sprintf("%T", client.(*httpClient).getHTTPClient()), `Type should be "core.HTTPClient"`)
	})
	t.Run("Builder returns ok", func(t *testing.T) {
		client := cBuilder.SetConnectionTimeout(5 * time.Second).SetRequestTimeout(5 * time.Second).Build()
		assert.Equal(t, 10*time.Second, client.(*httpClient).getHTTPClient().(*http.Client).Timeout, `Added timeout matches.`)
	})
	t.Run("Builder returns set client ignoring other values", func(t *testing.T) {
		client := cBuilder.SetConnectionTimeout(5 * time.Second).SetRequestTimeout(5 * time.Second).
			SetHTTPClient(&http.Client{Timeout: time.Second}).Build()
		assert.NotEqual(t, 10*time.Second, client.(*httpClient).getHTTPClient().(*http.Client).Timeout, `Timeout should be the one set.`)
	})
}
