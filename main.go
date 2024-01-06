package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	config := getConfig()
	startServer(config)
}

func getConfig() Config {
	f, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

type Config struct {
	Mode            string `yaml:"mode"`
	Port            int    `yaml:"port"`
	StringMaxLength int    `yaml:"stringMax"`
	StringMinLength int    `yaml:"stringMin"`
}
