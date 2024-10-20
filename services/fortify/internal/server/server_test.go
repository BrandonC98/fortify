package server

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BrandonC98/fortify/services/fortify/internal/model"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	decryptedTestvalue = "test2"
	encryptedTestvalue = `-----BEGIN PGP MESSAGE-----
Version: GopenPGP 2.7.4
Comment: https://gopenpgp.org

wy4ECQMIocqe6gm6h9XgqbdLH9Uv3S4LHAiT1TIhIO5NGzXAUdzD18u+J6zQ7ZUv
0jYB/9nn70fZvkW8iMCghmg8vdAK3dqPfomC2D1buSUed2DC1CO3UcqCgDi5E7c5
Yk3a3bdrvyc=
=Fki9
-----END PGP MESSAGE-----
`
)

type mockSecretRepository interface {
	AddRecord(creds *model.Secret)
	RetriveAllRecords() []model.Secret
	On(methodName string, arguments ...interface{}) *mock.Call
}

type MockEncryptedSecretsRepository struct {
	mock.Mock
}

func (m *MockEncryptedSecretsRepository) AddRecord(creds *model.Secret) {
	slog.Info("Mocking secrets", creds.Name, creds.Value)
}

func (m *MockEncryptedSecretsRepository) RetriveAllRecords() []model.Secret {
	c := []model.Secret{
		{
			ID:    1,
			Name:  "key",
			Value: encryptedTestvalue,
		},
	}
	return c
}

type MockDecryptedSecretsRepository struct {
	mock.Mock
}

func (m *MockDecryptedSecretsRepository) AddRecord(creds *model.Secret) {
	slog.Info("Mocking secrets", creds.Name, creds.Value)
}

func (m *MockDecryptedSecretsRepository) RetriveAllRecords() []model.Secret {
	c := []model.Secret{
		{
			ID:    1,
			Name:  "key",
			Value: decryptedTestvalue,
		},
	}
	return c
}
func TestShowEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config := model.Config{
		Port:         80,
		GeneratorURL: "generator",
		DBUser:       "user",
		DBHost:       "host",
		DBPassword:   "password",
		Key:          "TEST_KEY",
	}
	var tests = []struct {
		name                string
		endpoint            string
		requestType         string
		useEncryptedSecrets bool
		expectedStatusCode  int
		expectedBody        string
	}{
		{"successfully hit the show endpoint with encrypted values stored", "/show", "GET", true, 200, fmt.Sprintf("key: %s\n", decryptedTestvalue)},
		{"successfully hit the show endpoint with unencrypted values stored", "/show", "GET", false, 200, fmt.Sprintf("key: %s\n", decryptedTestvalue)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testRepo mockSecretRepository
			if test.useEncryptedSecrets == true {
				testRepo = new(MockEncryptedSecretsRepository)
			} else {
				testRepo = new(MockDecryptedSecretsRepository)
			}

			testRepo.On("retriveAllCreds")

			router := gin.Default()
			router.GET(test.endpoint, showHandler(testRepo, config))

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
	config := model.Config{
		Port:         80,
		GeneratorURL: "generator",
		DBUser:       "user",
		DBHost:       "host",
		DBPassword:   "password",
		Key:          "TEST_KEY",
	}
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
			testRepo := new(MockEncryptedSecretsRepository)
			testRepo.On("AddCredsRecord")

			router := gin.Default()
			router.POST(test.endpoint, saveHandler(testRepo, config))

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
	config := model.Config{
		Port:         80,
		GeneratorURL: "generator",
		DBUser:       "user",
		DBHost:       "host",
		DBPassword:   "password",
		Key:          "TEST_KEY",
	}

	encryptedPassword, err := helper.EncryptMessageWithPassword([]byte(config.Key), "passman_password")
	assert.NoError(t, err)

	var tests = []struct {
		name               string
		mockStatusCode     int
		inputEndpoint      string
		inputPayload       string
		expectedStatusCode int
		expectedPayload    string
	}{
		{"successfully generate password", http.StatusOK, "/generate", encryptedPassword, 200, `{"message":"passman_password"}`},
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
			router.GET(test.inputEndpoint, generateHandler("/testEndpoint", &mockClient, config))

			req, err := http.NewRequest(http.MethodGet, test.inputEndpoint, nil)
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, test.expectedStatusCode, recorder.Code)
			assert.Equal(t, test.expectedPayload, recorder.Body.String())
		})
	}
}
