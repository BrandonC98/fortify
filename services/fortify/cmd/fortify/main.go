package main

import (
	"github.com/BrandonC98/fortify/services/fortify/internal/model"
	"github.com/BrandonC98/fortify/services/fortify/internal/server"
)

func main() {
	server.StartServer(model.GetConfig())
}
