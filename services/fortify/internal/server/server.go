package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"

	"github.com/BrandonC98/fortify/services/fortify/internal/database"
	"github.com/BrandonC98/fortify/services/fortify/internal/model"
)

func StartServer(config model.Config) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	var client standardHTTPClient
	r := database.NewSecretsRepository(config.DBHost, "fortify_db", config.DBUser, config.DBPassword)
	r.Setup()

	router.GET("/ping", pingHandler)
	router.GET("/", homeHandler)
	router.GET("/generate", generateHandler(fmt.Sprintf("%s/generate", config.GeneratorURL), &client))
	router.POST("/save", saveHandler(r))
	router.GET("/show", showHandler(r))

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func showHandler(r database.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		creds := r.RetriveAllRecords()
		var sb strings.Builder
		for i := 0; i < len(creds); i++ {
			sb.WriteString(fmt.Sprint(creds[i].Name, ": ", creds[i].Value, "\n"))
		}

		c.String(http.StatusOK, sb.String())
	}
}

func saveHandler(r database.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials model.Secret

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		r.AddRecord(&credentials)
		c.String(http.StatusOK, "successful")
	}
}

func generateHandler(endpointURL string, client HTTPClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string

		if gin.Mode() == "release" {
			// use aws secrets manager to get key
			println("Functionality not yet implmented")
		} else {
			key = "GENERATOR_KEY"
		}

		s := generate(endpointURL, client)

		s, err := helper.DecryptMessageWithPassword([]byte(key), s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": s})
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
