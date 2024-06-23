package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	Resp *http.Response
	Err  error
}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	log.Print("Input Request url -> " + req.URL.String())
	return c.Resp, c.Err
}

func TestGetGeneratedPassword(t *testing.T) {
	password := "test"
	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(password)),
	}

	client := MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}

	var tests = []struct {
		name             string
		inputURL         string
		inputClient      HTTPClient
		expectedPassword string
	}{
		{"successfully get password", "/testEndpoint", &client, "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getGeneratedPassword(test.inputURL, test.inputClient)
			assert.Equal(t, test.expectedPassword, actual)
		})
	}
}
