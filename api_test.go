package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var tests = []struct {
		name        string
		endpoint    string
		requestType string
		statusCode  int
		body        string
	}{
		{"successfully hit the ping endpoint", "/ping", "GET", 200, `{"message":"pong"}`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := gin.Default()
			router.GET(test.endpoint, pingHandler)

			req, err := http.NewRequest(test.requestType, test.endpoint, nil)
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.statusCode, recorder.Code)
			assert.Equal(t, test.body, recorder.Body.String())
		})
	}
}

func TestGenerateEndpoint(t *testing.T) {
	config := Config{
		Mode:            "test",
		Port:            8080,
		StringMinLength: 7,
		StringMaxLength: 25,
	}
	gin.SetMode(config.Mode)
	var tests = []struct {
		name         string
		endpoint     string
		requestType  string
		statusCode   int
		bodyContains []string
	}{
		{"successfully hit the generate endpoint", "/generate", "GET", 200, []string{"BEGIN PGP MESSAGE", "END PGP MESSAGE"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			router := gin.Default()
			router.GET(test.endpoint, generateHandler(config))

			req, err := http.NewRequest(test.requestType, test.endpoint, nil)
			assert.Nil(t, err)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.statusCode, recorder.Code)
			for _, expectedString := range test.bodyContains {
				assert.Contains(t, recorder.Body.String(), expectedString)
			}
		})
	}
}
