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
	var passGenURL, dbUser, dbHost, dbPassword string

	port, errorMsg = strconv.Atoi(os.Getenv("PASSMAN_PORT"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}

	passGenURL = os.Getenv("PASSMAN_PASS_GEN_URL")
	dbUser = os.Getenv("DB_USER")
	dbHost = os.Getenv("DB_HOST")
	dbPassword = os.Getenv("DB_PASSWORD")

	c := Config{
		Port:       port,
		PassGenURL: passGenURL,
		DBUser:     dbUser,
		DBHost:     dbHost,
		DBPassword: dbPassword,
	}

	return c
}
