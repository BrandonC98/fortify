package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func startServer(config Config) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	var client standardHTTPClient
	db := Database(config)

	router.GET("/ping", pingHandler)

	router.GET("/", homeHandler)

	router.GET("/generatePassword", generatePasswordHandler(fmt.Sprintf("%s/generate", config.PassGenURL), &client))

	router.POST("/save", saveHandler(db))
	router.GET("/show", showHandler(db))

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func showHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		creds := retriveAllCreds(db)
		var sb strings.Builder
		for i := 0; i < len(creds); i++ {
			sb.WriteString(fmt.Sprint(creds[i].Name, ": ", creds[i].Passwd, "\n"))
		}

		slog.Info("List: " + sb.String())

		c.String(200, sb.String())
	}
}

func saveHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Saving")
		var credentials Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		slog.Info("Pass => " + credentials.Passwd)
		AddCredsRecord(&credentials, db)

	}

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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
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
