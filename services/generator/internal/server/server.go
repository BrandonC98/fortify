package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"

	g "github.com/BrandonC98/fortify/services/generator/internal/generation"
	"github.com/BrandonC98/fortify/services/generator/internal/model"
)

func router() *gin.Engine {
	return gin.Default()
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func generateHandler(config model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := config.Key

		plaintextPassword, err := g.GenerateRandString(g.RandNumber(config.StringMinLength, config.StringMaxLength))
		if err != nil {
			log.Fatal(err)
		}

		ecryptedPassword, err := helper.EncryptMessageWithPassword([]byte(key), plaintextPassword)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, ecryptedPassword)
	}
}

func StartServer(config model.Config) {
	router := router()

	router.GET("/ping", pingHandler)

	router.GET("/generate", generateHandler(config))

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}
