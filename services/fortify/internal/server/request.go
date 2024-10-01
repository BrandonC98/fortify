package server

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type standardHTTPClient struct{}

func (c *standardHTTPClient) Do(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	slog.Info("Sending request", "url", req.URL, "method", req.Method)
	return client.Do(req)
}

func generate(endpointURL string, client HTTPClient) string {
	req, err := http.NewRequest(http.MethodGet, endpointURL, bytes.NewBuffer([]byte{}))
	if err != nil {
		slog.Error(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Response received", "status", resp.Status, "body", string(body))

	return string(body)
}
