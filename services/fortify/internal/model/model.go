package model

import (
	"log"
	"os"
	"strconv"
)

type Secret struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Value string
}

type Config struct {
	Port         int
	GeneratorURL string
	DBUser       string
	DBHost       string
	DBPassword   string
	Key          string
}

func GetConfig() Config {
	var errorMsg error
	var port int
	var key, generatorURL, dbUser, dbHost, dbPassword string

	port, errorMsg = strconv.Atoi(os.Getenv("PORT"))
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}

	generatorURL = os.Getenv("GENERATOR_URL")
	dbUser = os.Getenv("DB_USER")
	dbHost = os.Getenv("DB_HOST")
	dbPassword = os.Getenv("DB_PASSWORD")
	key = os.Getenv("ENCRYPTION_KEY")

	c := Config{
		Port:         port,
		GeneratorURL: generatorURL,
		DBUser:       dbUser,
		DBHost:       dbHost,
		DBPassword:   dbPassword,
		Key:          key,
	}

	return c
}
