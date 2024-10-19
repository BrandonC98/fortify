package model

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port            int
	StringMaxLength int
	StringMinLength int
	Key             string
}

func GetConfig() Config {
	var errorMsg error
	var port, strMin, strMax int
	var key string

	port, errorMsg = strconv.Atoi(os.Getenv("PORT"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	strMin, errorMsg = strconv.Atoi(os.Getenv("STRING_MIN"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	strMax, errorMsg = strconv.Atoi(os.Getenv("STRING_MAX"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	key = os.Getenv("ENCRYPTION_KEY")

	c := Config{
		Port:            port,
		StringMinLength: strMin,
		StringMaxLength: strMax,
		Key:             key,
	}

	return c
}
