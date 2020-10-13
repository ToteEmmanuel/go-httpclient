package examples

import (
	"fmt"
)

type Endpoints struct {
	AuthorizationUrl string `json:"authorization_url"`
	CurrentUserUrl   string `json:"current_user_url"`
	RespositoryUrl   string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		//Deal with errors as needed.
		return nil, err
	}

	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Body:\n%s\n", response.BodyString())
	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil {
		return nil, err
	}
	fmt.Printf("Authorization URL: %s\n", endpoints.AuthorizationUrl)
	fmt.Printf("Current User URL: %s\n", endpoints.RespositoryUrl)
	fmt.Printf("Repository URL: %s\n", endpoints.AuthorizationUrl)
	return &endpoints, nil
}
