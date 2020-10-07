package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"ithub.com/ToteEmmanuel/go-httpclient/gohttp"
)

var (
	githubHTTPClient = getGithubClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	getUrls()
}

func getUrls() {
	response, err := githubHTTPClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHTTPClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func getGithubClient() gohttp.HTTPClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("XYZ", "RandomValue")
	client.SetHeaders(commonHeaders)
	return client
}
