package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/ToteEmmanuel/go-httpclient/gohttpmocks"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting Examples testing suite.")
	gohttpmocks.MockedServer.Start()
	os.Exit(m.Run())
}
func TestGetEndpoints(t *testing.T) {
	//Initialization

	t.Run("Error fetching from URL", func(t *testing.T) {
		gohttpmocks.MockedServer.DeleteMocks()
		gohttpmocks.MockedServer.AddMock(gohttpmocks.Mock{
			Method: http.MethodGet,
			URL:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})
		//Execution
		endpoints, err := GetEndpoints()
		assert.Nil(t, endpoints, "No endpoints were expected.")
		assert.EqualError(t, err, "timeout getting github endpoints", "Not expected error.")

	})
	t.Run("Error unmarshalling response body", func(t *testing.T) {
		gohttpmocks.MockedServer.DeleteMocks()
		gohttpmocks.MockedServer.AddMock(gohttpmocks.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseStatusCode: http.StatusOK,
		})
		//Execution
		endpoints, err := GetEndpoints()
		assert.Nil(t, endpoints, "No endpoints were expected.")
		assert.Contains(t, err.Error(), "cannot unmarshal", "Expected unmarshalling error.")

	})
	t.Run("Successful fetch from URL", func(t *testing.T) {
		gohttpmocks.MockedServer.DeleteMocks()
		gohttpmocks.MockedServer.AddMock(gohttpmocks.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
			ResponseStatusCode: http.StatusOK,
		})
		//Execution
		endpoints, err := GetEndpoints()
		assert.Nil(t, err, "No error was expected.")
		assert.NotNil(t, endpoints, "Response shouldn't be nil.")
		assert.Equal(t, "https://api.github.com/user", endpoints.CurrentUserUrl, "Current user URL is not the expected one.")

	})

}
