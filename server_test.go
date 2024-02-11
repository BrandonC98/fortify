package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/helper"
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

func TestGeneratePasswordEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	encryptedPassword, err := helper.EncryptMessageWithPassword([]byte("PASSMAN_PASS_GEN_KEY"), "passman_password")
	assert.NoError(t, err)

	var tests = []struct {
		name               string
		mockStatusCode     int
		inputEndpoint      string
		inputPayload       string
		expectedStatusCode int
		expectedPayload    string
	}{
		{"successfully generate password", http.StatusOK, "/generatePassword", encryptedPassword, 200, `{"message":"passman_password"}`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockResp := &http.Response{
				StatusCode: test.mockStatusCode,
				Body:       ioutil.NopCloser(strings.NewReader(test.inputPayload)),
			}

			mockClient := MockHTTPClient{
				Resp: mockResp,
				Err:  nil,
			}

			router := gin.Default()
			router.GET(test.inputEndpoint, generatePasswordHandler("/testEndpoint", &mockClient))

			req, err := http.NewRequest(http.MethodGet, test.inputEndpoint, nil)
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.expectedStatusCode, recorder.Code)
			assert.Equal(t, test.expectedPayload, recorder.Body.String())
		})
	}
}
