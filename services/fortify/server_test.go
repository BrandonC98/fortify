package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCredentialRepository struct {
	mock.Mock
}

func (m *MockCredentialRepository) AddCredsRecord(creds *Credentials) {
	slog.Info("Mocking credentials", creds.Name, creds.Passwd)
}

func (m *MockCredentialRepository) retriveAllCreds() []Credentials {
	c := []Credentials{
		{
			ID:     1,
			Name:   "key",
			Passwd: "val",
		},
	}

	return c
}
func TestShowEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var tests = []struct {
		name               string
		endpoint           string
		requestType        string
		expectedStatusCode int
		expectedBody       string
	}{
		{"successfully hit the show endpoint", "/show", "GET", 200, "key: val\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testRepo := new(MockCredentialRepository)
			testRepo.On("retriveAllCreds")

			router := gin.Default()
			router.GET(test.endpoint, showHandler(testRepo))

			req, err := http.NewRequest(test.requestType, test.endpoint, nil)
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.expectedStatusCode, recorder.Code)
			assert.Equal(t, test.expectedBody, recorder.Body.String())
		})
	}
}

func TestSaveEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var tests = []struct {
		name        string
		endpoint    string
		requestType string
		statusCode  int
		body        string
	}{
		{"successfully hit the show endpoint", "/save", "POST", 200, "{ \"website1\": \"pass1\"}"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testRepo := new(MockCredentialRepository)
			testRepo.On("AddCredsRecord")

			router := gin.Default()
			router.POST(test.endpoint, saveHandler(testRepo))

			req, err := http.NewRequest(test.requestType, test.endpoint, strings.NewReader(test.body))
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.statusCode, recorder.Code)
			assert.Equal(t, "successful", recorder.Body.String())
		})
	}
}

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
				Body:       io.NopCloser(strings.NewReader(test.inputPayload)),
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
