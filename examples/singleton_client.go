package examples

import (
	"net/http"
	"time"

	"github.com/ToteEmmanuel/go-httpclient/gohttp"
	"github.com/ToteEmmanuel/go-httpclient/mime"
)

var (
	httpClient = getHTTPClient()
)

func getHTTPClient() gohttp.HTTPClient {
	headers := make(http.Header)
	headers.Set(mime.HeaderContentType, mime.ContentTypeJSON)
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetRequestTimeout(3 * time.Second).
		SetHeaders(headers).
		Build()
	return client

}
