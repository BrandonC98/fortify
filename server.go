package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"
)

func startServer(config Config) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	var client standardHTTPClient

	router.GET("/ping", pingHandler)

	router.GET("/", homeHandler)

	router.GET("/generatePassword", generatePasswordHandler(fmt.Sprintf("%s/generate", config.PassGenURL), &client))

	router.POST("/save")

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func saveHandler(name string, password string) {

}

func generatePasswordHandler(endpointURL string, client HTTPClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string

		if gin.Mode() == "release" {
			// use aws secrets manager to get key
			println("Functionality not yet implmented")
		} else {
			key = "PASSMAN_PASS_GEN_KEY"
		}

		password := getGeneratedPassword(endpointURL, client)

		plainPassword, err := helper.DecryptMessageWithPassword([]byte(key), password)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": plainPassword})
	}
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "main website",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
