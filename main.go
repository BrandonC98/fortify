package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	config := getConfig()

	startServer(config)
}

func getConfig() Config {
	var error error
	var port, strMin, strMax int

	port, error = strconv.Atoi(os.Getenv("PASSMAN_PORT"))
	if error != nil {
		log.Fatal(error)
	}
	strMin, error = strconv.Atoi(os.Getenv("PASSMAN_STRING_MIN"))
	if error != nil {
		log.Fatal(error)
	}
	strMax, error = strconv.Atoi(os.Getenv("PASSMAN_STRING_MAX"))
	if error != nil {
		log.Fatal(error)
	}

	c := Config{
		Port:            port,
		StringMinLength: strMin,
		StringMaxLength: strMax,
	}

	return c
}

type Config struct {
	Port            int
	StringMaxLength int
	StringMinLength int
}
