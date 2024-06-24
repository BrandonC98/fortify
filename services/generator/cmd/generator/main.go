package main

import (
	"github.com/BrandonC98/fortify/services/generator/internal/model"
	"github.com/BrandonC98/fortify/services/generator/internal/server"
)

func main() {
	config := model.GetConfig()
	server.StartServer(config)
}
