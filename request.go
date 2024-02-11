package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type standardHTTPClient struct{}

func (c *standardHTTPClient) Do(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}

func getGeneratedPassword(endpointURL string, client HTTPClient) string {
	req, err := http.NewRequest(http.MethodGet, endpointURL, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
