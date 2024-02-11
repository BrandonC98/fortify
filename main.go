package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	startServer(getConfig())
}

func getConfig() Config {
	var errorMsg error
	var port int
	var passGenURL, dataAccessURL string

	port, errorMsg = strconv.Atoi(os.Getenv("PASSMAN_PORT"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}

	passGenURL = os.Getenv("PASSMAN_PASS_GEN_URL")
	dataAccessURL = os.Getenv("PASSMAN_DATA_ACCESS_URL")

	c := Config{
		Port:          port,
		PassGenURL:    passGenURL,
		DataAccessURL: dataAccessURL,
	}

	return c
}

type Config struct {
	Port          int
	PassGenURL    string
	DataAccessURL string
}
