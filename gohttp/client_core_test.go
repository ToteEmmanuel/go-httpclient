package gohttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	//Init
	testClient := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	testClient.Headers = commonHeaders

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
	testClient := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	testClient.Headers = commonHeaders

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
