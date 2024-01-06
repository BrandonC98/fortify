package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	return gin.Default()
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func generateHandler(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string
		if gin.Mode() == "release" {
			// use AWS Secrets manager to get key
		} else {
			key = "PASSMAN_PASS_GEN_KEY"
		}

		plaintextPassword, err := generatePassword(randNumber(config.StringMinLength, config.StringMaxLength))
		if err != nil {
			log.Fatal(err)
		}

		ecryptedPassword, err := helper.EncryptMessageWithPassword([]byte(key), plaintextPassword)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": ecryptedPassword,
		})
	}
}

func startServer(config Config) {
	router := router()
	gin.SetMode(config.Mode)

	router.GET("/ping", pingHandler)

	router.GET("/generate", generateHandler(config))

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}
