package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseBodyBytes(t *testing.T) {
	bodyText := "Random body text."
	response := Response{
		Body: []byte(bodyText),
	}
	assert.Equal(t, []byte(bodyText), response.Body, "Byte array in body should match.")
}

func TestResponseBodyString(t *testing.T) {
	bodyText := "Random body text."
	response := Response{
		Body: []byte(bodyText),
	}
	assert.Equal(t, bodyText, response.BodyString(), "Body string should match.")
}
func TestResponseBodyUnmarshal(t *testing.T) {
	bodyValue := "Random body text."
	bodyText := fmt.Sprintf(`{"value":"%s"}`, bodyValue)
	response := Response{
		Body: []byte(bodyText),
	}
	type testObject struct {
		Value string `json:"value"`
	}
	var testStruct testObject
	testStruct.Value = bodyValue
	response.UnmarshalJSON(&testStruct)
	assert.Equal(t, bodyValue, testStruct.Value, "unmarshalled value should match.")
}
